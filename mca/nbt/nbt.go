package nbt

import (
	"bufio"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"os"

	"kmschr.com/mc2brs/ansi"
)

type tagTreeInfo struct {
	r         io.Reader
	depth     int
	tagType   byte
	listIndex int
	debug     bool
}

func ReadZlib(r io.Reader, debug bool) (Tag, error) {
	z, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer z.Close()

	root := readTag(newTagTree(z, debug))
	return root, nil
}

func ReadGzip(r io.Reader, debug bool) (Tag, error) {
	g, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer g.Close()

	root := readTag(newTagTree(g, debug))
	return root, nil
}

func newTagTree(r io.Reader, debug bool) tagTreeInfo {
	return tagTreeInfo{
		r:         r,
		depth:     0,
		tagType:   0,
		listIndex: -1,
		debug:     debug,
	}
}

func readTag(treeInfo tagTreeInfo) Tag {
	var tagType byte
	var name string

	if treeInfo.tagType != 0 {
		tagType = treeInfo.tagType
		name = fmt.Sprintf("%d", treeInfo.listIndex)
	} else {
		tagType = readByte(treeInfo.r)
		if tagType != TAG_End {
			name = readString(treeInfo.r)
		}
	}

	var tag Tag

	if treeInfo.debug {
		fmt.Println()
		printTagDepth(treeInfo.depth)
		printTagName(name, tagType)
		printTagType(tagType)
		f := bufio.NewWriter(os.Stdout)
		f.Flush()
	}

	switch tagType {
	case TAG_End:
		tag = &EndNBT{"End"}
	case TAG_Byte:
		tag = &ByteNBT{name, readByte(treeInfo.r)}
	case TAG_Short:
		tag = &ShortNBT{name, readShort(treeInfo.r)}
	case TAG_Int:
		tag = &IntNBT{name, readInt32(treeInfo.r)}
	case TAG_Long:
		tag = &LongNBT{name, readInt64(treeInfo.r)}
	case TAG_Float:
		tag = &FloatNBT{name, readFloat(treeInfo.r)}
	case TAG_Double:
		tag = &DoubleNBT{name, readDouble(treeInfo.r)}
	case TAG_String:
		tag = &StringNBT{name, readString(treeInfo.r)}
	case TAG_Byte_Array:
		tag = &ByteArrayNBT{name, readByteArray(treeInfo.r)}
	case TAG_List:
		var list []Tag
		treeInfo.tagType = readByte(treeInfo.r)
		size := readInt32(treeInfo.r)
		treeInfo.depth++
		for i := 0; i < int(size); i++ {
			treeInfo.listIndex = i
			list = append(list, readTag(treeInfo))
		}
		tag = &ListNBT{name, list}
	case TAG_Compound:
		children := make(map[string]Tag)
		treeInfo.depth++
		treeInfo.tagType = 0
		treeInfo.listIndex = -1
		for {
			child := readTag(treeInfo)
			if child.Type() == TAG_End {
				break
			}
			children[child.Name()] = child
		}
		tag = &CompoundNBT{name, children}
	case TAG_Int_Array:
		tag = &IntArrayNBT{name, readIntArray(treeInfo.r)}
	case TAG_Long_Array:
		tag = &LongArrayNBT{name, readLongArray(treeInfo.r)}
	default:
		ansi.Println(ansi.Red, "Unknown tag")
		ansi.Quit()
	}

	if treeInfo.debug {
		switch tagType {
		case TAG_Byte:
			ansi.Print(ansi.BrightBlue, fmt.Sprint(tag.Byte()))
		case TAG_Short:
			ansi.Print(ansi.BrightBlue, fmt.Sprint(tag.Short()))
		case TAG_Int:
			ansi.Print(ansi.BrightBlue, fmt.Sprint(tag.Int()))
		case TAG_Long:
			ansi.Print(ansi.BrightBlue, fmt.Sprint(tag.Long()))
		case TAG_String:
			ansi.Print(ansi.BrightBlue, tag.String())
		case TAG_Int_Array:
			ansi.Print(ansi.BrightBlue, fmt.Sprint(len(tag.IntArray())))
		case TAG_Long_Array:
			ansi.Print(ansi.BrightBlue, fmt.Sprint(len(tag.LongArray())))
		}
	}

	return tag
}
