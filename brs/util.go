package brs

import (
	"image/color"
	"math"
)

func isZero(i [3]int) bool {
	return i[0] == 0 && i[1] == 0 && i[2] == 0
}

func combineOrientation(dir byte, rot byte) byte {
	return (dir&0x7)<<2 | rot&0x3
}

func colorBytes(c color.RGBA) []byte {
	return []byte{c.R, c.G, c.B}
}

func ConvertColor(c color.RGBA) color.RGBA {
	return color.RGBA{
		R: convert(c.R),
		G: convert(c.G),
		B: convert(c.B),
		A: c.A,
	}
}

func convert(n uint8) uint8 {
	u := float64(n) / 255.0
	if u <= 0.04045 {
		return uint8(((u * 25) / 323) * 255.0)
	}
	p := (u + 0.055) / 1.055
	return uint8(math.Pow(p, 2.4) * 255.0)
}
