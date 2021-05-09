package mca

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/util"
)

// World represents an entire minecraft world which contains regions
type World struct {
	Name      string
	Path      string
	Regions   map[regionCoords]Region
	Loaded    map[regionCoords]bool
	Dimension string
}

// NewWorld creates a World for a given world directory
func NewWorld(name string, path string) World {
	world := World{
		Name:      name,
		Path:      path,
		Regions:   make(map[regionCoords]Region),
		Loaded:    make(map[regionCoords]bool),
		Dimension: ansi.Dimension(),
	}
	world.indexRegions()
	return world
}

var dimensionFolders = map[string]string{
	"overworld": "\\region",
	"nether":    "\\DIM-1\\region",
	"end":       "\\DIM1\\region",
}

var dimensionColors = map[string][]string{
	"overworld": {
		ansi.RGB(59, 73, 29),
		ansi.RGB(121, 150, 60),
		ansi.RGB(159, 196, 78),
	},
	"nether": {
		ansi.RGB(77, 18, 18),
		ansi.RGB(87, 28, 28),
		ansi.RGB(97, 38, 38),
	},
	"end": {
		ansi.RGB(187, 180, 115),
		ansi.RGB(207, 200, 135),
		ansi.RGB(217, 220, 155),
	},
}

// indexRegions finds all regions in a world and stages them for reading
func (w *World) indexRegions() {
	regionList, err := ioutil.ReadDir(w.Path + dimensionFolders[w.Dimension])
	if err != nil {
		ansi.Println(ansi.Red, "error reading region folder for world")
		ansi.Quit()
	}
	for _, regionInfo := range regionList {
		parts := strings.Split(regionInfo.Name(), ".")
		if parts[0] != "r" {
			continue
		}
		rx, _ := strconv.Atoi(parts[1])
		rz, _ := strconv.Atoi(parts[2])
		w.Regions[regionCoords{rx, rz}] = NewRegion(w.regionPath(rx, rz), rx, rz)
	}
}

// Overview prints console art to show the general layout of regions in the world
func (w *World) Overview() {
	fmt.Println("Found", len(w.Regions), w.Dimension, "regions")
	b := bounds{}
	for coords := range w.Regions {
		b.updateBounds(coords.X, coords.Z)
	}
	fmt.Printf("(%v,%v) -> (%v,%v)\n", b.X1, b.Z1, b.X2, b.Z2)
	fmt.Println("World Map:")

	minZ := util.Max(b.Z1, -18)
	maxZ := util.Min(b.Z2, 18)
	minX := util.Max(b.X1, -18)
	maxX := util.Min(b.X2, 18)

	for x := minX - 3; x < 0; x++ {
		fmt.Print("  ")
	}
	fmt.Printf("-Z (%s)\n", ansi.Sprint(ansi.BrightGreen, "NORTH"))

	for z := minZ; z <= maxZ; z++ {
		if z == 0 {
			fmt.Printf("\x1b[0m -X (%s) ", ansi.Sprint(ansi.BrightGreen, "WEST"))
		} else {
			fmt.Print("\x1b[0m           ")
		}
		fmt.Print("\x1b[48;2;0;0;0m")
		for x := minX; x <= maxX; x++ {
			sym := "\u2588\u2588"
			if x == 0 || z == 0 {
				sym = "\u2593\u2593"
			}
			if x == 0 && z == 0 {
				sym = "\u2588\u2588"
			}
			r, exists := w.Regions[regionCoords{x, z}]
			if !exists {
				fmt.Print(ansi.RGB(20, 20, 20) + sym)
				continue
			}
			size := len(r.ChunkTable)
			if size < 256 {
				fmt.Print(dimensionColors[w.Dimension][0] + sym)
			} else if size < 1024 {
				fmt.Print(dimensionColors[w.Dimension][1] + sym)
			} else {
				fmt.Print(dimensionColors[w.Dimension][2] + sym)
			}
		}
		if z == 0 {
			fmt.Printf("\x1b[0m +X (%s)\n", ansi.Sprint(ansi.BrightGreen, "EAST"))
			continue
		}
		fmt.Println("\x1b[0m")
	}

	for x := minX - 3; x < 0; x++ {
		fmt.Print("  ")
	}
	fmt.Printf("+Z (%s)\n", ansi.Sprint(ansi.BrightGreen, "SOUTH"))
}

// SelectRegion lets a user select which region to load in a world
func (w *World) SelectRegion() (int, int) {
	rxString := ansi.Sprint(ansi.BrightBlue, "RegionX")
	rzString := ansi.Sprint(ansi.BrightBlue, "RegionZ")
	fmt.Printf("Regions are selected in format %s, %s\n", rxString, rzString)
	fmt.Println("Regions are an area of 32 by 32 chunks")
	fmt.Println("You can determine the region from in game coordinates")
	fmt.Printf("%s = %s(x / 512)\n", rxString, ansi.Sprint(ansi.BrightRed, "floor"))
	fmt.Printf("%s = %s(z / 512)\n", rzString, ansi.Sprint(ansi.BrightRed, "floor"))
	for {
		selectedRegion := ansi.BasicPrompt("Enter region coords: (default 0, 0)")
		var x int
		var z int
		if selectedRegion != "" {
			selectedCoords := strings.Split(selectedRegion, ",")
			if len(selectedCoords) != 2 {
				continue
			}
			x, _ = strconv.Atoi(selectedCoords[0])
			z, _ = strconv.Atoi(strings.TrimSpace(selectedCoords[1]))
		}
		_, exists := w.Regions[regionCoords{x, z}]
		if exists {
			return x, z
		}
	}
}

// SurroundingAt returns the blocks surrounding a worldpsace coordinate
func (w *World) SurroundingAt(x, z, y int) []BlockState {
	var blocks []BlockState
	indices := [...][3]int{
		{x, z, y + 1},
		{x, z, y - 1},
		{x, z + 1, y},
		{x, z - 1, y},
		{x + 1, z, y},
		{x - 1, z, y},
	}
	for _, i := range indices {
		b, exists := w.BlockAt(i[0], i[1], i[2])
		if !exists {
			continue
		}
		blocks = append(blocks, b)
	}
	return blocks
}

// BlockAt returns the block of a given worldspace coordinate
func (w *World) BlockAt(x, z, y int) (BlockState, bool) {
	r, exists := w.RegionAt(x, z)
	if !exists {
		return BlockState{}, false
	}
	return r.BlockAt(x, z, y)
}

// BiomeAt returns the biome at any worldspace coordinate
func (w *World) BiomeAt(x, z, y int) int32 {
	r, exists := w.RegionAt(x, z)
	if !exists {
		return -1
	}
	return r.BiomeAt(x, z, y)
}

// RegionAt returns the region at a worldspace coordinate
func (w *World) RegionAt(x, z int) (Region, bool) {
	regionX := x >> 9
	regionZ := z >> 9
	r, exists := w.Regions[regionCoords{regionX, regionZ}]
	if !exists {
		return Region{}, false
	}
	return r, true
}

// LoadRegion reads the contents of the region file
func (w *World) LoadRegion(x, z int) {
	r, exists := w.Regions[regionCoords{x, z}]
	if !exists {
		ansi.Println(ansi.Red, "Tried to load region that doesnt exist")
		return
	}
	r.LoadAll()
	w.loadSurroundingChunks(x, z)
	w.Loaded[regionCoords{x, z}] = true
}

// loadSurroundingChunks loads the chunks in the immediate area around a region for blending purposes.
func (w *World) loadSurroundingChunks(x, z int) {
	regionLeft, haveRegionLeft := w.Regions[regionCoords{x - 1, z}]
	regionRight, haveRegionRight := w.Regions[regionCoords{x + 1, z}]
	regionAbove, haveRegionAbove := w.Regions[regionCoords{x, z + 1}]
	regionBelow, haveRegionBelow := w.Regions[regionCoords{x, z - 1}]
	for i := 0; i < 32; i++ {
		if haveRegionLeft {
			regionLeft.LoadChunk(31, i)
		}
		if haveRegionRight {
			regionRight.LoadChunk(0, i)
		}
		if haveRegionAbove {
			regionAbove.LoadChunk(i, 0)
		}
		if haveRegionBelow {
			regionBelow.LoadChunk(i, 31)
		}
	}

	regionTR, haveRegionTR := w.Regions[regionCoords{x + 1, z + 1}]
	regionBR, haveRegionBR := w.Regions[regionCoords{x + 1, z - 1}]
	regionBL, haveRegionBL := w.Regions[regionCoords{x - 1, z - 1}]
	regionTL, haveRegionTL := w.Regions[regionCoords{x - 1, z + 1}]
	if haveRegionTR {
		regionTR.LoadChunk(0, 0)
	}
	if haveRegionBR {
		regionBR.LoadChunk(0, 31)
	}
	if haveRegionBL {
		regionBL.LoadChunk(31, 31)
	}
	if haveRegionTL {
		regionTL.LoadChunk(31, 0)
	}
}

// regionPath gets the path of a region within a world
func (w *World) regionPath(x, z int) string {
	dimension := ansi.Dimension()
	return w.Path + fmt.Sprintf("%s\\r.%d.%d.mca", dimensionFolders[dimension], x, z)
}

// regionCoords contains x and z coordinates of a region
type regionCoords struct {
	X int
	Z int
}

type bounds struct {
	X1 int
	X2 int
	Z1 int
	Z2 int
}

func (b *bounds) updateBounds(x, z int) {
	if x > b.X2 {
		b.X2 = x
	}
	if x < b.X1 {
		b.X1 = x
	}
	if z > b.Z2 {
		b.Z2 = z
	}
	if z < b.Z1 {
		b.Z1 = z
	}
}
