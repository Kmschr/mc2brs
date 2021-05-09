package mca

import (
	"fmt"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/brs"
	"kmschr.com/mc2brs/util"
)

// optimize merges bricks with their neighbors of the same size
func optimize(bricks map[util.Vec3]brs.Brick) {
	ansi.Print(ansi.BrightYellow, "Optimizing")
	ansi.Print(ansi.BrightBlue, "      [")
	total := len(bricks)
	i := 0
	for pos, brick := range bricks {
		if i%(total/16) == 0 {
			ansi.Print(ansi.BrightBlue, "#")
		}

		if brick.AssetIndex != 0 {
			i++
			continue
		}

		above := util.Vec3{X: pos.X, Z: pos.Z, Y: pos.Y + brick.Size[2]*2}
		east := util.Vec3{X: pos.X + brick.Size[0]*2, Z: pos.Z, Y: pos.Y}
		south := util.Vec3{X: pos.X, Z: pos.Z + brick.Size[1]*2, Y: pos.Y}
		southeast := util.Vec3{X: pos.X + brick.Size[0]*2, Z: pos.Z + brick.Size[1]*2, Y: pos.Y}
		aboveeast := util.Vec3{X: pos.X + brick.Size[0]*2, Z: pos.Z, Y: pos.Y + brick.Size[2]*2}
		aboveBrick, aboveExists := bricks[above]
		eastBrick, eastExists := bricks[east]
		southBrick, southExists := bricks[south]
		aboveeastBrick, aboveeastExists := bricks[aboveeast]
		southeastBrick, southeastExists := bricks[southeast]
		equalAbove := aboveExists && brick.EqualTo(aboveBrick)
		equalEast := eastExists && brick.EqualTo(eastBrick)
		equalSouth := southExists && brick.EqualTo(southBrick)
		equalSoutheast := southeastExists && brick.EqualTo(southeastBrick)
		equalAboveeast := aboveeastExists && brick.EqualTo(aboveeastBrick)

		if equalAbove && equalAboveeast && equalEast {
			brick.Pos[0] += brick.Size[0]
			brick.Pos[2] += brick.Size[2]
			brick.Size[0] *= 2
			brick.Size[2] *= 2
			bricks[pos] = brick
			delete(bricks, above)
			delete(bricks, east)
			delete(bricks, aboveeast)
		} else if equalAbove {
			brick.Pos[2] += brick.Size[2]
			brick.Size[2] *= 2
			bricks[pos] = brick
			delete(bricks, above)
		} else if equalEast && equalSouth && equalSoutheast {
			brick.Pos[0] += eastBrick.Size[0]
			brick.Pos[1] += southBrick.Size[1]
			brick.Size[0] = brick.Size[0] + eastBrick.Size[0]
			brick.Size[1] = brick.Size[1] + southBrick.Size[1]
			bricks[pos] = brick
			delete(bricks, east)
			delete(bricks, south)
			delete(bricks, southeast)
		} else if equalEast {
			brick.Pos[0] += eastBrick.Size[0]
			brick.Size[0] = brick.Size[0] + eastBrick.Size[0]
			bricks[pos] = brick
			delete(bricks, east)
		} else if equalSouth {
			brick.Pos[1] += southBrick.Size[1]
			brick.Size[1] = brick.Size[1] + southBrick.Size[1]
			bricks[pos] = brick
			delete(bricks, south)
		}

		i++
	}
	ansi.Print(ansi.BrightBlue, "] ")
	percent := (float32(len(bricks)) / float32(total)) * 100.0
	ansi.Println(ansi.BrightGreen, fmt.Sprintf("%.1f%% original size", percent))
}
