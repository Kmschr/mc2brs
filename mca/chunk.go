package mca

import (
	"image/color"
	"math/rand"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/mca/biome"
	"kmschr.com/mc2brs/mca/nbt"
	"kmschr.com/mc2brs/util"
)

// Chunk is a 16x16x256 area of blocks
type Chunk struct {
	DataVersion  int32
	Biomes       []int32
	Heightmaps   map[string][]int64
	CarvingMasks map[string][]byte
	Sections     map[int]Section
	Status       string
	maxSection   int
	minSection   int
}

func (c *Chunk) Symbol() string {
	var symbol string

	avgBiome := util.MostFrequent(c.Biomes)
	biomeSymbol, found := biome.BiomeSymbols[avgBiome]
	if found {
		heightmap, exists := c.Heightmaps["OCEAN_FLOOR"]
		elevation := 64
		if exists {
			elevation = int(heightmap[18] & 0x1FF)
		}
		var biomeColor color.RGBA
		if biome.IsWater(avgBiome) {
			biomeColor = biome.Biomes[avgBiome].WaterColor
		} else {
			biomeColor = biome.GrassColor(avgBiome, 8, 8, elevation)
		}
		symbol = ansi.Color(biomeColor)
		if len(biomeSymbol) > 1 {
			if biome.IsMountain(avgBiome) {
				if elevation < 80 {
					symbol += biomeSymbol[0]
				} else if elevation < 130 {
					symbol += biomeSymbol[1]
				} else {
					symbol += biomeSymbol[2]
				}
			} else if biome.IsBadlandsWoodedPlateau(avgBiome) {
				if elevation < 82 {
					symbol += biomeSymbol[0]
				} else {
					symbol += biomeSymbol[1]
				}
			} else {
				index := rand.Intn(2)
				symbol += biomeSymbol[index]
			}
		} else {
			symbol += biomeSymbol[0]
		}
	} else {
		symbol = "?"
	}

	return symbol
}

// newChunk creates a chunk using a root tag
func newChunk(rootNBT nbt.Tag) Chunk {
	chunk := Chunk{}
	levelNBT, exists := rootNBT.Compound()["Level"]
	if !exists {
		ansi.Println(ansi.Red, "Chunk does not contain level tag")
	}
	dataVersionNBT := rootNBT.Compound()["DataVersion"]
	statusNBT := levelNBT.Compound()["Status"]
	biomesNBT := levelNBT.Compound()["Biomes"]
	if dataVersionNBT != nil {
		chunk.DataVersion = dataVersionNBT.Int()
	}
	if biomesNBT != nil {
		chunk.Biomes = biomesNBT.IntArray()
	}
	if statusNBT != nil {
		chunk.Status = levelNBT.Compound()["Status"].String()
	}
	chunk.processMapsAndMasks(levelNBT)
	chunk.processSections(levelNBT)
	return chunk
}

// processMapsAndMasks processes any available heightmaps/carving masks
func (c *Chunk) processMapsAndMasks(levelNBT nbt.Tag) {
	heightmapsNBT := levelNBT.Compound()["Heightmaps"]
	carvingMasksNBT := levelNBT.Compound()["CarvingMasks"]
	if heightmapsNBT != nil {
		c.Heightmaps = make(map[string][]int64)
		for _, heightmapNBT := range heightmapsNBT.Compound() {
			c.Heightmaps[heightmapNBT.Name()] = heightmapNBT.LongArray()
		}
	}
	if carvingMasksNBT != nil {
		c.CarvingMasks = map[string][]byte{}
		for _, carvingMaskNBT := range carvingMasksNBT.Compound() {
			c.CarvingMasks[carvingMaskNBT.Name()] = carvingMaskNBT.ByteArray()
		}
	}
}

// processSections processes and organizes the sections within a chunk
func (c *Chunk) processSections(levelNBT nbt.Tag) {
	sectionsNBT := levelNBT.Compound()["Sections"]
	if sectionsNBT == nil {
		return
	}
	c.Sections = make(map[int]Section)
	for _, sectionNBT := range sectionsNBT.List() {
		sectionIndex := int8(sectionNBT.Compound()["Y"].Byte())
		section := newSection(sectionNBT, c.DataVersion)
		c.Sections[int(sectionIndex)] = section
		if sectionIndex > int8(c.maxSection) {
			c.maxSection = int(sectionIndex)
		}
		if sectionIndex < int8(c.minSection) {
			c.minSection = int(sectionIndex)
		}
	}
}

// biomeAt accepts positions in unmodified chunk, region, and world space because of the byte masking
func (c *Chunk) biomeAt(x, z, y int) int32 {
	if c.Biomes == nil || len(c.Biomes) == 0 {
		return -1
	}
	if c.DataVersion < 2202 {
		return c.Biomes[biomeIndexOld(x, z)]
	}
	return c.Biomes[biomeIndex(x, z, y)]
}

// BlockAt returns the blockstate of any position within that chunk
func (c *Chunk) BlockAt(x, z, y int) (BlockState, bool) {
	section, exists := c.Sections[y>>4]
	if !exists {
		return BlockState{}, false
	}
	return section.blockAt(x, z, y)
}

// biomeIndex gets the index within a chunk's Biome array for a given position in that chunk
// Biomes has 1024 entries
// Each biome cube is 4x4x4
// X is between 0-16 (4 bits), only top 2 bits are used to represent 4 cubes horizontally
// Z is between 0-16 (4 bits), only top 2 bits are used to represent 4 cubes horizontally
// Y is between 0-256 (8 bits), only top 6 bits are used to represent 64 cubes vertically
// Final int in format YYYYYYZZXX
func biomeIndex(x, z, y int) int {
	return (y&0xFC)<<2 | z&0xC | (x&0xF)>>2
}

// biomeIndexOld gets the index within a chunk's Biome array for a given posiiton in that chunk
// Biomes has 256 entries
func biomeIndexOld(x int, z int) int {
	return (z&0xF)<<4 | x&0xF
}
