package brs

import (
	"bytes"
	"compress/zlib"
	"io"
	"math"
)

type BitWriter struct {
	buf *bytes.Buffer
	cur byte
	bit byte
}

func NewBitWriter() BitWriter {
	return BitWriter{
		buf: bytes.NewBuffer([]byte{}),
	}
}

func (b *BitWriter) WriteCompressed(w io.Writer) {
	internal := NewWriter()
	internal.WriteInt(b.buf.Len())
	internalCompressed := bytes.NewBuffer([]byte{})
	z, _ := zlib.NewWriterLevel(internalCompressed, zlib.BestCompression)
	_, _ = z.Write(b.buf.Bytes())
	z.Close()
	internal.WriteInt(internalCompressed.Len())
	internal.Write(internalCompressed.Bytes())
	internal.WriteTo(w)
}

func (b *BitWriter) WriteTo(w Writer) {
	w.Write(b.buf.Bytes())
}

func (b *BitWriter) WriteBit(bitBool bool) bool {
	var bit byte
	if bitBool {
		bit = 1
	}
	b.cur |= bit << b.bit
	b.bit++
	if b.bit >= 8 {
		b.FlushByte()
	}
	return bitBool
}

func (b *BitWriter) WriteBits(src int, len int) {
	for bit := 0; bit < len; bit++ {
		mask := 1 << (bit & 7)
		b.WriteBit(src&mask != 0)
	}
}

func (b *BitWriter) WriteBytes(src []byte) {
	b.WriteBitArray(src, len(src)*8)
}

func (b *BitWriter) WriteBitArray(src []byte, len int) {
	for bit := 0; bit < len; bit++ {
		bit := (src[bit>>3]&0xFF)&(1<<(bit&7)) != 0
		b.WriteBit(bit)
	}
}

func (b *BitWriter) WriteInt(i int) {
	w := NewWriter()
	w.WriteInt(i)
	buf := w.b.Bytes()
	b.WriteBitArray(buf, w.b.Len()*8)
}

func (b *BitWriter) WriteIntMax(val int, max int) {
	newVal := 0
	mask := 1
	for newVal+mask < max && mask != 0 {
		b.WriteBit((val & mask) != 0)
		if val&mask != 0 {
			newVal |= mask
		}
		mask <<= 1
	}
}

func (b *BitWriter) WriteIntPacked(val int) {
	for {
		src := val & 0x7F
		val = int(uint(val) >> 7)
		b.WriteBit(val != 0)
		b.WriteBits(src, 7)
		if val == 0 {
			break
		}
	}
}

func (b *BitWriter) WriteFloat(f float32) {
	i := math.Float32bits(f)
	b.WriteInt(int(i))
}

func (b *BitWriter) WriteString(s string) {
	b.WriteInt(len(s) + 1)
	buf := []byte(s)
	buf = append(buf, 0)
	b.WriteBytes(buf)
}

func (b *BitWriter) PrependLength() {
	buf := b.buf.Bytes()
	newBit := NewBitWriter()
	newBit.WriteInt(len(buf))
	newBit.WriteBytes(buf)
	b.buf = newBit.buf
}

func (b *BitWriter) WritePositiveIntVectorPacked(v [3]int) {
	b.WriteIntPacked(v[0])
	b.WriteIntPacked(v[1])
	b.WriteIntPacked(v[2])
}

func (b *BitWriter) WriteIntVectorPacked(v [3]int) {
	b.WriteIntPacked(mapInt(v[0]))
	b.WriteIntPacked(mapInt(v[1]))
	b.WriteIntPacked(mapInt(v[2]))
}

func mapInt(n int) int {
	var sign int
	if n > 0 {
		sign = 1
	} else {
		n = -n
	}

	return n<<1 | sign
}

func (b *BitWriter) FlushByte() {
	if b.bit > 0 {
		b.buf.WriteByte(b.cur)
		b.cur = 0
		b.bit = 0
	}
}
