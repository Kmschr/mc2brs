package mca

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/mca/nbt"
)

const regionWidth = 32 * 16
const chunkSectorSize = 4096

type RegionInfo struct {
	Path      string
	FileInfo  os.FileInfo
	RegionX   int
	RegionZ   int
	WorldName string
}

// Region is a 32x32 area of chunks
type Region struct {
	RegionX    int
	RegionZ    int
	Chunks     map[int]Chunk
	ChunkTable map[int]int64
	File       *os.File
}

// NewRegion creates a region given a specified path, loading the
// positions for compressed buffers for chunks but not reading them
func NewRegion(path string, rx int, rz int) Region {
	region := Region{
		Chunks:     make(map[int]Chunk),
		ChunkTable: make(map[int]int64),
	}
	region.RegionX = rx
	region.RegionZ = rz
	regionFile, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		ansi.Println(ansi.Red, "Error opening regionFile for reading\n"+err.Error())
		ansi.Quit()
	}
	for z := 0; z < 32; z++ {
		for x := 0; x < 32; x++ {
			chunkOffset, sectorCount := chunkHeader(regionFile, x, z)
			if sectorCount == 0 || chunkOffset == 0 {
				continue
			}
			region.ChunkTable[chunksIndex(x, z)] = chunkOffset * chunkSectorSize
		}
	}
	region.File = regionFile
	return region
}

// LoadAll loads every available chunk in the region.
func (r *Region) LoadAll() {
	for z := 0; z < 32; z++ {
		for x := 0; x < 32; x++ {
			if !r.ChunkExists(x, z) {
				continue
			}
			r.LoadChunk(x, z)
		}
	}
}

// LoadChunk loads the chunk at a specified index.
func (r *Region) LoadChunk(chunkX int, chunkZ int) {
	chunkIndex := chunksIndex(chunkX, chunkZ)
	chunkOffset, exists := r.ChunkTable[chunkIndex]
	if !exists || chunkOffset == 0 {
		return
	}
	chunkLenBuf := make([]byte, 4)
	typeBuf := make([]byte, 1)
	r.File.Seek(chunkOffset, 0)
	r.File.Read(chunkLenBuf)
	r.File.Read(typeBuf)
	chunkLen := binary.BigEndian.Uint32(chunkLenBuf)
	if typeBuf[0] != 2 {
		if typeBuf[0] == 0 {
			return
		}
		ansi.Println(ansi.Red, "Chunks not in zlib format")
		fmt.Print(typeBuf[0])
		ansi.Quit()
	}
	chunkCompressedBuf := make([]byte, chunkLen-1)
	r.File.Read(chunkCompressedBuf)
	b := bytes.NewReader(chunkCompressedBuf)
	chunkNBT, err := nbt.ReadZlib(b, false)
	if err != nil {
		ansi.Quit()
	}
	chunk := newChunk(chunkNBT)
	r.Chunks[chunkIndex] = chunk
}

// ChunkLoaded returns whether or not a chunk is loaded within the region
func (r *Region) ChunkLoaded(chunkX int, chunkZ int) bool {
	_, loaded := r.Chunks[chunksIndex(chunkX, chunkZ)]
	return loaded
}

// ChunkExists returns whether or not a chunk exists within the region
func (r *Region) ChunkExists(chunkX int, chunkZ int) bool {
	_, exists := r.ChunkTable[chunksIndex(chunkX, chunkZ)]
	return exists
}

// ChunkAt returns the chunk at a chunk coordinate within the region
func (r *Region) ChunkAt(chunkX int, chunkZ int) (Chunk, bool) {
	if r.Chunks == nil {
		return Chunk{}, false
	}
	chunk, exists := r.Chunks[chunksIndex(chunkX, chunkZ)]
	return chunk, exists
}

// BiomeAt returns the biome at a point within the region
func (r *Region) BiomeAt(x, z, y int) int32 {
	chunkIndex := chunksIndex(chunkCoords(x, z))
	chunk, exists := r.Chunks[chunkIndex]
	if !exists {
		return -1
	}
	return chunk.biomeAt(x, z, y)
}

// BlockAt returns the blockstate at a point within the region
func (r *Region) BlockAt(x, z, y int) (BlockState, bool) {
	i := chunksIndex(chunkCoords(x, z))
	chunk, exists := r.Chunks[i]
	if !exists {
		return BlockState{}, false
	}
	return chunk.BlockAt(x, z, y)
}

// SurroundingAt returns up to 6 block with adjacent faces to a block
func (r *Region) SurroundingAt(x, z, y int) []BlockState {
	var blocks []BlockState
	indices := [...][3]int{
		{x, z, y + 1},
		{x, z, y - 1},
		{x, z + 1, y},
		{x, z - 1, y},
		{x + 1, z, y},
		{x - 1, z, y},
	}
	for _, i := range indices {
		if !inBounds(i[0], i[1], i[2]) {
			continue
		}
		b, exists := r.BlockAt(i[0], i[1], i[2])
		if !exists {
			continue
		}
		blocks = append(blocks, b)
	}
	return blocks
}

// inBounds checks if a block is within the region
func inBounds(x, z, y int) bool {
	return x >= 0 && x < regionWidth &&
		z >= 0 && z < regionWidth &&
		y >= 0 && y < 256
}

// chunksIndex returns the index of a chunk within a regions chunks array
func chunksIndex(chunkX int, chunkZ int) int {
	return (chunkZ&0x1F)<<5 | chunkX&0x1F
}

// chunkCoords returns the chunk coordinates of a block within a region
func chunkCoords(x int, z int) (int, int) {
	return x >> 4, z >> 4
}

// chunkHeader reads the header for a chunk within a region file
// and returns the chunk's offset into the regionfile as well as
// the number of sectors the chunk fills.
func chunkHeader(f *os.File, x int, z int) (int64, int) {
	chunkOffsetBuf := make([]byte, 3)
	sectorCountBuf := make([]byte, 1)
	chunkHeaderByteOffset := chunkHeaderLocation(x, z)
	f.Seek(chunkHeaderByteOffset, 0)
	f.Read(chunkOffsetBuf)
	f.Read(sectorCountBuf)
	chunkOffsetBuf = append([]byte{0}, chunkOffsetBuf...)
	chunkOffset := int64(binary.BigEndian.Uint32(chunkOffsetBuf))
	sectorCount := int(sectorCountBuf[0])
	return chunkOffset, sectorCount
}

// chunkHeaderLocation returns the location of a chunk's
// header within a region file relative to the origin.
func chunkHeaderLocation(x int, z int) int64 {
	return int64(4 * ((x % 32) + (z%32)*32))
}
