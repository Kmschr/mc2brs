package nbt

type FloatNBT struct {
	name string
	data float32
}

func (t *FloatNBT) Type() byte {
	return TAG_Float
}

func (t *FloatNBT) Name() string {
	return t.name
}

func (t *FloatNBT) Byte() byte {
	return 0
}

func (t *FloatNBT) Short() int16 {
	return 0
}

func (t *FloatNBT) Int() int32 {
	return 0
}

func (t *FloatNBT) Long() int64 {
	return 0
}

func (t *FloatNBT) Float() float32 {
	return t.data
}

func (t *FloatNBT) Double() float64 {
	return 0
}

func (t *FloatNBT) String() string {
	return ""
}

func (t *FloatNBT) ByteArray() []byte {
	return nil
}

func (t *FloatNBT) List() []Tag {
	return nil
}

func (t *FloatNBT) Compound() map[string]Tag {
	return nil
}

func (t *FloatNBT) IntArray() []int32 {
	return nil
}

func (t *FloatNBT) LongArray() []int64 {
	return nil
}
