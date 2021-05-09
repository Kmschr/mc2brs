package nbt

type StringNBT struct {
	name string
	data string
}

func (t *StringNBT) Type() byte {
	return TAG_String
}

func (t *StringNBT) Name() string {
	return t.name
}

func (t *StringNBT) Byte() byte {
	return 0
}

func (t *StringNBT) Short() int16 {
	return 0
}

func (t *StringNBT) Int() int32 {
	return 0
}

func (t *StringNBT) Long() int64 {
	return 0
}

func (t *StringNBT) Float() float32 {
	return 0
}

func (t *StringNBT) Double() float64 {
	return 0
}

func (t *StringNBT) String() string {
	return t.data
}

func (t *StringNBT) ByteArray() []byte {
	return nil
}

func (t *StringNBT) List() []Tag {
	return nil
}

func (t *StringNBT) Compound() map[string]Tag {
	return nil
}

func (t *StringNBT) IntArray() []int32 {
	return nil
}

func (t *StringNBT) LongArray() []int64 {
	return nil
}
