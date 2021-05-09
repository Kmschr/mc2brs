package nbt

type IntArrayNBT struct {
	name string
	data []int32
}

func (t *IntArrayNBT) Type() byte {
	return TAG_Int_Array
}

func (t *IntArrayNBT) Name() string {
	return t.name
}

func (t *IntArrayNBT) Byte() byte {
	return 0
}

func (t *IntArrayNBT) Short() int16 {
	return 0
}

func (t *IntArrayNBT) Int() int32 {
	return 0
}

func (t *IntArrayNBT) Long() int64 {
	return 0
}

func (t *IntArrayNBT) Float() float32 {
	return 0
}

func (t *IntArrayNBT) Double() float64 {
	return 0
}

func (t *IntArrayNBT) String() string {
	return ""
}

func (t *IntArrayNBT) ByteArray() []byte {
	return nil
}

func (t *IntArrayNBT) List() []Tag {
	return nil
}

func (t *IntArrayNBT) Compound() map[string]Tag {
	return nil
}

func (t *IntArrayNBT) IntArray() []int32 {
	return t.data
}

func (t *IntArrayNBT) LongArray() []int64 {
	return nil
}
