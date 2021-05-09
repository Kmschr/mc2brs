package util

import (
	"image/color"
)

const blendRadius = 8

func RGBInt(c color.RGBA) int {
	return int(c.R)<<16 | int(c.G)<<8 | int(c.B)
}

func IntRGB(c int) color.RGBA {
	return color.RGBA{
		R: byte(c >> 16),
		G: byte(c >> 8),
		B: byte(c),
		A: 255,
	}
}

func TintColor(base color.RGBA, tint color.RGBA) color.RGBA {
	blendR := float32(base.R) / 255.0
	blendG := float32(base.G) / 255.0
	blendB := float32(base.B) / 255.0
	blendA := float32(base.A) / 255.0
	r := float32(tint.R) * blendR
	g := float32(tint.G) * blendG
	b := float32(tint.B) * blendB
	a := float32(tint.A) * blendA
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

func BlendColor(img [32 * 16][32 * 16]color.RGBA, x int, z int) color.RGBA {
	if img[x][z].A == 0 {
		return color.RGBA{0, 0, 0, 0}
	}

	n := 0
	r := 0
	g := 0
	b := 0
	for cx := x - blendRadius; cx < x+blendRadius; cx++ {
		for cz := z - blendRadius; cz < z+blendRadius; cz++ {
			if cx < 0 || cz < 0 || cx >= 32*16 || cz >= 32*16 {
				continue
			}
			curColor := img[cx][cz]
			if curColor.A == 0 {
				continue
			}
			r += int(curColor.R)
			g += int(curColor.G)
			b += int(curColor.B)
			n++
		}
	}

	avgR := uint8(r / n)
	avgG := uint8(g / n)
	avgB := uint8(b / n)
	return color.RGBA{avgR, avgG, avgB, img[x][z].A}
}
