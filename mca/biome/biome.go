package biome

import (
	"image/color"
	"strings"

	"kmschr.com/mc2brs/util"
)

// TODO: Add temperature modifier for frozen biomes

type Biome struct {
	Name        string
	Temperature float32
	Rainfall    float32
	WaterColor  color.RGBA
}

var (
	// Hard-coded water tints available for biomes
	TintWaterDefault  = color.RGBA{0x3F, 0x76, 0xE4, 0xFF}
	TintWaterCold     = color.RGBA{0x3D, 0x57, 0xD6, 0xFF}
	TintWaterFrozen   = color.RGBA{0x39, 0x38, 0xC9, 0xFF}
	TintWaterLukeWarm = color.RGBA{0x45, 0xAD, 0xF2, 0xFF}
	TintWaterSwamp    = color.RGBA{0x61, 0x7B, 0x64, 0xFF}
	TintWaterWarm     = color.RGBA{0x43, 0xD5, 0xEE, 0xFF}

	// Base colors
	BaseWater   = color.RGBA{0xA5, 0xA5, 0xA5, 0xF4}
	BaseGrass   = color.RGBA{0x95, 0x95, 0x95, 0xFF}
	BaseFoliage = color.RGBA{0x95, 0x95, 0x95, 0xFF}

	// Blended water colors
	ColorWaterDefault  = util.TintColor(BaseWater, TintWaterDefault)
	ColorWaterCold     = util.TintColor(BaseWater, TintWaterCold)
	ColorWaterFrozen   = util.TintColor(BaseWater, TintWaterFrozen)
	ColorWaterLukeWarm = util.TintColor(BaseWater, TintWaterLukeWarm)
	ColorWaterSwamp    = util.TintColor(BaseWater, TintWaterSwamp)
	ColorWaterWarm     = util.TintColor(BaseWater, TintWaterWarm)

	// Noise generators
	TemperatureNoise       = NewPerlinNoiseGen(NewSeededRandom(1234), []int{0})
	FrozenTemperatureNoise = NewPerlinNoiseGen(NewSeededRandom(3456), []int{-2, -1, 0})
	BiomeInfoNoise         = NewPerlinNoiseGen(NewSeededRandom(2345), []int{0})

	// Biomes that are specifically checked against
	DarkForest int32 = 29
	Swamp      int32 = 6

	// Special colors
	TintSwampGrassDark  = color.RGBA{0x4C, 0x76, 0x3C, 0xFF}
	TintSwampGrassLight = color.RGBA{0x6A, 0x70, 0x39, 0xFF}
	TintBirchLeaves     = color.RGBA{0x80, 0xA7, 0x55, 0xFF}
	TintSpruceLeaves    = color.RGBA{0x61, 0x99, 0x61, 0xFF}

	// Blended special colors
	ColorBirchLeaves  = util.TintColor(BaseFoliage, TintBirchLeaves)
	ColorSpruceLeaves = util.TintColor(BaseFoliage, TintSpruceLeaves)
	ColorAzaleaLeaves = color.RGBA{86, 110, 43, 255}
)

var Biomes = map[int32]Biome{
	0:   {Name: "Ocean", Temperature: 0.5, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	1:   {Name: "Plains", Temperature: 0.8, Rainfall: 0.4, WaterColor: ColorWaterDefault},
	2:   {Name: "Desert", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	3:   {Name: "Mountains", Temperature: 0.2, Rainfall: 0.3, WaterColor: ColorWaterDefault},
	4:   {Name: "Forest", Temperature: 0.7, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	5:   {Name: "Taiga", Temperature: 0.25, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	6:   {Name: "Swamp", Temperature: 0.8, Rainfall: 0.9, WaterColor: ColorWaterSwamp},
	7:   {Name: "River", Temperature: 0.5, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	8:   {Name: "Nether", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	9:   {Name: "The End", Temperature: 0.5, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	10:  {Name: "Frozen Ocean", Temperature: 0.0, Rainfall: 0.5, WaterColor: ColorWaterFrozen},
	11:  {Name: "Frozen River", Temperature: 0.0, Rainfall: 0.5, WaterColor: ColorWaterFrozen},
	12:  {Name: "Snowy Tundra", Temperature: 0.0, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	13:  {Name: "Snowy Mountains", Temperature: 0.0, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	14:  {Name: "Mushroom Fields", Temperature: 0.9, Rainfall: 1.0, WaterColor: ColorWaterDefault},
	15:  {Name: "Mushroom Field Shore", Temperature: 0.9, Rainfall: 1.0, WaterColor: ColorWaterDefault},
	16:  {Name: "Beach", Temperature: 0.8, Rainfall: 0.4, WaterColor: ColorWaterDefault},
	17:  {Name: "Desert Hills", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	18:  {Name: "Wooded Hills", Temperature: 0.7, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	19:  {Name: "Taiga Hills", Temperature: 0.25, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	20:  {Name: "Mountain Edge", Temperature: 0.2, Rainfall: 0.3, WaterColor: ColorWaterDefault},
	21:  {Name: "Jungle", Temperature: 0.95, Rainfall: 0.9, WaterColor: ColorWaterDefault},
	22:  {Name: "Jungle Hills", Temperature: 0.95, Rainfall: 0.9, WaterColor: ColorWaterDefault},
	23:  {Name: "Jungle Edge", Temperature: 0.95, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	24:  {Name: "Deep Ocean", Temperature: 0.5, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	25:  {Name: "Stone Shore", Temperature: 0.2, Rainfall: 0.3, WaterColor: ColorWaterDefault},
	26:  {Name: "Snowy Beach", Temperature: 0.05, Rainfall: 0.3, WaterColor: ColorWaterDefault},
	27:  {Name: "Birch Forest", Temperature: 0.6, Rainfall: 0.6, WaterColor: ColorWaterDefault},
	28:  {Name: "Birch Forest Hills", Temperature: 0.6, Rainfall: 0.6, WaterColor: ColorWaterDefault},
	29:  {Name: "Dark Forest", Temperature: 0.7, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	30:  {Name: "Snowy Taiga", Temperature: -0.5, Rainfall: 0.4, WaterColor: ColorWaterCold},
	31:  {Name: "Snowy Taiga Hills", Temperature: -0.5, Rainfall: 0.4, WaterColor: ColorWaterCold},
	32:  {Name: "Giant Tree Taiga", Temperature: 0.3, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	33:  {Name: "Giant Tree Taiga Hills", Temperature: 0.3, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	34:  {Name: "Wooded Mountains", Temperature: 0.2, Rainfall: 0.3, WaterColor: ColorWaterDefault},
	35:  {Name: "Savanna", Temperature: 1.2, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	36:  {Name: "Savanna Plateau", Temperature: 1.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	37:  {Name: "Badlands", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	38:  {Name: "Wooded Badlands Plateau", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	39:  {Name: "Badlands Plateau", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	40:  {Name: "Small End Islands", Temperature: 0.5, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	41:  {Name: "End Midlands", Temperature: 0.5, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	42:  {Name: "End Highlands", Temperature: 0.5, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	43:  {Name: "End Barrens", Temperature: 0.5, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	44:  {Name: "Warm Ocean", Temperature: 0.8, Rainfall: 0.5, WaterColor: ColorWaterWarm},
	45:  {Name: "Lukewarm Ocean", Temperature: 0.8, Rainfall: 0.5, WaterColor: ColorWaterLukeWarm},
	46:  {Name: "Cold Ocean", Temperature: 0.8, Rainfall: 0.5, WaterColor: ColorWaterCold},
	47:  {Name: "Deep Warm Ocean", Temperature: 0.8, Rainfall: 0.5, WaterColor: ColorWaterWarm},
	48:  {Name: "Deep Lukewarm Ocean", Temperature: 0.8, Rainfall: 0.5, WaterColor: ColorWaterLukeWarm},
	49:  {Name: "Deep Cold Ocean", Temperature: 0.8, Rainfall: 0.5, WaterColor: ColorWaterCold},
	50:  {Name: "Deep Frozen Ocean", Temperature: 0.8, Rainfall: 0.5, WaterColor: ColorWaterFrozen},
	127: {Name: "The Void", Temperature: 0.5, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	128: {Name: "Unknown Biome", Temperature: 0.8, Rainfall: 0.4, WaterColor: ColorWaterDefault},
	129: {Name: "Sunflower Plains", Temperature: 0.8, Rainfall: 0.4, WaterColor: ColorWaterDefault},
	130: {Name: "Desert Lakes", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	131: {Name: "Gravelly Mountains", Temperature: 0.2, Rainfall: 0.3, WaterColor: ColorWaterDefault},
	132: {Name: "Flower Forest", Temperature: 0.7, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	133: {Name: "Taiga Mountains", Temperature: 0.25, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	134: {Name: "Swamp Hills", Temperature: 0.8, Rainfall: 0.9, WaterColor: ColorWaterSwamp},
	140: {Name: "Ice Spikes", Temperature: 0.0, Rainfall: 0.5, WaterColor: ColorWaterDefault},
	149: {Name: "Modified Jungle", Temperature: 0.95, Rainfall: 0.9, WaterColor: ColorWaterDefault},
	151: {Name: "Modified Jungle Edge", Temperature: 0.95, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	155: {Name: "Tall Birch Forest", Temperature: 0.6, Rainfall: 0.6, WaterColor: ColorWaterDefault},
	156: {Name: "Tall Birch Hills", Temperature: 0.6, Rainfall: 0.6, WaterColor: ColorWaterDefault},
	157: {Name: "Dark Forest Hills", Temperature: 0.7, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	158: {Name: "Snowy Taiga Mountains", Temperature: -0.5, Rainfall: 0.4, WaterColor: ColorWaterDefault},
	160: {Name: "Giant Spruce Taiga", Temperature: 0.25, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	161: {Name: "Giant Spruce Taiga Hills", Temperature: 0.25, Rainfall: 0.8, WaterColor: ColorWaterDefault},
	162: {Name: "Gravelly Mountains+", Temperature: 0.2, Rainfall: 0.3, WaterColor: ColorWaterDefault},
	163: {Name: "Shattered Savanna", Temperature: 1.1, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	164: {Name: "Shattered Savanna Plateau", Temperature: 1.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	165: {Name: "Eroded Badlands", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	166: {Name: "Modified Wooded Badlands Plateau", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	167: {Name: "Modified Badlands Plateau", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	168: {Name: "Bamboo Jungle", Temperature: 0.95, Rainfall: 0.9, WaterColor: ColorWaterDefault},
	169: {Name: "Bamboo Jungle Hills", Temperature: 0.95, Rainfall: 0.9, WaterColor: ColorWaterDefault},
	170: {Name: "Soul Sand Valley", Temperature: .0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	171: {Name: "Crimson Forest", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	172: {Name: "Warped Forest", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	173: {Name: "Basalt Deltas", Temperature: 2.0, Rainfall: 0.0, WaterColor: ColorWaterDefault},
	174: {Name: "Dripstone Caves", Temperature: 0.8, Rainfall: 0.4, WaterColor: ColorWaterDefault},
	175: {Name: "Lush Caves", Temperature: 0.8, Rainfall: 0.4, WaterColor: ColorWaterDefault},
}

var grassCorners = [3]color.RGBA{
	{191, 183, 85, 255},  // lower left
	{128, 180, 151, 255}, // lower right
	{71, 205, 51, 255},   // upper left
}

var foliageCorners = [3]color.RGBA{
	{174, 164, 42, 255}, // lower left
	{96, 161, 123, 255}, // lower right
	{26, 191, 0, 255},   // upper left
}

func heightAdjustedTemperature(biome int32, x, z, y int) float32 {
	baseTemp := Biomes[biome].Temperature
	if y > 64 {
		f := TemperatureNoise.GetValue(float64(x>>3), float64(z>>3), false) * 4.0
		return baseTemp - (float32(f)+float32(y)-64.0)*0.05/30.0
	} else {
		return baseTemp
	}
}

func biomeTint(temperature float32, rainfall float32, corners [3]color.RGBA) color.RGBA {
	//temperature = util.Clamp(temperature-float32(elevation)*0.00166667, 0, 1)
	rainfall = util.Clampf(rainfall, 0, 1)
	rainfall *= temperature

	var lambda [3]float32
	lambda[0] = temperature - rainfall
	lambda[1] = 1 - temperature
	lambda[2] = rainfall

	var red float32
	var green float32
	var blue float32
	for i := 0; i < 3; i++ {
		red += lambda[i] * float32(corners[i].R)
		green += lambda[i] * float32(corners[i].G)
		blue += lambda[i] * float32(corners[i].B)
	}
	r := uint8(util.Clampf(red, 0, 255))
	g := uint8(util.Clampf(green, 0, 255))
	b := uint8(util.Clampf(blue, 0, 255))
	return color.RGBA{r, g, b, 0xFF}
}

func swampTint(x int, z int) color.RGBA {
	d0 := BiomeInfoNoise.GetValue(float64(x)*0.0225, float64(z)*0.0225, false)
	if d0 < -0.1 {
		return TintSwampGrassLight
	}
	return TintSwampGrassDark
}

func GrassColor(biome int32, x, z, y int) color.RGBA {
	if biome == Swamp {
		return util.TintColor(BaseGrass, swampTint(x, z))
	}

	temp := heightAdjustedTemperature(biome, x, z, y)
	tint := biomeTint(temp, Biomes[biome].Rainfall, grassCorners)
	if biome == DarkForest {
		tint = util.IntRGB((util.RGBInt(tint) & 0xFEFEFE) + 0x28340A>>1)
	}
	return util.TintColor(BaseGrass, tint)
}

func FoliageColor(biome int32, x, z, y int) color.RGBA {
	if biome == Swamp {
		return util.TintColor(BaseFoliage, swampTint(x, z))
	}

	temp := heightAdjustedTemperature(biome, x, z, y)
	tint := biomeTint(temp, Biomes[biome].Rainfall, foliageCorners)
	return util.TintColor(BaseFoliage, tint)
}

func IsWater(biome int32) bool {
	switch biome {
	case 0, 24, 10, 50, 46, 49, 45, 48, 44, 47, 7, 11:
		return true
	default:
		return false
	}
}

func IsMountain(biome int32) bool {
	return strings.Contains(Biomes[biome].Name, "Mountains")
}

func IsBadlandsWoodedPlateau(biome int32) bool {
	return biome == 38 || biome == 166
}
