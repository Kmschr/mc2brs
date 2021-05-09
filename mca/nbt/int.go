package nbt

type IntNBT struct {
	name string
	data int32
}

func (t *IntNBT) Type() byte {
	return TAG_Int
}

func (t *IntNBT) Name() string {
	return t.name
}

func (t *IntNBT) Byte() byte {
	return 0
}

func (t *IntNBT) Short() int16 {
	return 0
}

func (t *IntNBT) Int() int32 {
	return t.data
}

func (t *IntNBT) Long() int64 {
	return 0
}

func (t *IntNBT) Float() float32 {
	return 0
}

func (t *IntNBT) Double() float64 {
	return 0
}

func (t *IntNBT) String() string {
	return ""
}

func (t *IntNBT) ByteArray() []byte {
	return nil
}

func (t *IntNBT) List() []Tag {
	return nil
}

func (t *IntNBT) Compound() map[string]Tag {
	return nil
}

func (t *IntNBT) IntArray() []int32 {
	return nil
}

func (t *IntNBT) LongArray() []int64 {
	return nil
}
