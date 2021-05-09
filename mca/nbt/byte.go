package nbt

type ByteNBT struct {
	name string
	data byte
}

func (t *ByteNBT) Type() byte {
	return TAG_Byte
}

func (t *ByteNBT) Name() string {
	return t.name
}

func (t *ByteNBT) Byte() byte {
	return t.data
}

func (t *ByteNBT) Short() int16 {
	return 0
}

func (t *ByteNBT) Int() int32 {
	return 0
}

func (t *ByteNBT) Long() int64 {
	return 0
}

func (t *ByteNBT) Float() float32 {
	return 0
}

func (t *ByteNBT) Double() float64 {
	return 0
}

func (t *ByteNBT) String() string {
	return ""
}

func (t *ByteNBT) ByteArray() []byte {
	return nil
}

func (t *ByteNBT) List() []Tag {
	return nil
}

func (t *ByteNBT) Compound() map[string]Tag {
	return nil
}

func (t *ByteNBT) IntArray() []int32 {
	return nil
}

func (t *ByteNBT) LongArray() []int64 {
	return nil
}
