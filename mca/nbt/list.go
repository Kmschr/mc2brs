package nbt

type ListNBT struct {
	name string
	data []Tag
}

func (t *ListNBT) Type() byte {
	return TAG_List
}

func (t *ListNBT) Name() string {
	return t.name
}

func (t *ListNBT) Byte() byte {
	return 0
}

func (t *ListNBT) Short() int16 {
	return 0
}

func (t *ListNBT) Int() int32 {
	return 0
}

func (t *ListNBT) Long() int64 {
	return 0
}

func (t *ListNBT) Float() float32 {
	return 0
}

func (t *ListNBT) Double() float64 {
	return 0
}

func (t *ListNBT) String() string {
	return ""
}

func (t *ListNBT) ByteArray() []byte {
	return nil
}

func (t *ListNBT) List() []Tag {
	return t.data
}

func (t *ListNBT) Compound() map[string]Tag {
	return nil
}

func (t *ListNBT) IntArray() []int32 {
	return nil
}

func (t *ListNBT) LongArray() []int64 {
	return nil
}
