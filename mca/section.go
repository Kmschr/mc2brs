package mca

import (
	"math"
	"math/big"

	"kmschr.com/mc2brs/mca/nbt"
)

// Section is a 16x16x16 cube of blocks within a chunk
type Section struct {
	BlockLight  []byte
	BlockStates []int64
	Palette     []BlockState
	SkyLight    []byte
	Y           int8
	DataVersion int32
}

// newSection creates a new section struct from an assumed to be valid section tag
func newSection(sectionNBT nbt.Tag, dataVersion int32) Section {
	s := Section{}
	s.DataVersion = dataVersion
	s.Y = int8(sectionNBT.Compound()["Y"].Byte())
	paletteNBT := sectionNBT.Compound()["Palette"]
	blockLightNBT := sectionNBT.Compound()["BlockLight"]
	blockStatesNBT := sectionNBT.Compound()["BlockStates"]
	skyLightNBT := sectionNBT.Compound()["SkyLight"]
	if paletteNBT == nil {
		return s
	}
	palette := paletteNBT.List()
	if palette == nil {
		return s
	}
	s.Palette = make([]BlockState, len(palette))
	for i, blockTag := range palette {
		s.Palette[i] = newBlock(blockTag)
	}
	if blockLightNBT != nil {
		s.BlockLight = blockLightNBT.ByteArray()
	}
	if blockStatesNBT != nil {
		s.BlockStates = blockStatesNBT.LongArray()
	}
	if skyLightNBT != nil {
		s.SkyLight = skyLightNBT.ByteArray()
	}
	return s
}

func (s *Section) BlocklightAt(x, z, y int) byte {
	if s.BlockLight == nil {
		return 0
	}
	index := (x + z*16 + y*16*16) >> 1
	if x&0x1 == 1 {
		return (s.BlockLight[index] & 0xF0 >> 4)
	}
	return s.BlockLight[index] & 0xF
}

func (s *Section) SkylightAt(x, z, y int) byte {
	if s.SkyLight == nil {
		return 0
	}
	index := (x + z*16 + y*16*16) >> 1
	if x&0x1 == 1 {
		return (s.SkyLight[index] & 0xF0 >> 4)
	}
	return s.SkyLight[index] & 0xF
}

// blockAt returns the blockstate of a position within the section
func (s *Section) blockAt(x, z, y int) (BlockState, bool) {
	if s.BlockStates == nil {
		return BlockState{}, false
	}
	paletteIndex := s.paletteIndex(sectionBlockIndex(x, z, y))
	if paletteIndex < 0 {
		return BlockState{}, false
	}
	return s.Palette[paletteIndex], true
}

// sectionBlockIndex returns the index of a block within the block states array of a section
func sectionBlockIndex(x, z, y int) int {
	return (y&0xF)<<8 | (z&0xF)<<4 | x&0xF
}

// paletteIndex returns the index of the block data within the section's palette
func (s *Section) paletteIndex(blockStateIndex int) int {
	bits := len(s.BlockStates) >> 6
	if s.DataVersion < 2527 {
		return s.paletteIndexOld(blockStateIndex, bits)
	}
	return s.paletteIndexCurrent(blockStateIndex, bits)
}

func (s *Section) paletteIndexCurrent(blockStateIndex int, bits int) int {
	indicesPerLong := int(64.0 / float64(bits))
	blockStatesIndex := blockStateIndex / indicesPerLong
	startBit := (blockStateIndex % indicesPerLong) * bits
	return int(bitRange(s.BlockStates[blockStatesIndex], startBit, startBit+bits))
}

func (s *Section) paletteIndexOld(blockStateIndex int, bits int) int {
	blockStatesIndex := float64(blockStateIndex) / (4096.0 / float64(len(s.BlockStates)))
	longIndex := int(blockStatesIndex)
	startBit := int((blockStatesIndex - math.Floor(blockStatesIndex)) * 64.0)
	if startBit+bits > 64 {
		prev := bitRange(s.BlockStates[longIndex], startBit, 64)
		next := bitRange(s.BlockStates[longIndex+1], 0, startBit+bits-64)
		nextBig := big.NewInt(next)
		startBitBig := big.NewInt(int64(startBit))
		nextBig.Lsh(nextBig, 64)
		nextBig.Sub(nextBig, startBitBig)
		return int(nextBig.Int64() + prev)
	}
	return int(bitRange(s.BlockStates[longIndex], startBit, startBit+bits))
}

// bitRange is used for extracting data from a sections blockstates array
func bitRange(value int64, from int, to int) int64 {
	waste := 64 - to
	return int64(uint64(value<<waste) >> (waste + from))
}
