package nbt

type DoubleNBT struct {
	name string
	data float64
}

func (t *DoubleNBT) Type() byte {
	return TAG_Double
}

func (t *DoubleNBT) Name() string {
	return t.name
}

func (t *DoubleNBT) Byte() byte {
	return 0
}

func (t *DoubleNBT) Short() int16 {
	return 0
}

func (t *DoubleNBT) Int() int32 {
	return 0
}

func (t *DoubleNBT) Long() int64 {
	return 0
}

func (t *DoubleNBT) Float() float32 {
	return 0
}

func (t *DoubleNBT) Double() float64 {
	return t.data
}

func (t *DoubleNBT) String() string {
	return ""
}

func (t *DoubleNBT) ByteArray() []byte {
	return nil
}

func (t *DoubleNBT) List() []Tag {
	return nil
}

func (t *DoubleNBT) Compound() map[string]Tag {
	return nil
}

func (t *DoubleNBT) IntArray() []int32 {
	return nil
}

func (t *DoubleNBT) LongArray() []int64 {
	return nil
}
