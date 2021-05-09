package brs

import (
	"bytes"
	"image/png"
	"io"
	"os"
	"time"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/util"
)

var Magic = []byte{'B', 'R', 'S'}

const Version int16 = 10
const GameVersion int = 6781

var UE4DateTimeBase = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

func Write(save Save, name string) {
	w, _ := os.Create(name + ".brs")
	defer w.Close()

	writePreamble(save, w)
	writeHeader1(save, w)
	writeHeader2(save, w)
	writeScreenshot(save, w)
	writeBricks(&save, w)
	WriteComponents(save, w)
}

func writePreamble(save Save, w io.Writer) {
	b := NewWriter()
	b.Write(Magic)
	b.WriteShort(Version)
	b.WriteInt(GameVersion)
	b.WriteTo(w)
}

func writeHeader1(save Save, w io.Writer) {
	b := NewWriter()
	b.WriteString(save.Map)
	b.WriteString(save.Author.Name)
	b.WriteString(save.Description)
	b.Write(save.Author.UUID)
	b.WriteUserNameFirst(save.Host)
	b.WriteDateTime(save.SaveTime)
	b.WriteInt(len(save.Bricks))
	b.WriteNotCompressed(w)
}

func writeHeader2(save Save, w io.Writer) {
	b := NewWriter()
	b.WriteStrings(nil) // mods
	b.WriteStrings(save.BrickAssets)
	b.WriteInt(0) // color palette
	b.WriteStrings(save.Materials)
	b.WriteInt(len(save.BrickOwners))
	for _, user := range save.BrickOwners {
		b.Write(user.UUID)
		b.WriteString(user.Name)
		b.WriteInt(len(save.Bricks))
	}
	b.WriteStrings([]string{"BPMC_Default"}) // physical mats
	b.WriteNotCompressed(w)
}

func writeScreenshot(save Save, w io.Writer) {
	b := NewWriter()
	if save.Screenshot == nil {
		b.WriteByte(screenshotFormatNone)
	} else {
		b.WriteByte(screenshotFormatPNG)
		screenshotBuf := bytes.NewBuffer([]byte{})
		err := png.Encode(screenshotBuf, save.Screenshot)
		if err != nil {
			ansi.Println(ansi.Red, err.Error())
			ansi.Quit()
		}
		b.WriteInt(screenshotBuf.Len())
		b.Write(screenshotBuf.Bytes())
	}
	b.WriteTo(w)
}

func writeBricks(save *Save, w io.Writer) {
	b := NewBitWriter()
	maxBrickAssetIndex := util.Max(len(save.BrickAssets), 2)
	maxMaterialIndex := util.Max(len(save.Materials), 2)
	total := len(save.Bricks)
	i := 0
	ansi.Print(ansi.BrightBlue, "[")
	for _, brick := range save.Bricks {
		if i%(total/16) == 0 {
			ansi.Print(ansi.BrightBlue, "#")
		}
		b.FlushByte()
		b.WriteIntMax(brick.AssetIndex, maxBrickAssetIndex)
		if b.WriteBit(!isZero(brick.Size)) {
			b.WritePositiveIntVectorPacked(brick.Size)
		}
		b.WriteIntVectorPacked(brick.Pos)
		orientation := combineOrientation(brick.Direction, brick.Rotation)
		b.WriteIntMax(int(orientation), 24)
		b.WriteBit(brick.Collision) // Player Collision
		b.WriteBit(brick.Collision) // Weapon Collision
		b.WriteBit(true)            // Interaction Collision
		b.WriteBit(true)            // Tool Collision
		b.WriteBit(brick.Visibility)
		b.WriteIntMax(brick.MaterialIndex, maxMaterialIndex)
		b.WriteIntMax(0, 2)                // Physical Index
		b.WriteIntMax(brick.Intensity, 11) // Material Intensity
		b.WriteBit(true)                   // use non palette color
		b.WriteBytes(colorBytes(brick.Color))
		var ownerIndex int
		if brick.OwnerIndex != nil {
			ownerIndex = *brick.OwnerIndex + 1
		}
		b.WriteIntPacked(ownerIndex)
		if brick.Light {
			save.brickIndices = append(save.brickIndices, i)
		}
		i++
	}
	b.FlushByte()
	b.WriteCompressed(w)
	ansi.Println(ansi.BrightBlue, "]")
}

func WriteComponents(save Save, w io.Writer) {
	b := NewWriter()
	maxBrickIndex := util.Max(len(save.Bricks), 2)
	b.WriteInt(1)                   // Num components
	b.WriteString("BCD_PointLight") // Component Type
	bits := NewBitWriter()
	bits.WriteInt(1) // version

	// brick indices
	bits.WriteInt(len(save.brickIndices))
	for _, i := range save.brickIndices {
		bits.WriteIntMax(i, maxBrickIndex)
	}

	// list of name and type properties
	bits.WriteInt(6)
	bits.WriteString("Brightness")
	bits.WriteString("Float")
	bits.WriteString("Color")
	bits.WriteString("Color")
	bits.WriteString("Radius")
	bits.WriteString("Float")
	bits.WriteString("bCastShadows")
	bits.WriteString("Boolean")
	bits.WriteString("bMatchBrickShape")
	bits.WriteString("Boolean")
	bits.WriteString("bUseBrickColor")
	bits.WriteString("Boolean")

	scale := float32(ansi.Scale())
	brightness := 1.06*scale*scale + 3.04*scale
	radius := 25 * float32(scale)

	// components for each brick
	for i := 0; i < len(save.brickIndices); i++ {
		bits.WriteFloat(brightness)         // Brightness
		bits.WriteBytes([]byte{0, 0, 0, 0}) // Color, unused
		bits.WriteFloat(radius)             // Radius
		bits.WriteInt(0)                    // Cast Shadows (False)
		bits.WriteInt(0)                    // Match Brick Shape (False)
		bits.WriteInt(1)                    // Use brick color (True)
	}

	bits.FlushByte()
	bits.PrependLength()
	bits.WriteTo(b)

	b.WriteCompressed(w)
}
