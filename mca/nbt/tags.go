package nbt

import (
	"fmt"

	"kmschr.com/mc2brs/ansi"
)

// https://minecraft.fandom.com/wiki/NBT_format
const (
	TAG_End = iota
	TAG_Byte
	TAG_Short
	TAG_Int
	TAG_Long
	TAG_Float
	TAG_Double
	TAG_Byte_Array
	TAG_String
	TAG_List
	TAG_Compound
	TAG_Int_Array
	TAG_Long_Array
)

var tagStrings = [...]string{
	"End",
	"Byte",
	"Short",
	"Int",
	"Long",
	"Float",
	"Double",
	"ByteArray",
	"String",
	"List",
	"CompoundTag",
	"IntArray",
	"LongArray",
}

type Tag interface {
	Type() byte
	Name() string
	Byte() byte
	Short() int16
	Int() int32
	Long() int64
	Float() float32
	Double() float64
	String() string
	ByteArray() []byte
	List() []Tag
	Compound() map[string]Tag
	IntArray() []int32
	LongArray() []int64
}

func printTagType(t byte) {
	if int(t) > len(tagStrings) {
		ansi.Println(ansi.Red, "Invalid tag type")
		return
	}
	if t == TAG_End {
		return
	}
	fmt.Printf("%s> ", tagStrings[t])
}

func printTagDepth(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
}

func printTagName(n string, t byte) {
	if t == TAG_End {
		ansi.Print(ansi.BrightRed, "<End>")
		return
	}
	if n == "" {
		n = "root"
	}
	ansi.Print(ansi.Yellow, fmt.Sprintf("<%s: ", n))
}
