package biome

import (
	"kmschr.com/mc2brs/ansi"
)

// https://minecraft.fandom.com/wiki/Java_Edition_data_values#Biomes
var BiomeSymbols = map[int32][]string{
	0:  {"\u2248"},                         // Ocean
	24: {"\u2248"},                         // Deep Ocean
	10: {"\u2248"},                         // Frozen Ocean
	50: {"\u2248"},                         // Deep Frozen ocean
	46: {"\u2248"},                         // Cold Ocean
	49: {"\u2248"},                         // Deep Cold Ocean
	45: {"\u2248"},                         // Lukewarm Ocean
	48: {"\u2248"},                         // Deep Lukewarm Ocean
	44: {"\u2248"},                         // Warm Ocean
	47: {"\u2248"},                         // Deep Warm Ocean
	7:  {ansi.RGB(0, 243, 247) + "\u2502"}, // River
	11: {ansi.RGB(0, 243, 247) + "\u2502"}, // Frozen River
	16: {ansi.RGB(255, 255, 0) + "\u2248",
		ansi.RGB(255, 255, 0) + "~"}, // Beach
	25:  {ansi.RGB(128, 128, 128) + "\u2248"}, // Stone Shore
	26:  {ansi.RGB(255, 255, 255) + "\u2248"}, // Snowy Beach
	4:   {"\u2660", "\u2663"},                 // Forest
	18:  {"\u2229", "n"},                      // Wooded Hills
	132: {"\u2660", "\u2663"},                 // Flower Forest
	27:  {"\u2660", "\u2663"},                 // Birch Forest
	28:  {"\u2229", "n"},                      // Birch Forest Hills
	155: {"\u2663", "\u2663"},                 // Tall Birch Forest
	156: {"\u2229", "n"},                      // Tall Birch Hills
	29:  {"\u2660", "\u2663"},                 // Dark Forest
	157: {"\u2663", "n"},                      // Dark Forest Hills
	21:  {"\u2660", "\u0393"},                 // Jungle
	22:  {"\u2229", "n"},                      // Jungle Hills
	149: {"\u2660", "\u0393"},                 // Modified Jungle
	23:  {"\u2660", "\u0393"},                 // Jungle Edge
	151: {"\u2660", "\u0393"},                 // Modified Jungle Edge
	168: {"\u2660", "\u2560"},                 // Bamboo Jungle
	169: {"\u2229", "n"},                      // Bamboo Jungle Hills
	5:   {"\u2191", "\u21A8"},                 // Taiga
	19:  {"\u2229", "n"},                      // Taiga Hills
	133: {ansi.RGB(192, 192, 192) + "\u2302",
		ansi.RGB(128, 128, 128) + "\u25B2",
		ansi.RGB(192, 192, 192) + "\u25B2"}, // Taiga Mountains
	30: {ansi.RGB(255, 255, 255) + "\u2660",
		ansi.RGB(255, 255, 255) + "\u2663"}, // Snowy Taiga
	31: {ansi.RGB(255, 255, 255) + "\u2229",
		ansi.RGB(255, 255, 255) + "n"}, // Snowy Taiga Hills
	158: {ansi.RGB(255, 255, 255) + "\u2302",
		ansi.RGB(255, 255, 255) + "\u25B2",
		ansi.RGB(255, 255, 255) + "\u25B2"}, // Snowy Taiga Mountains
	32:  {"\u2191", "\u21A8"}, // Giant tree taiga
	33:  {"\u2229", "n"},      // Giant tree taiga hills
	160: {"\u2191", "\u21A8"}, // Giant spruce taiga
	161: {"\u2229", "n"},      // Giant tree spruce taiga hills
	14:  {"\"", "\u03C4"},     // Mushrom fields
	15:  {"\"", "\u03C4"},     // Mushroom fields shore
	6:   {"\"", "\u2320"},     // Swamp
	134: {"\"", "n"},          // Swamp Hills
	35: {ansi.RGB(255, 255, 0) + "\"",
		ansi.RGB(255, 255, 0) + "n"}, // Savanna
	36: {ansi.RGB(255, 255, 0) + "\"",
		ansi.RGB(255, 255, 0) + "n"}, //  Savanna Plateau
	163: {ansi.RGB(255, 255, 0) + "\"",
		ansi.RGB(255, 255, 0) + "n"}, // Shattered Savannna
	164: {ansi.RGB(255, 255, 0) + "\"",
		ansi.RGB(255, 255, 0) + "n"}, // Shattered Savanna Plataue
	1:   {"n", "."}, // Plains
	129: {"n", "."}, // Sunflower Plains
	2: {ansi.RGB(255, 255, 0) + "\u2248",
		ansi.RGB(255, 255, 0) + "~"}, // Desert
	17: {ansi.RGB(255, 255, 0) + "\u2229",
		ansi.RGB(255, 255, 0) + "n"}, // Desert Hills
	130: {ansi.RGB(255, 255, 0) + "\u2248",
		ansi.RGB(10, 10, 200) + "~"}, // Desert Lakes
	12: {ansi.RGB(0, 243, 247) + ".",
		ansi.RGB(0, 243, 247) + "\u2219"}, // Snowy Tundra
	13: {ansi.RGB(255, 255, 255) + "\u2302",
		ansi.RGB(255, 255, 255) + "\u25B2",
		ansi.RGB(0, 243, 247) + "\u25B2"}, // Snowy Mountains
	140: {ansi.RGB(0, 243, 247) + "\u2592",
		ansi.RGB(0, 243, 247) + "\u2591"}, // Ice Spikes
	3: {ansi.RGB(192, 192, 192) + "\u2302",
		ansi.RGB(128, 128, 128) + "\u25B2",
		ansi.RGB(192, 192, 192) + "\u25B2"}, // Mountains
	34: {"\u2660",
		ansi.RGB(192, 192, 192) + "\u2302",
		ansi.RGB(128, 128, 128) + "\u25B2"}, // Wooded Mountains
	131: {ansi.RGB(192, 192, 192) + "\u2302",
		ansi.RGB(128, 128, 128) + "\u25B2",
		ansi.RGB(192, 192, 192) + "\u25B2"}, // Gravelly Mountains
	162: {ansi.RGB(192, 192, 192) + "\u2302",
		ansi.RGB(128, 128, 128) + "\u25B2",
		ansi.RGB(192, 192, 192) + "\u25B2"}, // Modified Gravelly Mountains
	20: {ansi.RGB(192, 192, 192) + "\u2302"}, // Mountain Edge
	37: {ansi.RGB(189, 102, 32) + "v",
		ansi.RGB(189, 102, 32) + "\u221A"}, // Badlands
	39:  {ansi.RGB(189, 102, 32) + "\u221A"}, // Badlands Plateau
	167: {ansi.RGB(189, 102, 32) + "\u221A"}, // Modified Badlands Plateau
	38: {ansi.RGB(189, 102, 32) + "\u221A",
		ansi.RGB(61, 113, 50) + "\u2660"}, // Wooded Badlands Plateau
	166: {ansi.RGB(189, 102, 32) + "\u221A",
		ansi.RGB(61, 113, 50) + "\u2660"}, // Modified Wooded Badlands Plateau
	165: {ansi.RGB(189, 102, 32) + "v",
		ansi.RGB(189, 102, 32) + "\u221A"}, // Eroded Badlands
	174: {"n", "."}, // Dripstone Caves
	175: {"n", "."}, // Lush Caves
	8: {ansi.RGB(97, 38, 38) + "\u2593",
		ansi.RGB(97, 38, 38) + "\u2592"}, // Nether Wastes
	171: {ansi.RGB(93, 26, 31) + "\u2660",
		ansi.RGB(93, 26, 31) + "\u2663"}, // Crimson Forest
	172: {ansi.RGB(58, 59, 77) + "\u2660",
		ansi.RGB(58, 59, 77) + "\u2663"}, // Warped Forest
	170: {ansi.RGB(77, 58, 47) + "\u2593",
		ansi.RGB(77, 58, 47) + "\u2592"}, // Soul Sand Valley
	173: {ansi.RGB(71, 70, 76) + "^",
		ansi.RGB(213, 90, 18) + "\u2593"}, // Basalt Deltas
	9: {ansi.RGB(255, 255, 255) + "\u2219",
		ansi.RGB(255, 255, 255) + "\u00A0"}, // The End
	40:  {ansi.RGB(217, 220, 155) + "\u25CB"}, // Small End Islands
	41:  {ansi.RGB(217, 220, 155) + "\u2248"}, // End Midlands
	42:  {ansi.RGB(217, 220, 155) + "\u2248"}, // End Highlands
	43:  {ansi.RGB(217, 220, 155) + "\u2248"}, // End Barrens
	127: {ansi.RGB(40, 40, 40) + "\u2588"},    // The Void
}
