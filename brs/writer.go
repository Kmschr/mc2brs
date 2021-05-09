package brs

import (
	"bytes"
	"compress/zlib"
	"io"
	"time"
)

type Writer struct {
	b *bytes.Buffer
}

func NewWriter() Writer {
	b := bytes.NewBuffer([]byte{})
	return Writer{b}
}

func (w *Writer) WriteTo(t io.Writer) (int64, error) {
	return w.b.WriteTo(t)
}

func (w *Writer) WriteCompressed(t io.Writer) {
	internal := NewWriter()
	internal.WriteInt(w.b.Len())
	internalCompressed := bytes.NewBuffer([]byte{})
	z, _ := zlib.NewWriterLevel(internalCompressed, zlib.BestCompression)
	_, _ = z.Write(w.b.Bytes())
	z.Close()
	internal.WriteInt(internalCompressed.Len())
	internal.Write(internalCompressed.Bytes())
	internal.WriteTo(t)
}

func (w *Writer) WriteNotCompressed(t io.Writer) {
	internal := NewWriter()
	internal.WriteInt(w.b.Len())
	internal.WriteInt(0)
	internal.Write(w.b.Bytes())
	internal.WriteTo(t)
}

func (w *Writer) Write(b []byte) {
	w.b.Write(b)
}

func (w *Writer) WriteByte(b byte) error {
	return w.b.WriteByte(b)
}

func (w *Writer) WriteShort(s int16) {
	b := []byte{byte(s), byte(s >> 8)}
	w.b.Write(b)
}

func (w *Writer) WriteInt(i int) {
	b := []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)}
	w.b.Write(b)
}

func (w *Writer) WriteString(s string) {
	w.WriteInt(len(s) + 1)
	b := []byte(s)
	b = append(b, 0)
	w.b.Write(b)
}

func (w *Writer) WriteLong(l int64) {
	b := []byte{byte(l), byte(l >> 8), byte(l >> 16), byte(l >> 24), byte(l >> 32), byte(l >> 40), byte(l >> 48), byte(l >> 56)}
	w.b.Write(b)
}

func (w *Writer) WriteUserNameFirst(u User) {
	w.WriteString(u.Name)
	w.Write(u.UUID)
}

func (w *Writer) WriteDateTime(t time.Time) {
	w.WriteLong(0)
}

func (w *Writer) WriteStrings(strings []string) {
	if strings == nil {
		w.WriteInt(0)
		return
	}
	w.WriteInt(len(strings))
	for _, s := range strings {
		w.WriteString(s)
	}
}
