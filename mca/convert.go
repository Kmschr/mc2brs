package mca

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/brs"
	"kmschr.com/mc2brs/mca/biome"
	"kmschr.com/mc2brs/util"
)

// mapAndPreview creates an image for a region, saves the image,
// and then returns the image as a Brickadia save preview image
func (w *World) mapAndPreview(r Region) *image.RGBA {
	minimap := w.makeMinimap(r)
	f, _ := os.Create("minimap.png")
	png.Encode(f, minimap)
	f.Close()
	return util.PadImage(minimap)
}

// ConvertRegion makes bricks out of blocks for a region
func (w *World) ConvertRegion(regionX int, regionZ int) brs.Save {
	r, exists := w.Regions[regionCoords{regionX, regionZ}]
	if !exists {
		ansi.Println(ansi.Red, "Tried to convert non existent region")
		ansi.Quit()
	}
	save := brs.NewSave()
	save.Screenshot = w.mapAndPreview(r)
	dx := regionX * regionWidth
	dz := regionZ * regionWidth
	bricks := make(map[util.Vec3]brs.Brick)
	ansi.Println(ansi.BrightRed, fmt.Sprintf("\nGenerating %s_%s_%d_%d...", w.Name, w.Dimension, regionX, regionZ))
	for chunkZ := 0; chunkZ < 32; chunkZ++ {
		for chunkX := 0; chunkX < 32; chunkX++ {
			// get chunk, skip if not loaded
			chunk, loaded := r.ChunkAt(chunkX, chunkZ)
			if !loaded {
				fmt.Print("  ")
				continue
			}
			for sectionY := chunk.minSection; sectionY < chunk.maxSection; sectionY++ {
				// get section, skip if not loaded
				_, exists := chunk.Sections[sectionY]
				if !exists {
					continue
				}
				// final layer of iteration, coordinates within a sector
				for cz := 0; cz < 16; cz++ {
					for cx := 0; cx < 16; cx++ {
						for sy := 0; sy < 16; sy++ {
							rx := chunkX<<4 + cx  // x within the region
							rz := chunkZ<<4 + cz  // z within the region
							y := sectionY<<4 + sy // worldspace y
							x := rx + dx          // worldspace x
							z := rz + dz          // worldspace z

							// Skip air/non rendered blocks
							block, exists := r.BlockAt(rx, rz, y)
							if !exists || block.isSkipped() || w.surrounded(block, x, z, y) {
								continue
							}
							color, known := w.ColorAtBRS(block, x, z, y)
							if !known {
								continue
							}
							pos := util.Vec3{X: x << 1 * ansi.Scale(), Z: z << 1 * ansi.Scale(), Y: y << 1 * ansi.Scale()}
							bricks[pos] = block.resolveBrick(x, z, y, color)
						}
					}
				}
			}
			// print chunk symbol to indicate chunk has finished processing
			fmt.Print(chunk.Symbol() + chunk.Symbol())
		}
		fmt.Println("")
	}
	postProcess(bricks)
	if ansi.Optimize() {
		optimize(bricks)
		optimize(bricks)
		optimize(bricks)
	}
	save.Bricks = bricks
	return save
}

// resolveBrick returns the brick conversion of a block
func (b BlockState) resolveBrick(x, z, y int, color color.RGBA) brs.Brick {
	var brick brs.Brick
	if b.isStairs() {
		brick = b.resolveStairs(x, z, y, color)
	} else if b.isSlab() {
		brick = b.resolveSlab(x, z, y, color)
	} else {
		brick = brs.NewBrick(color, x, z, y+64, b.Name)
	}
	modifyMaterial(b, &brick)
	return brick
}

// resolveStairs handles stairs block to brick conversion
func (b BlockState) resolveStairs(x, z, y int, color color.RGBA) brs.Brick {
	var brick brs.Brick
	shape := b.Properties["shape"]
	half := b.Properties["half"]
	facing := b.Properties["facing"]
	if shape == "straight" {
		brick = brs.NewWedge(color, x, z, y+64, b.Name, facing)
		if half == "top" {
			brick.Rotation++
		}
	} else if shape == "outer_right" || shape == "outer_left" {
		brick = brs.NewCorner(color, x, z, y+64, b.Name, facing)
		if shape == "outer_left" {
			brick.Rotation += 3
		}
		if half == "top" {
			brick.Direction = brs.ZNegative
			brick.Rotation++
			if (shape == "outer_right" && (facing == "north" || facing == "south")) ||
				(shape == "outer_left" && (facing == "east" || facing == "west")) {
				brick.Rotation += 2
			}
		}
	} else if shape == "inner_right" || shape == "inner_left" {
		brick = brs.NewInnerCorner(color, x, z, y+64, b.Name, facing)
		if shape == "inner_left" {
			brick.Rotation += 3
		}
		if half == "top" {
			brick.Direction = brs.ZNegative
			brick.Rotation++
		}
	}
	return brick
}

func (b BlockState) resolveSlab(x, z, y int, color color.RGBA) brs.Brick {
	if ansi.Scale() == 1 {
		return brs.NewBrick(color, x, z, y+64, b.Name)
	}

	var brick brs.Brick

	half := b.Properties["type"]
	switch half {
	case "top":
		brick = brs.NewBrick(color, x, z, y+64, b.Name)
		brick.Pos[2] += brick.Size[2] / 2
		brick.Size[2] /= 2
	case "bottom":
		brick = brs.NewBrick(color, x, z, y+64, b.Name)
		brick.Pos[2] -= brick.Size[2] / 2
		brick.Size[2] /= 2
	default:
		brick = brs.NewBrick(color, x, z, y+64, b.Name)
	}

	return brick
}

// modifyMaterial adjusts a bricks material according to a blocks category
func modifyMaterial(b BlockState, brick *brs.Brick) {
	if b.Glows() {
		brick.MaterialIndex = brs.MaterialGlow
		brick.Intensity = 2
	} else if b.isWatery() {
		brick.MaterialIndex = brs.MaterialGlass
		brick.Intensity = 8
		brick.Collision = false
	} else if b.isGlass() {
		brick.MaterialIndex = brs.MaterialGlass
		brick.Intensity = 7
	} else if b.isPortal() {
		brick.MaterialIndex = brs.MaterialGhost
		brick.Collision = false
	} else if b.isIce() {
		brick.MaterialIndex = brs.MaterialGlass
		brick.Intensity = 9
	} else if b.isLava() {
		brick.Collision = false
		brick.MaterialIndex = brs.MaterialGlow
		brick.Intensity = 1
	}

	if ansi.Lights() && b.isLight() {
		brick.Light = true
	}
}

// postProcess goes through the existing bricks and makes adjustments,
// such as making snow go into the brick below it
func postProcess(bricks map[util.Vec3]brs.Brick) {
	ansi.Print(ansi.BrightYellow, "\nPost Processing")
	ansi.Print(ansi.BrightBlue, " [")
	total := len(bricks)
	i := 0
	for pos, brick := range bricks {
		if i%(total/16) == 0 {
			ansi.Print(ansi.BrightBlue, "#")
		}
		if brick.Name == "minecraft:snow" {
			below := util.Vec3{X: pos.X, Z: pos.Z, Y: pos.Y - brick.Size[2]*2}
			belowBrick, exists := bricks[below]
			if exists {
				belowBrick.Color = brick.Color
				bricks[below] = belowBrick
			}
			delete(bricks, pos)
		}
		i++
	}
	ansi.Println(ansi.BrightBlue, "]")
}

// surrounded returns whether or not a block is visible due to being covered by other blocks
func (w *World) surrounded(block BlockState, x, z, y int) bool {
	surrounding := w.SurroundingAt(x, z, y)
	if len(surrounding) != 6 && y != 0 && y != 255 {
		return false
	}
	for _, b := range surrounding {
		if block.isNotBlockedBy(b) {
			return false
		}
	}
	return true
}

// ColorAt returns the BRS color of a block at any worldspace coordinate, it also returns
// whether or not the color is known
func (w *World) ColorAtBRS(block BlockState, x, z, y int) (color.RGBA, bool) {
	var c color.RGBA
	if block.Name == "minecraft:grass_block" {
		c = brs.ConvertColor(w.GrassBlend(x, z, y))
	} else if block.isFoliage() {
		if block.Name == "minecraft:birch_leaves" {
			c = brs.ConvertColor(biome.ColorBirchLeaves)
		} else if block.Name == "minecraft:spruce_leaves" {
			c = brs.ConvertColor(biome.ColorSpruceLeaves)
		} else if block.Name == "minecraft:azalea_leaves" {
			c = brs.ConvertColor(biome.ColorSpruceLeaves)
		} else {
			c = brs.ConvertColor(w.FoliageBlend(x, z, y))
		}
	} else if block.isWatery() {
		c = brs.ConvertColor(w.WaterBlend(x, z, y))
	} else {
		col, hasColor := block.colorBRS()
		if !hasColor {
			return c, false
		}
		c = col
	}
	return c, true
}

// ColorAt returns the color of a block at any worldspace coordinate, it also returns
// whether or not the color is known
func (w *World) ColorAt(block BlockState, x, z, y int) (color.RGBA, bool) {
	var c color.RGBA
	if block.Name == "minecraft:grass_block" {
		c = w.GrassBlend(x, z, y)
	} else if block.isFoliage() {
		if block.Name == "minecraft:birch_leaves" {
			c = biome.ColorBirchLeaves
		} else if block.Name == "minecraft:spruce_leaves" {
			c = biome.ColorSpruceLeaves
		} else {
			c = w.FoliageBlend(x, z, y)
		}
	} else if block.isWatery() {
		c = w.WaterBlend(x, z, y)
	} else {
		col, hasColor := block.color()
		if !hasColor {
			return c, false
		}
		c = col
	}
	return c, true
}
