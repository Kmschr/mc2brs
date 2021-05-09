package brs

import (
	"image/color"
)

const (
	XPositive = iota
	XNegative
	YPositive
	YNegative
	ZPositive
	ZNegative

	Deg0 = iota
	Deg90
	Deg180
	Deg270
)

type Brick struct {
	AssetIndex    int
	Size          [3]int
	Pos           [3]int
	Direction     byte
	Rotation      byte
	Collision     bool
	Visibility    bool
	MaterialIndex int
	Intensity     int
	Color         color.RGBA
	Light         bool
	OwnerIndex    *int
	Name          string
}

// EqualTo returns true if a brick is matching another in
// size, orientation, and color
func (b Brick) EqualTo(o Brick) bool {
	return b.Size == o.Size &&
		b.AssetIndex == o.AssetIndex &&
		b.Direction == o.Direction &&
		b.Rotation == o.Rotation &&
		b.Color.R == o.Color.R &&
		b.Color.G == o.Color.G &&
		b.Color.B == o.Color.B
}

func (b Brick) EqualZAxisTo(o Brick) bool {
	return b.Size[0] == o.Size[0] &&
		b.Size[2] == o.Size[2] &&
		b.AssetIndex == o.AssetIndex &&
		b.Direction == o.Direction &&
		b.Rotation == o.Rotation &&
		b.Color.R == o.Color.R &&
		b.Color.G == o.Color.G &&
		b.Color.B == o.Color.B
}

func (b Brick) EqualXAxisTo(o Brick) bool {
	return b.Size[1] == o.Size[1] &&
		b.Size[2] == o.Size[2] &&
		b.AssetIndex == o.AssetIndex &&
		b.Direction == o.Direction &&
		b.Rotation == o.Rotation &&
		b.Color.R == o.Color.R &&
		b.Color.G == o.Color.G &&
		b.Color.B == o.Color.B
}
func (b Brick) EqualYAxisTo(o Brick) bool {
	return b.Size[0] == o.Size[0] &&
		b.Size[1] == o.Size[1] &&
		b.AssetIndex == o.AssetIndex &&
		b.Direction == o.Direction &&
		b.Rotation == o.Rotation &&
		b.Color.R == o.Color.R &&
		b.Color.G == o.Color.G &&
		b.Color.B == o.Color.B
}

func NewBrick(c color.RGBA, x, z, y int, name string) Brick {
	i := 0
	return Brick{
		AssetIndex:    0,
		Size:          [3]int{1, 1, 1},
		Pos:           [3]int{x, z, y},
		Direction:     ZPositive,
		Rotation:      Deg0,
		Collision:     true,
		Visibility:    true,
		MaterialIndex: 0,
		Intensity:     5,
		Color:         c,
		OwnerIndex:    &i,
		Name:          name,
	}
}

func NewWedge(c color.RGBA, x, z, y int, name string, facing string) Brick {
	var dir byte
	switch facing {
	case "north":
		dir = XNegative
	case "south":
		dir = XPositive
	case "east":
		dir = YNegative
	case "west":
		dir = YPositive
	}
	i := 0
	return Brick{
		AssetIndex:    1,
		Size:          [3]int{1, 1, 1},
		Pos:           [3]int{x, z, y},
		Direction:     dir,
		Rotation:      Deg180,
		Collision:     true,
		Visibility:    true,
		MaterialIndex: 0,
		Intensity:     5,
		Color:         c,
		OwnerIndex:    &i,
		Name:          name,
	}
}

func NewCorner(c color.RGBA, x, z, y int, name string, facing string) Brick {
	var rot byte
	switch facing {
	case "north":
		rot = Deg270
	case "south":
		rot = Deg90
	case "east":
		rot = Deg0
	case "west":
		rot = Deg180
	}
	i := 0
	return Brick{
		AssetIndex:    2,
		Size:          [3]int{1, 1, 1},
		Pos:           [3]int{x, z, y},
		Direction:     ZPositive,
		Rotation:      rot,
		Collision:     true,
		Visibility:    true,
		MaterialIndex: 0,
		Intensity:     5,
		Color:         c,
		OwnerIndex:    &i,
		Name:          name,
	}
}

func NewInnerCorner(c color.RGBA, x, z, y int, name string, facing string) Brick {
	var rot byte
	switch facing {
	case "north":
		rot = Deg270
	case "south":
		rot = Deg90
	case "east":
		rot = Deg0
	case "west":
		rot = Deg180
	}
	i := 0
	return Brick{
		AssetIndex:    3,
		Size:          [3]int{1, 1, 1},
		Pos:           [3]int{x, z, y},
		Direction:     ZPositive,
		Rotation:      rot,
		Collision:     true,
		Visibility:    true,
		MaterialIndex: 0,
		Intensity:     5,
		Color:         c,
		OwnerIndex:    &i,
		Name:          name,
	}
}
