package biome

import "testing"

func TestSeededRandom(t *testing.T) {
	rand := NewSeededRandom(0)
	for i := 0; i < 10; i++ {
		t.Log(rand.NextInt(10))
	}
}
