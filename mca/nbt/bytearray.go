package nbt

type ByteArrayNBT struct {
	name string
	data []byte
}

func (t *ByteArrayNBT) Type() byte {
	return TAG_Byte_Array
}

func (t *ByteArrayNBT) Name() string {
	return t.name
}

func (t *ByteArrayNBT) Byte() byte {
	return 0
}

func (t *ByteArrayNBT) Short() int16 {
	return 0
}

func (t *ByteArrayNBT) Int() int32 {
	return 0
}

func (t *ByteArrayNBT) Long() int64 {
	return 0
}

func (t *ByteArrayNBT) Float() float32 {
	return 0
}

func (t *ByteArrayNBT) Double() float64 {
	return 0
}

func (t *ByteArrayNBT) String() string {
	return ""
}

func (t *ByteArrayNBT) ByteArray() []byte {
	return t.data
}

func (t *ByteArrayNBT) List() []Tag {
	return nil
}

func (t *ByteArrayNBT) Compound() map[string]Tag {
	return nil
}

func (t *ByteArrayNBT) IntArray() []int32 {
	return nil
}

func (t *ByteArrayNBT) LongArray() []int64 {
	return nil
}
