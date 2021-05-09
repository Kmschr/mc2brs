package biome

import "math"

const SQRT_3 = 1.73205080757
const F2 = 0.5 * (SQRT_3 - 1)
const G2 = (3.0 - SQRT_3) / 6.0

var GRADIENT = [][]int{
	{1, 1, 0},
	{-1, 1, 0},
	{1, -1, 0},
	{-1, -1, 0},
	{1, 0, 1},
	{-1, 0, 1},
	{1, 0, -1},
	{-1, 0, -1},
	{0, 1, 1},
	{0, -1, 1},
	{0, 1, -1},
	{0, -1, -1},
	{1, 1, 0},
	{0, -1, 1},
	{-1, 1, 0},
	{0, -1, -1},
}

type SimplexNoiseGen struct {
	p  [512]int
	xo float64
	yo float64
	zo float64
}

func NewSimplexNoiseGen(rand SharedSeedRandom) SimplexNoiseGen {
	s := SimplexNoiseGen{}

	s.xo = rand.NextDouble() * 256.00
	s.yo = rand.NextDouble() * 256.00
	s.zo = rand.NextDouble() * 256.00

	for i := 0; i < 256; i++ {
		s.p[i] = i
	}

	for i := 0; i < 256; i++ {
		j := rand.NextInt(256 - i)
		k := s.p[i]
		s.p[i] = s.p[j+i]
		s.p[j+i] = k
	}

	return s
}

func (s *SimplexNoiseGen) P(i int) int {
	return s.p[i&255]
}

func (s *SimplexNoiseGen) GetValue2(x float64, z float64) float64 {
	d0 := (x + z) * F2
	i := int(math.Floor(x + d0))
	j := int(math.Floor(z + d0))
	d1 := float64(i+j) * G2
	d2 := float64(i) - d1
	d3 := float64(j) - d1
	d4 := x - d2
	d5 := z - d3
	var k int
	var l int
	if d4 > d5 {
		k = 1
	} else {
		l = 1
	}

	d6 := d4 - float64(k) + G2
	d7 := d5 - float64(l) + G2
	d8 := d4 - 1.0 + 2.0*G2
	d9 := d5 - 1.0 + 2.0*G2
	i1 := i & 255
	j1 := j & 255
	k1 := s.P(i1+s.P(j1)) % 12
	l1 := s.P(i1+k+s.P(j1+l)) % 12
	i2 := s.P(i1+1+s.P(j1+1)) % 12
	d10 := getCornerNoise3D(k1, d4, d5, 0.0, 0.5)
	d11 := getCornerNoise3D(l1, d6, d7, 0.0, 0.5)
	d12 := getCornerNoise3D(i2, d8, d9, 0.0, 0.5)
	return 70.0 * (d10 + d11 + d12)
}

func (s *SimplexNoiseGen) GetValue(x float64, y float64, z float64) float64 {
	d1 := (x + y + z) / 3
	i := int(math.Floor(x + d1))
	j := int(math.Floor(y + d1))
	k := int(math.Floor(z + d1))
	d3 := float64(i+j+k) / 6.0
	d4 := float64(i) - d3
	d5 := float64(j) - d3
	d6 := float64(k) - d3
	d7 := x - d4
	d8 := y - d5
	d9 := z - d6
	var l int
	var i1 int
	var j1 int
	var k1 int
	var l1 int
	var i2 int
	if d7 >= d8 {
		if d8 >= d9 {
			l = 1
			k1 = 1
			l1 = 1
		} else if d7 >= d9 {
			l = 1
			k1 = 1
			i2 = 1
		} else {
			j1 = 1
			k1 = 1
			i2 = 1
		}
	} else if d8 < d9 {
		j1 = 1
		l1 = 1
		i2 = 1
	} else if d7 < d9 {
		i1 = 1
		l1 = 1
		i2 = 1
	} else {
		i1 = 1
		k1 = 1
		l1 = 1
	}

	d10 := d7 - float64(l) + 1.0/6.0
	d11 := d8 - float64(i1) + 1.0/6.0
	d12 := d9 - float64(j1) + 1.0/6.0
	d13 := d7 - float64(k1) + 1.0/3.0
	d14 := d8 - float64(l1) + 1.0/3.0
	d15 := d9 - float64(i2) + 1.0/3.0
	d16 := d7 - 0.5
	d17 := d8 - 0.5
	d18 := d9 - 0.5
	j2 := i & 255
	k2 := j & 255
	l2 := k & 255
	i3 := s.P(j2+s.P(k2+s.P(l2))) % 12
	j3 := s.P(j2+l+s.P(k2+i1+s.P(l2+j1))) % 12
	k3 := s.P(j2+k1+s.P(k2+l1+s.P(l2+i2))) % 12
	l3 := s.P(j2+1+s.P(k2+1+s.P(l2+1))) % 12
	d19 := getCornerNoise3D(i3, d7, d8, d9, 0.6)
	d20 := getCornerNoise3D(j3, d10, d11, d12, 0.6)
	d21 := getCornerNoise3D(k3, d13, d14, d15, 0.6)
	d22 := getCornerNoise3D(l3, d16, d17, d18, 0.6)
	return 32.0 * (d19 + d20 + d21 + d22)

}

func dot(grad []int, x float64, y float64, z float64) float64 {
	return float64(grad[0])*x + float64(grad[1])*y + float64(grad[2])*z
}

func getCornerNoise3D(i int, x float64, y float64, z float64, dist float64) float64 {
	d1 := dist - x*x - y*y - z*z
	var d0 float64
	if d1 < 0.0 {
		d0 = 0.0
	} else {
		d1 = d1 * d1
		d0 = d1 * d1 * dot(GRADIENT[i], x, y, z)
	}
	return d0
}
