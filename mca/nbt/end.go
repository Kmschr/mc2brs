package nbt

type EndNBT struct {
	name string
}

func (t *EndNBT) Type() byte {
	return TAG_End
}

func (t *EndNBT) Name() string {
	return t.name
}

func (t *EndNBT) Byte() byte {
	return 0
}

func (t *EndNBT) Short() int16 {
	return 0
}

func (t *EndNBT) Int() int32 {
	return 0
}

func (t *EndNBT) Long() int64 {
	return 0
}

func (t *EndNBT) Float() float32 {
	return 0
}

func (t *EndNBT) Double() float64 {
	return 0
}

func (t *EndNBT) String() string {
	return ""
}

func (t *EndNBT) ByteArray() []byte {
	return nil
}

func (t *EndNBT) List() []Tag {
	return nil
}

func (t *EndNBT) Compound() map[string]Tag {
	return nil
}

func (t *EndNBT) IntArray() []int32 {
	return nil
}

func (t *EndNBT) LongArray() []int64 {
	return nil
}
