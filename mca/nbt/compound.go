package nbt

type CompoundNBT struct {
	name string
	data map[string]Tag
}

func (t *CompoundNBT) Type() byte {
	return TAG_Compound
}

func (t *CompoundNBT) Name() string {
	return t.name
}

func (t *CompoundNBT) Byte() byte {
	return 0
}

func (t *CompoundNBT) Short() int16 {
	return 0
}

func (t *CompoundNBT) Int() int32 {
	return 0
}

func (t *CompoundNBT) Long() int64 {
	return 0
}

func (t *CompoundNBT) Float() float32 {
	return 0
}

func (t *CompoundNBT) Double() float64 {
	return 0
}

func (t *CompoundNBT) String() string {
	return ""
}

func (t *CompoundNBT) ByteArray() []byte {
	return nil
}

func (t *CompoundNBT) List() []Tag {
	return nil
}

func (t *CompoundNBT) Compound() map[string]Tag {
	if t == nil {
		return nil
	}
	return t.data
}

func (t *CompoundNBT) IntArray() []int32 {
	return nil
}

func (t *CompoundNBT) LongArray() []int64 {
	return nil
}
