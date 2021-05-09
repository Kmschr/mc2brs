package mca

import "kmschr.com/mc2brs/mca/nbt"

// BlockState contains the name and properties of a block
type BlockState struct {
	Name       string
	Properties map[string]string
}

// newBlock creates a block state from it's assumed to be valid root tag
func newBlock(t nbt.Tag) BlockState {
	block := BlockState{}
	block.Name = t.Compound()["Name"].String()
	block.parseProperties(t)
	return block
}

// parseProperties loads the properties from the compound tag into a map
func (b *BlockState) parseProperties(t nbt.Tag) {
	propertiesNBT := t.Compound()["Properties"]
	if propertiesNBT == nil {
		return
	}
	b.Properties = make(map[string]string)
	properties := propertiesNBT.Compound()
	for _, propertyTag := range properties {
		b.Properties[propertyTag.Name()] = propertyTag.String()
	}
}
