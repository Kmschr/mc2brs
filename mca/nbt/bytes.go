package nbt

import (
	"encoding/binary"
	"io"
	"math"
)

func readByte(r io.Reader) byte {
	buf := make([]byte, 1)
	io.ReadFull(r, buf)
	return buf[0]
}

func readShort(r io.Reader) int16 {
	buf := make([]byte, 2)
	io.ReadFull(r, buf)
	return int16(buf[0])<<8 | int16(buf[1])
}

func readUInt16(r io.Reader) uint16 {
	buf := make([]byte, 2)
	io.ReadFull(r, buf)
	return uint16(buf[0])<<8 | uint16(buf[1])
}

func readInt32(r io.Reader) int32 {
	buf := make([]byte, 4)
	io.ReadFull(r, buf)
	return int32(binary.BigEndian.Uint32(buf))
}

func readInt64(r io.Reader) int64 {
	buf := make([]byte, 8)
	io.ReadFull(r, buf)
	return int64(binary.BigEndian.Uint64(buf))
}

func readFloat(r io.Reader) float32 {
	buf := make([]byte, 4)
	io.ReadFull(r, buf)
	return math.Float32frombits(binary.BigEndian.Uint32(buf))
}

func readDouble(r io.Reader) float64 {
	buf := make([]byte, 8)
	io.ReadFull(r, buf)
	return math.Float64frombits(binary.BigEndian.Uint64(buf))
}

func readString(r io.Reader) string {
	buf := make([]byte, readUInt16(r))
	io.ReadFull(r, buf)
	return string(buf)
}

func readByteArray(r io.Reader) []byte {
	buf := make([]byte, readInt32(r))
	io.ReadFull(r, buf)
	return buf
}

func readIntArray(r io.Reader) []int32 {
	size := readInt32(r)
	buf := make([]byte, size*4)
	io.ReadFull(r, buf)
	vals := make([]int32, size)
	for i := 0; i < int(size); i++ {
		bytes := buf[i*4 : (i*4 + 4)]
		vals[i] = int32(binary.BigEndian.Uint32(bytes))
	}
	return vals
}

func readLongArray(r io.Reader) []int64 {
	size := readInt32(r)
	buf := make([]byte, size*8)
	io.ReadFull(r, buf)
	vals := make([]int64, size)
	for i := 0; i < int(size); i++ {
		bytes := buf[i*8 : (i*8 + 8)]
		vals[i] = int64(binary.BigEndian.Uint64(bytes))
	}
	return vals
}
