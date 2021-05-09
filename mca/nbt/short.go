package nbt

type ShortNBT struct {
	name string
	data int16
}

func (t *ShortNBT) Type() byte {
	return TAG_Short
}

func (t *ShortNBT) Name() string {
	return t.name
}

func (t *ShortNBT) Byte() byte {
	return 0
}

func (t *ShortNBT) Short() int16 {
	return t.data
}

func (t *ShortNBT) Int() int32 {
	return 0
}

func (t *ShortNBT) Long() int64 {
	return 0
}

func (t *ShortNBT) Float() float32 {
	return 0
}

func (t *ShortNBT) Double() float64 {
	return 0
}

func (t *ShortNBT) String() string {
	return ""
}

func (t *ShortNBT) ByteArray() []byte {
	return nil
}

func (t *ShortNBT) List() []Tag {
	return nil
}

func (t *ShortNBT) Compound() map[string]Tag {
	return nil
}

func (t *ShortNBT) IntArray() []int32 {
	return nil
}

func (t *ShortNBT) LongArray() []int64 {
	return nil
}
