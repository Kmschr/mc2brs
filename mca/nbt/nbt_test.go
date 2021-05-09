package nbt_test

import (
	"bytes"
	"encoding/binary"
	"os"
	"testing"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/mca/nbt"
)

/*
func TestReadLevelDat(t *testing.T) {
	f, err := os.Open("level.dat")
	if err != nil {
		t.FailNow()
	}
	fmt.Println(nbt.ReadGzip(f, true))
}*/

func TestRead(t *testing.T) {
	regionFile, err := os.OpenFile("r.0.0.mca", os.O_RDONLY, 0444)
	if err != nil {
		t.Fatal(err.Error())
	}
	for z := 0; z < 32; z++ {
		for x := 0; x < 32; x++ {
			chunkOffsetBuf := make([]byte, 3)
			sectorCountBuf := make([]byte, 1)
			chunkHeaderByteOffset := int64(4 * ((x % 32) + (z%32)*32))
			regionFile.Seek(chunkHeaderByteOffset, 0)
			regionFile.Read(chunkOffsetBuf)
			regionFile.Read(sectorCountBuf)

			chunkOffsetBuf = append([]byte{0}, chunkOffsetBuf...)
			chunkOffset := int64(binary.BigEndian.Uint32(chunkOffsetBuf))
			sectorCount := int(sectorCountBuf[0])

			if sectorCount == 0 {
				continue
			}

			chunkLenBuf := make([]byte, 4)
			typeBuf := make([]byte, 1)
			regionFile.Seek(chunkOffset*4096, 0)
			regionFile.Read(chunkLenBuf)
			regionFile.Read(typeBuf)

			chunkLen := binary.BigEndian.Uint32(chunkLenBuf)

			if typeBuf[0] == 1 || typeBuf[0] == 3 {
				t.Fatal(ansi.Sprint(ansi.Red, "Chunks not in zlib format\n"))
			}

			chunkCompressedBuf := make([]byte, chunkLen-1)
			regionFile.Read(chunkCompressedBuf)
			b := bytes.NewReader(chunkCompressedBuf)
			_, err := nbt.ReadZlib(b, true)
			if err != nil {
				t.Fatal(err.Error())
			}
			break
		}
	}
}
