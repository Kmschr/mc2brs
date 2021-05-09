package nbt

type LongNBT struct {
	name string
	data int64
}

func (t *LongNBT) Type() byte {
	return TAG_Long
}

func (t *LongNBT) Name() string {
	return t.name
}

func (t *LongNBT) Byte() byte {
	return 0
}

func (t *LongNBT) Short() int16 {
	return 0
}

func (t *LongNBT) Int() int32 {
	return 0
}

func (t *LongNBT) Long() int64 {
	return t.data
}

func (t *LongNBT) Float() float32 {
	return 0
}

func (t *LongNBT) Double() float64 {
	return 0
}

func (t *LongNBT) String() string {
	return ""
}

func (t *LongNBT) ByteArray() []byte {
	return nil
}

func (t *LongNBT) List() []Tag {
	return nil
}

func (t *LongNBT) Compound() map[string]Tag {
	return nil
}

func (t *LongNBT) IntArray() []int32 {
	return nil
}

func (t *LongNBT) LongArray() []int64 {
	return nil
}
