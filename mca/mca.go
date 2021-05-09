package mca

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/mca/biome"
	"kmschr.com/mc2brs/util"
)

// genHeightmapImg creates an image to visualize a region's heightmap
func (r *Region) genHeightmapImg(name string) {
	heightmap := image.NewRGBA(image.Rect(0, 0, regionWidth, regionWidth))
	for chunkZ := 0; chunkZ < 32; chunkZ++ {
		for chunkX := 0; chunkX < 32; chunkX++ {
			chunk, exists := r.Chunks[chunksIndex(chunkX, chunkZ)]
			if !exists {
				continue
			}
			buf, exists := chunk.Heightmaps[name]
			if !exists {
				continue
			}
			//fmt.Println(buf)
			var heights [256]byte
			for i, vals := range buf {
				h := readHeightsFromLong(vals)
				for j := 0; j < 7; j++ {
					ind := i*7 + j
					if ind >= 256 {
						break
					}
					heights[ind] = h[j]
				}
			}
			for cz := 0; cz < 16; cz++ {
				for cx := 0; cx < 16; cx++ {
					gray := heights[cz*16+cx]
					heightmap.SetRGBA(chunkX*16+cx, chunkZ*16+cz, color.RGBA{gray, gray, gray, 255})
				}
			}
		}
	}
	output, _ := os.Create(fmt.Sprintf("%s.png", name))
	png.Encode(output, heightmap)
	output.Close()
}

// genHeightmapImgs creates images to visualize a region's heightmaps
func (r *Region) GenHeightmapImgs() {
	r.genHeightmapImg("MOTION_BLOCKING")
	r.genHeightmapImg("MOTION_BLOCKING_NO_LEAVES")
	r.genHeightmapImg("OCEAN_FLOOR")
	r.genHeightmapImg("OCEAN_FLOOR_WG")
	r.genHeightmapImg("WORLD_SURFACE")
	r.genHeightmapImg("WORLD_SURFACE_WG")
}

// genColorImgs creates images to visualize possible grass, foliage, and water colors
// for every spot in a region
func (r *Region) GenColorImgs() {
	grassImg := image.NewRGBA(image.Rect(0, 0, regionWidth, regionWidth))
	foliageImg := image.NewRGBA(image.Rect(0, 0, regionWidth, regionWidth))
	waterImg := image.NewRGBA(image.Rect(0, 0, regionWidth, regionWidth))
	var grassColors [regionWidth][regionWidth]color.RGBA
	var foliageColors [regionWidth][regionWidth]color.RGBA
	var waterColors [regionWidth][regionWidth]color.RGBA
	for chunkZ := 0; chunkZ < 32; chunkZ++ {
		for chunkX := 0; chunkX < 32; chunkX++ {
			chunk, exists := r.Chunks[chunksIndex(chunkX, chunkZ)]
			if !exists {
				continue
			}
			heightmap, exists := chunk.Heightmaps["MOTION_BLOCKING_NO_LEAVES"]
			if !exists {
				continue
			}
			var heights [256]byte
			for i, vals := range heightmap {
				h := readHeightsFromLong(vals)
				for j := 0; j < 7; j++ {
					ind := i*7 + j
					if ind >= 256 {
						break
					}
					heights[ind] = h[j]
				}
			}
			for cz := 0; cz < 16; cz++ {
				for cx := 0; cx < 16; cx++ {
					rx := chunkX*16 + cx
					rz := chunkZ*16 + cz

					y := int(heights[cz*16+cx])
					b := chunk.biomeAt(cx, cz, y)
					grassColor := biome.GrassColor(b, rx, rz, y)
					foliageColor := biome.FoliageColor(b, rx, rz, y)
					waterColor := biome.Biomes[b].WaterColor
					grassColors[rx][rz] = grassColor
					foliageColors[rx][rz] = foliageColor
					waterColors[rx][rz] = waterColor
				}
			}
		}
	}
	for z := 0; z < regionWidth; z++ {
		for x := 0; x < regionWidth; x++ {
			grassBlended := util.BlendColor(grassColors, x, z)
			grassImg.SetRGBA(x, z, grassBlended)
			foliageBlended := util.BlendColor(foliageColors, x, z)
			foliageImg.SetRGBA(x, z, foliageBlended)
			waterBlended := util.BlendColor(waterColors, x, z)
			waterImg.SetRGBA(x, z, waterBlended)
		}
	}
	grassOutfile, _ := os.Create("grass.png")
	foliageOutfile, _ := os.Create("foliage.png")
	waterOutfile, _ := os.Create("water.png")
	png.Encode(grassOutfile, grassImg)
	png.Encode(foliageOutfile, foliageImg)
	png.Encode(waterOutfile, waterImg)
	grassOutfile.Close()
	foliageOutfile.Close()
	waterOutfile.Close()
}

// makeMinimap creates a colored and shaded image that depicts the region
// from a top down perspective
func (w *World) makeMinimap(r Region) *image.RGBA {
	maxY := 256
	if ansi.Dimension() == "nether" {
		maxY = 64
	}
	ansi.Println(ansi.BrightMagenta, "\nGenerating Preview Image...")
	var topBlocks [regionWidth][regionWidth]BlockState
	var elevations [regionWidth][regionWidth]int
	minimap := image.NewRGBA(image.Rect(0, 0, regionWidth, regionWidth))
	ansi.Print(ansi.BrightYellow, "Processing")
	ansi.Print(ansi.BrightBlue, " [")
	for rx := 0; rx < regionWidth; rx++ {
		if rx%32 == 0 {
			ansi.Print(ansi.BrightBlue, "#")
		}
		for rz := 0; rz < regionWidth; rz++ {
			for y := 0; y < maxY; y++ {
				block, exists := r.BlockAt(rx, rz, y)
				if !exists || block.isSkipped() {
					continue
				}
				topBlocks[rx][rz] = block
				elevations[rx][rz] = y
			}
		}
	}
	ansi.Println(ansi.BrightBlue, "]")
	dx := r.RegionX * regionWidth
	dz := r.RegionZ * regionWidth
	ansi.Print(ansi.BrightYellow, "Coloring")
	ansi.Print(ansi.BrightBlue, "   [")
	for rx := 0; rx < regionWidth; rx++ {
		if rx%32 == 0 {
			ansi.Print(ansi.BrightBlue, "#")
		}
		for rz := 0; rz < regionWidth; rz++ {
			block := topBlocks[rx][rz]
			x := rx + dx
			z := rz + dz
			if w.surrounded(block, x, z, elevations[rx][rz]) {
				continue
			}
			c, known := w.ColorAt(block, x, z, elevations[rx][rz])
			if !known {
				continue
			}
			minimap.SetRGBA(rx, rz, c)
		}
	}
	ansi.Println(ansi.BrightBlue, "]")
	// Apply lighting
	ansi.Print(ansi.BrightYellow, "Lighting")
	ansi.Print(ansi.BrightBlue, "   [")
	for rx := 0; rx < regionWidth; rx++ {
		if rx%32 == 0 {
			ansi.Print(ansi.BrightBlue, "#")
		}
		for rz := 0; rz < regionWidth; rz++ {
			if rx == (regionWidth)-1 || rz == (regionWidth)-1 || rx == 0 || rz == 0 {
				continue
			}
			cur := minimap.RGBAAt(rx, rz)
			y := elevations[rx][rz]
			top := util.Clamp(elevations[rx][rz+1]-y, 0, 4)
			cornerTR := util.Clamp(elevations[rx+1][rz]-y, 0, 4)
			right := util.Clamp(elevations[rx+1][rz+1]-y, 0, 4)
			cornerBR := util.Clamp(elevations[rx+1][rz-1]-y, 0, 4)
			bot := util.Clamp(elevations[rx][rz-1]-y, 0, 4)
			cornerBL := util.Clamp(elevations[rx-1][rz-1]-y, 0, 4)
			shade := byte(255 - (top+cornerTR+right)*6 - (cornerBR+bot+cornerBL)*3)
			base := color.RGBA{shade, shade, shade, 255}
			minimap.SetRGBA(rx, rz, util.TintColor(base, cur))
		}
	}
	ansi.Println(ansi.BrightBlue, "]")
	return minimap
}

func readHeightsFromLong(b int64) [7]byte {
	//fmt.Printf("%064b\n", b)
	var heights [7]byte
	for i := 0; i < 7; i++ {
		heights[i] = byte(b)
		b = b >> 9
	}
	return heights
}
