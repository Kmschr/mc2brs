package brs

import "testing"

func TestWriteBit(t *testing.T) {
	b := NewBitWriter()
	b.WriteBit(false)
	b.WriteBit(true)
	b.WriteBit(true)
	b.WriteBit(false)
	b.WriteBit(false)
	b.WriteBit(false)
	b.WriteBit(false)
	b.WriteBit(false)
	t.Log(b.buf.Bytes())
	if b.buf.Bytes()[0] != 6 {
		t.Fail()
	}
}

func TestWriteIntMax(t *testing.T) {
	b := NewBitWriter()
	b.WriteIntMax(7, 512)
	b.FlushByte()
	t.Log(b.buf.Bytes())
}

func TestWriteIntPacked(t *testing.T) {
	b := NewBitWriter()
	b.WriteIntPacked(5402)
	b.FlushByte()
	t.Log(b.buf.Bytes())
}
