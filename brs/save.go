package brs

import (
	"image"
	"time"

	"kmschr.com/mc2brs/util"
)

var smallguy = User{
	Name: "Smallguy",
	UUID: []byte{35, 235, 250, 142, 142, 66, 130, 94, 211, 13, 117, 181, 110, 20, 112, 2},
}

var minecraft = User{
	Name: "Minecraft",
	UUID: []byte{35, 235, 250, 142, 142, 66, 130, 94, 211, 13, 117, 181, 110, 20, 112, 2},
}

var zeblote = User{
	Name: "Zeblote",
	UUID: []byte{0x6D, 0x5D, 0x5C, 0xB7, 0xAB, 0x46, 0xBF, 0xFB, 0xA0, 0x47, 0x0E, 0x87, 0x63, 0x9E, 0x91, 0x16},
}

const (
	// Screenshot Format Enum
	screenshotFormatNone = iota
	screenshotFormatPNG
)

const (
	// Material Indices
	MaterialPlastic = iota
	MaterialGlass
	MaterialGlow
	MaterialHidden
	MaterialGhost
	MaterialGhostFail
	MaterialMetallic
	MaterialHologram
)

type User struct {
	Name string
	UUID []byte
}

type Save struct {
	Map          string
	Host         User
	Author       User
	Description  string
	SaveTime     time.Time
	Bricks       map[util.Vec3]Brick
	BrickAssets  []string
	Materials    []string
	BrickOwners  []User
	Screenshot   image.Image
	brickIndices []int
}

func NewSave() Save {
	s := Save{
		Map:         "Plate",
		Host:        zeblote,
		Author:      minecraft,
		Description: "Generated using mc2brs, a Minecraft to Brickadia save converter tool",
		BrickAssets: []string{"PB_DefaultMicroBrick", "PB_DefaultMicroWedge", "PB_DefaultMicroWedgeCorner", "PB_DefaultMicroWedgeInnerCorner"},
		Materials:   []string{"BMC_Plastic", "BMC_Glass", "BMC_Glow", "BMC_Hidden", "BMC_Ghost", "BMC_Ghost_Fail", "BMC_Metallic", "BMC_Hologram"},
		BrickOwners: []User{smallguy},
		SaveTime:    time.Now(),
	}
	return s
}
