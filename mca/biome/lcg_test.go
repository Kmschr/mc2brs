package biome

import "testing"

func TestRandomness(t *testing.T) {
	rand := NewLCG(0)
	zero := 0
	ones := 0
	for i := 0; i < 1000; i++ {
		bits := rand.Next(1)
		if bits == 0 {
			zero++
		}
		if bits == 1 {
			ones++
		}
	}
	t.Log(ones)
	t.Log(zero)
}
