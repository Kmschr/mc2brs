package nbt

type LongArrayNBT struct {
	name string
	data []int64
}

func (t *LongArrayNBT) Type() byte {
	return TAG_Long_Array
}

func (t *LongArrayNBT) Name() string {
	return t.name
}

func (t *LongArrayNBT) Byte() byte {
	return 0
}

func (t *LongArrayNBT) Short() int16 {
	return 0
}

func (t *LongArrayNBT) Int() int32 {
	return 0
}

func (t *LongArrayNBT) Long() int64 {
	return 0
}

func (t *LongArrayNBT) Float() float32 {
	return 0
}

func (t *LongArrayNBT) Double() float64 {
	return 0
}

func (t *LongArrayNBT) String() string {
	return ""
}

func (t *LongArrayNBT) ByteArray() []byte {
	return nil
}

func (t *LongArrayNBT) List() []Tag {
	return nil
}

func (t *LongArrayNBT) Compound() map[string]Tag {
	return nil
}

func (t *LongArrayNBT) IntArray() []int32 {
	return nil
}

func (t *LongArrayNBT) LongArray() []int64 {
	return t.data
}
