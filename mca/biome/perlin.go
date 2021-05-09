package biome

import (
	"math"

	"kmschr.com/mc2brs/util"
)

type PerlinNoiseGen struct {
	noiseLevels            []SimplexNoiseGen
	highestFreqValueFactor float64
	highestFreqInputFactor float64
}

func NewPerlinNoiseGen(seed SharedSeedRandom, octaves []int) PerlinNoiseGen {
	p := PerlinNoiseGen{}

	i := -octaves[0]
	j := octaves[len(octaves)-1]
	k := i + j + 1

	simplexNoiseGen := NewSimplexNoiseGen(seed)
	l := j

	p.noiseLevels = make([]SimplexNoiseGen, k)
	if j >= 0 && j < k && util.Contains(octaves, 0) {
		p.noiseLevels[j] = simplexNoiseGen
	}

	for i1 := j + 1; i1 < k; i1++ {
		if i1 >= 0 && util.Contains(octaves, l-i1) {
			p.noiseLevels[i1] = NewSimplexNoiseGen(seed)
		} else {
			seed.ConsumeCount(262)
		}
	}

	if j > 0 {
		k1 := uint64(simplexNoiseGen.GetValue(simplexNoiseGen.xo, simplexNoiseGen.yo, simplexNoiseGen.zo) * 9.223372e18)
		sharedseedrandom := NewSeededRandom(k1)

		for j1 := l - 1; j1 >= 0; j1-- {
			if j1 < k && util.Contains(octaves, l-j1) {
				p.noiseLevels[j1] = NewSimplexNoiseGen(sharedseedrandom)
			} else {
				sharedseedrandom.ConsumeCount(262)
			}
		}
	}

	p.highestFreqInputFactor = math.Pow(2, float64(j))
	p.highestFreqValueFactor = 1.0 / (math.Pow(2, float64(k)) - 1.0)

	return p
}

func (p *PerlinNoiseGen) GetValue(x float64, z float64, amp bool) float64 {
	d0 := 0.0
	d1 := p.highestFreqInputFactor
	d2 := p.highestFreqValueFactor

	for _, simplexnoisegen := range p.noiseLevels {
		xAmp := 0.0
		zAmp := 0.0
		if amp {
			xAmp = simplexnoisegen.xo
			zAmp = simplexnoisegen.yo
		}
		d0 += simplexnoisegen.GetValue2(x*d1+xAmp, z*d1+zAmp*d2) * d2

		d1 /= 2.0
		d2 *= 2.0
	}

	return d0
}
