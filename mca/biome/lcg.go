package biome

// linear congruential pseudorandom number generator
type LCG struct {
	seed uint64
}

func NewLCG(seed uint64) LCG {
	return LCG{seed}
}

func (l *LCG) Next(bits int) int {
	l.updateSeed()
	return int(l.seed >> (48 - bits))
}

func (l *LCG) updateSeed() {
	l.seed = (l.seed*0x5DEECE66D + 0xB) & ((1 << 48) - 1)
}
