package biome

type SharedSeedRandom struct {
	rand  LCG
	count int
}

func NewSeededRandom(seed uint64) SharedSeedRandom {
	source := NewLCG(seed)
	return SharedSeedRandom{
		rand:  source,
		count: 0,
	}
}

func (s *SharedSeedRandom) ConsumeCount(count int) {
	for i := 0; i < count; i++ {
		s.next(1)
	}
}

func (s *SharedSeedRandom) next(bits int) int {
	s.count++
	return s.rand.Next(bits)
}

func (s *SharedSeedRandom) NextInt(bound int) int {
	if bound&-bound == bound { // bound is a power of 2
		return int(int64(bound) * int64(s.next(31)) >> 31)
	}

	val := 0
	bits := 0
	for {
		bits = s.next(31)
		val = bits % bound

		if bits-val+(bound-1) >= 0 {
			break
		}
	}

	return val
}

func (s *SharedSeedRandom) NextDouble() float64 {
	return float64(int64(s.next(26))<<27+int64(s.next(27))) / float64(1<<53)
}
