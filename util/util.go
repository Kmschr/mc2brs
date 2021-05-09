package util

import (
	"image"
	"image/color"
)

type Vec3 struct {
	X int
	Z int
	Y int
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func PadImage(img *image.RGBA) *image.RGBA {
	bounds := img.Bounds()
	newWidth := bounds.Max.Y * 741 / 417

	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, bounds.Max.Y))

	offset := (newWidth - bounds.Max.X) / 2

	for y := 0; y < bounds.Max.Y-1; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			newImg.SetRGBA(x+offset, y, img.RGBAAt(x, y))
		}
	}

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < offset; x++ {
			newImg.SetRGBA(x, y, color.RGBA{27, 47, 74, 255})
			newImg.SetRGBA(newWidth-x, y, color.RGBA{27, 47, 74, 255})
		}
	}
	return newImg
}

func MostFrequent(s []int32) int32 {
	m := map[int32]int{}
	var maxCnt int
	var freq int32
	for _, a := range s {
		m[a]++
		if m[a] > maxCnt {
			maxCnt = m[a]
			freq = a
		}
	}

	return freq
}

func Clampf(n float32, lo float32, hi float32) float32 {
	if n > hi {
		n = hi
	} else if n < lo {
		n = lo
	}
	return n
}

func Clamp(n int, lo int, hi int) int {
	if n > hi {
		n = hi
	} else if n < lo {
		n = lo
	}
	return n
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
