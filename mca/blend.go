package mca

import (
	"image/color"

	"kmschr.com/mc2brs/mca/biome"
)

// blendRadius describes how many blocks will be considered when blending biomes
// the resulting area that gets iterated is (2*blendRadius)^2 many blocks
const blendRadius = 8

// BiomeBlend returns the blend of biomes at a given location
func (w *World) BiomeBlend(blockX int, blockZ int, y int) map[int32]int {
	blend := make(map[int32]int)
	for x := blockX - blendRadius; x < blockX+blendRadius; x++ {
		for z := blockZ - blendRadius; z < blockZ+blendRadius; z++ {
			b := w.BiomeAt(x, z, y)
			if b != -1 {
				blend[b]++
			}
		}
	}
	return blend
}

// GrassBlend returns the blend of grass color at any block
func (w *World) GrassBlend(x, z, y int) color.RGBA {
	blend := w.BiomeBlend(x, z, y)
	r := 0
	g := 0
	b := 0
	total := 0
	for bi, freq := range blend {
		c := biome.GrassColor(bi, x, z, y)
		r += int(c.R) * freq
		g += int(c.G) * freq
		b += int(c.B) * freq
		total += freq
	}
	return color.RGBA{byte(r / total), byte(g / total), byte(b / total), 255}
}

// FoliageBlend returns the blend of foliage color at any block
func (w *World) FoliageBlend(x, z, y int) color.RGBA {
	blend := w.BiomeBlend(x, z, y)
	r := 0
	g := 0
	b := 0
	total := 0
	for bi, freq := range blend {
		c := biome.FoliageColor(bi, x, z, y)
		r += int(c.R) * freq
		g += int(c.G) * freq
		b += int(c.B) * freq
		total += freq
	}
	return color.RGBA{byte(r / total), byte(g / total), byte(b / total), 240}
}

// WaterBlend returns the blend of water color at any block
func (w *World) WaterBlend(x, z, y int) color.RGBA {
	blend := w.BiomeBlend(x, z, y)
	r := 0
	g := 0
	b := 0
	a := 0
	total := 0
	for bi, freq := range blend {
		c := biome.Biomes[bi].WaterColor
		r += int(c.R) * freq
		g += int(c.G) * freq
		b += int(c.B) * freq
		a += int(c.A) * freq
		total += freq
	}
	return color.RGBA{byte(r / total), byte(g / total), byte(b / total), byte(a / total)}
}
