package mca

// air is a "constant set" of air block names.
var air = map[string]bool{
	"minecraft:air":            true,
	"minecraft:cave_air":       true,
	"minecraft:structure_void": true,
	"minecraft:void_air":       true,
}

// transparent is a "constant set" of block names for
// transparent blocks that dont fit in any other category
var transparent = map[string]bool{
	"minecraft:ice":           true,
	"minecraft:nether_portal": true,
	"minecraft:end_portal":    true,
}

var portal = map[string]bool{
	"minecraft:nether_portal": true,
	"minecraft:end_portal":    true,
}

var ice = map[string]bool{
	"minecraft:ice":         true,
	"minecraft:frosted_ice": true,
}

// glow is a "constant set" of block names for
// light emitting blocks that "glow"
var glow = map[string]bool{
	"minecraft:glowstone":      true,
	"minecraft:shroomlight":    true,
	"minecraft:glow_lichen":    true,
	"minecraft:jack_o_lantern": true,
	"minecraft:lit_pumpkin":    true,
	"minecraft:sea_lantern":    true,
}

var light = map[string]bool{
	"minecraft:glowstone":   true,
	"minecraft:shroomlight": true,
	"minecraft:glow_lichen": true,
}

// foliage is a "constant set" of foliage block names.
var foliage = map[string]bool{
	"minecraft:azalea_leaves":   true,
	"minecraft:oak_leaves":      true,
	"minecraft:spruce_leaves":   true,
	"minecraft:birch_leaves":    true,
	"minecraft:jungle_leaves":   true,
	"minecraft:acacia_leaves":   true,
	"minecraft:dark_oak_leaves": true,
}

// water is a "constant set" of water block names.
// this includes sea plants that are treated as water
// despite not having the waterlogged property
var water = map[string]bool{
	"minecraft:water":         true,
	"minecraft:tall_seagrass": true,
	"minecraft:seagrass":      true,
	"minecraft:kelp_plant":    true,
	"minecraft:kelp":          true,
	"minecraft:bubble_column": true,
}

// glass is a "constant set" of block names for all glass type blocks
var glass = map[string]bool{
	"minecraft:glass":                         true,
	"minecraft:glass_pane":                    true,
	"minecraft:white_stained_glass":           true,
	"minecraft:orange_stained_glass":          true,
	"minecraft:magenta_stained_glass":         true,
	"minecraft:light_blue_stained_glass":      true,
	"minecraft:yellow_stained_glass":          true,
	"minecraft:lime_stained_glass":            true,
	"minecraft:pink_stained_glass":            true,
	"minecraft:gray_stained_glass":            true,
	"minecraft:light_gray_stained_glass":      true,
	"minecraft:cyan_stained_glass":            true,
	"minecraft:purple_stained_glass":          true,
	"minecraft:blue_stained_glass":            true,
	"minecraft:brown_stained_glass":           true,
	"minecraft:green_stained_glass":           true,
	"minecraft:red_stained_glass":             true,
	"minecraft:black_stained_glass":           true,
	"minecraft:white_stained_glass_pane":      true,
	"minecraft:orange_stained_glass_pane":     true,
	"minecraft:magenta_stained_glass_pane":    true,
	"minecraft:light_blue_stained_glass_pane": true,
	"minecraft:yellow_stained_glass_pane":     true,
	"minecraft:lime_stained_glass_pane":       true,
	"minecraft:pink_stained_glass_pane":       true,
	"minecraft:gray_stained_glass_pane":       true,
	"minecraft:light_gray_stained_glass_pane": true,
	"minecraft:cyan_stained_glass_pane":       true,
	"minecraft:purple_stained_glass_pane":     true,
	"minecraft:blue_stained_glass_pane":       true,
	"minecraft:brown_stained_glass_pane":      true,
	"minecraft:green_stained_glass_pane":      true,
	"minecraft:red_stained_glass_pane":        true,
	"minecraft:black_stained_glass_pane":      true,
	"minecraft:tinted_glass":                  true,
}

// stairs is a "constant set" of block names for all stair type blocks
var stairs = map[string]bool{
	"minecraft:purpur_stairs":                    true,
	"minecraft:oak_stairs":                       true,
	"minecraft:cobblestone_stairs":               true,
	"minecraft:brick_stairs":                     true,
	"minecraft:stone_brick_stairs":               true,
	"minecraft:nether_brick_stairs":              true,
	"minecraft:sandstone_stairs":                 true,
	"minecraft:spruce_stairs":                    true,
	"minecraft:birch_stairs":                     true,
	"minecraft:jungle_stairs":                    true,
	"minecraft:quartz_stairs":                    true,
	"minecraft:acacia_stairs":                    true,
	"minecraft:dark_oak_stairs":                  true,
	"minecraft:prismarine_stairs":                true,
	"minecraft:prismarine_brick_stairs":          true,
	"minecraft:dark_prismarine_stairs":           true,
	"minecraft:red_sandstone_stairs":             true,
	"minecraft:polished_granite_stairs":          true,
	"minecraft:smooth_red_sandstone_stairs":      true,
	"minecraft:mossy_stone_brick_stairs":         true,
	"minecraft:polished_diorite_stairs":          true,
	"minecraft:mossy_cobblestone_stairs":         true,
	"minecraft:end_stone_brick_stairs":           true,
	"minecraft:stone_stairs":                     true,
	"minecraft:smooth_sandstone_stairs":          true,
	"minecraft:smooth_quartz_stairs":             true,
	"minecraft:granite_stairs":                   true,
	"minecraft:andesite_stairs":                  true,
	"minecraft:red_nether_brick_stairs":          true,
	"minecraft:polished_andesite_stairs":         true,
	"minecraft:diorite_stairs":                   true,
	"minecraft:crimson_stairs":                   true,
	"minecraft:warped_stairs":                    true,
	"minecraft:blackstone_stairs":                true,
	"minecraft:polished_blackstone_stairs":       true,
	"minecraft:polished_blackstone_brick_stairs": true,
	"minecraft:deepslate_brick_stairs":           true,
}

// stairs is a "constant set" of block names for all slab type blocks
var slabs = map[string]bool{
	"minecraft:acacia_slab":                    true,
	"minecraft:andesite_slab":                  true,
	"minecraft:birch_slab":                     true,
	"minecraft:blackstone_slab":                true,
	"minecraft:brick_slab":                     true,
	"minecraft:cobbled_deepslate_slab":         true,
	"minecraft:cobblestone_slab":               true,
	"minecraft:crimson_slab":                   true,
	"minecraft:cut_copper_slab":                true,
	"minecraft:cut_red_sandstone_slab":         true,
	"minecraft:cut_sandstone_slab":             true,
	"minecraft:dark_oak_slab":                  true,
	"minecraft:dark_prismarine_slab":           true,
	"minecraft:deepslate_brick_slab":           true,
	"minecraft:deepslate_tile_slab":            true,
	"minecraft:diorite_slab":                   true,
	"minecraft:end_stone_brick_slab":           true,
	"minecraft:exposed_cut_copper_slab":        true,
	"minecraft:granite_slab":                   true,
	"minecraft:jungle_slab":                    true,
	"minecraft:mossy_cobblestone_slab":         true,
	"minecraft:mossy_stone_brick_slab":         true,
	"minecraft:nether_brick_slab":              true,
	"minecraft:oak_slab":                       true,
	"minecraft:oxidized_copper_slab":           true,
	"minecraft:oxidized_cut_copper_slab":       true,
	"minecraft:petrified_oak_slab":             true,
	"minecraft:polished_andesite_slab":         true,
	"minecraft:polished_blackstone_brick_slab": true,
	"minecraft:polished_blackstone_slab":       true,
	"minecraft:polished_diorite_slab":          true,
	"minecraft:polished_deepslate_slab":        true,
	"minecraft:polished_granite_slab":          true,
	"minecraft:prismarine_brick_slab":          true,
	"minecraft:prismarine_slab":                true,
	"minecraft:purpur_slab":                    true,
	"minecraft:quartz_slab":                    true,
	"minecraft:red_nether_brick_slab":          true,
	"minecraft:red_sandstone_slab":             true,
	"minecraft:sandstone_slab":                 true,
	"minecraft:smooth_red_sandstone_slab":      true,
	"minecraft:smooth_sandstone_slab":          true,
	"minecraft:smooth_quartz_slab":             true,
	"minecraft:smooth_stone_slab":              true,
	"minecraft:spruce_slab":                    true,
	"minecraft:stone_brick_slab":               true,
	"minecraft:stone_slab":                     true,
	"minecraft:warped_slab":                    true,
}

// sprite is a "constant set" of sprite block names.
// A sprite here is not purely just sprite blocks, but
// any block that is not treated as a standard cube or
// handled specially otherwise. This includes but is not
// limited to flowers, trapdoors, carpets, heads, banners,
// doors, beds, and other odd shape blocks
var sprite = map[string]bool{
	"minecraft:acacia_button":                      true,
	"minecraft:acacia_door":                        true,
	"minecraft:acacia_pressure_plate":              true,
	"minecraft:acacia_sapling":                     true,
	"minecraft:acacia_sign":                        true,
	"minecraft:acacia_trapdoor":                    true,
	"minecraft:acacia_wall_sign":                   true,
	"minecraft:activator_rail":                     true,
	"minecraft:allium":                             true,
	"minecraft:anvil":                              true,
	"minecraft:attached_melon_stem":                true,
	"minecraft:attached_pumpkin_stem":              true,
	"minecraft:amethyst_cluster":                   true,
	"minecraft:azalea_leaves_flowers":              true,
	"minecraft:azure_bluet":                        true,
	"minecraft:bamboo":                             true,
	"minecraft:bamboo_sapling":                     true,
	"minecraft:barrier":                            true,
	"minecraft:beetroots":                          true,
	"minecraft:bell":                               true,
	"minecraft:big_dripleaf":                       true,
	"minecraft:big_dripleaf_stem":                  true,
	"minecraft:birch_button":                       true,
	"minecraft:birch_door":                         true,
	"minecraft:birch_pressure_plate":               true,
	"minecraft:birch_sapling":                      true,
	"minecraft:birch_sign":                         true,
	"minecraft:birch_trapdoor":                     true,
	"minecraft:birch_wall_sign":                    true,
	"minecraft:black_banner":                       true,
	"minecraft:black_bed":                          true,
	"minecraft:black_candle":                       true,
	"minecraft:black_candle_cake":                  true,
	"minecraft:black_carpet":                       true,
	"minecraft:black_wall_banner":                  true,
	"minecraft:blue_banner":                        true,
	"minecraft:blue_bed":                           true,
	"minecraft:blue_candle":                        true,
	"minecraft:blue_candle_cake":                   true,
	"minecraft:blue_carpet":                        true,
	"minecraft:blue_orchid":                        true,
	"minecraft:blue_wall_banner":                   true,
	"minecraft:brain_coral":                        true,
	"minecraft:brain_coral_fan":                    true,
	"minecraft:brain_coral_wall_fan":               true,
	"minecraft:brewing_stand":                      true,
	"minecraft:brown_banner":                       true,
	"minecraft:brown_bed":                          true,
	"minecraft:brown_candle":                       true,
	"minecraft:brown_candle_cake":                  true,
	"minecraft:brown_carpet":                       true,
	"minecraft:brown_mushroom":                     true,
	"minecraft:brown_wall_banner":                  true,
	"minecraft:bubble_coral":                       true,
	"minecraft:bubble_coral_fan":                   true,
	"minecraft:bubble_coral_wall_fan":              true,
	"minecraft:cake":                               true,
	"minecraft:candle":                             true,
	"minecraft:candle_cake":                        true,
	"minecraft:cave_vines_plant":                   true,
	"minecraft:cave_vines":                         true,
	"minecraft:chain":                              true,
	"minecraft:chipped_anvil":                      true,
	"minecraft:cobweb":                             true,
	"minecraft:cocoa":                              true,
	"minecraft:comparator":                         true,
	"minecraft:conduit":                            true,
	"minecraft:cornflower":                         true,
	"minecraft:creeper_head":                       true,
	"minecraft:creeper_wall_head":                  true,
	"minecraft:crimson_button":                     true,
	"minecraft:crimson_door":                       true,
	"minecraft:crimson_fungus":                     true,
	"minecraft:crimson_pressure_plate":             true,
	"minecraft:crimson_roots":                      true,
	"minecraft:crimson_sign":                       true,
	"minecraft:crimson_trapdoor":                   true,
	"minecraft:crimson_wall_sign":                  true,
	"minecraft:cyan_banner":                        true,
	"minecraft:cyan_bed":                           true,
	"minecraft:cyan_candle":                        true,
	"minecraft:cyan_candle_cake":                   true,
	"minecraft:cyan_carpet":                        true,
	"minecraft:cyan_wall_banner":                   true,
	"minecraft:dandelion":                          true,
	"minecraft:damaged_anvil":                      true,
	"minecraft:dark_oak_button":                    true,
	"minecraft:dark_oak_door":                      true,
	"minecraft:dark_oak_pressure_plate":            true,
	"minecraft:dark_oak_sapling":                   true,
	"minecraft:dark_oak_sign":                      true,
	"minecraft:dark_oak_trapdoor":                  true,
	"minecraft:dark_oak_wall_sign":                 true,
	"minecraft:dead_bubble_coral":                  true,
	"minecraft:dead_bubble_coral_fan":              true,
	"minecraft:dead_bubble_coral_wall_fan":         true,
	"minecraft:dead_brain_coral_wall_fan":          true,
	"minecraft:dead_brain_coral_fan":               true,
	"minecraft:dead_brain_coral":                   true,
	"minecraft:dead_bush":                          true,
	"minecraft:dead_horn_coral":                    true,
	"minecraft:dead_horn_coral_wall_fan":           true,
	"minecraft:dead_horn_coral_fan":                true,
	"minecraft:dead_tube_coral":                    true,
	"minecraft:dead_tube_coral_fan":                true,
	"minecraft:dead_tube_coral_wall_fan":           true,
	"minecraft:dead_fire_coral":                    true,
	"minecraft:dead_fire_coral_fan":                true,
	"minecraft:dead_fire_coral_wall_fan":           true,
	"minecraft:detector_rail":                      true,
	"minecraft:dragon_egg":                         true,
	"minecraft:dragon_head":                        true,
	"minecraft:dragon_wall_head":                   true,
	"minecraft:enchanting_table":                   true,
	"minecraft:end_rod":                            true,
	"minecraft:fern":                               true,
	"minecraft:fire":                               true,
	"minecraft:fire_coral":                         true,
	"minecraft:fire_coral_fan":                     true,
	"minecraft:fire_coral_wall_fan":                true,
	"minecraft:flower_pot":                         true,
	"minecraft:flowering_azalea":                   true,
	"minecraft:grass":                              true,
	"minecraft:gray_banner":                        true,
	"minecraft:gray_bed":                           true,
	"minecraft:gray_candle":                        true,
	"minecraft:gray_candle_cake":                   true,
	"minecraft:gray_carpet":                        true,
	"minecraft:gray_wall_banner":                   true,
	"minecraft:green_banner":                       true,
	"minecraft:green_bed":                          true,
	"minecraft:green_candle":                       true,
	"minecraft:green_candle_cake":                  true,
	"minecraft:green_carpet":                       true,
	"minecraft:green_wall_banner":                  true,
	"minecraft:grindstone":                         true,
	"minecraft:hanging_roots":                      true,
	"minecraft:heavy_weighted_pressure_plate":      true,
	"minecraft:hopper":                             true,
	"minecraft:horn_coral":                         true,
	"minecraft:horn_coral_fan":                     true,
	"minecraft:horn_coral_wall_fan":                true,
	"minecraft:iron_bars":                          true,
	"minecraft:iron_door":                          true,
	"minecraft:iron_trapdoor":                      true,
	"minecraft:jungle_button":                      true,
	"minecraft:jungle_door":                        true,
	"minecraft:jungle_pressure_plate":              true,
	"minecraft:jungle_sapling":                     true,
	"minecraft:jungle_sign":                        true,
	"minecraft:jungle_trapdoor":                    true,
	"minecraft:jungle_wall_sign":                   true,
	"minecraft:kelp":                               true,
	"minecraft:kelp_plant":                         true,
	"minecraft:ladder":                             true,
	"minecraft:lantern":                            true,
	"minecraft:large_amethyst_bud":                 true,
	"minecraft:large_fern":                         true,
	"minecraft:lectern":                            true,
	"minecraft:level":                              true,
	"minecraft:lever":                              true,
	"minecraft:light":                              true,
	"minecraft:light_blue_banner":                  true,
	"minecraft:light_blue_bed":                     true,
	"minecraft:light_blue_candle":                  true,
	"minecraft:light_blue_candle_cake":             true,
	"minecraft:light_blue_carpet":                  true,
	"minecraft:light_blue_wall_banner":             true,
	"minecraft:light_gray_banner":                  true,
	"minecraft:light_gray_bed":                     true,
	"minecraft:light_gray_candle":                  true,
	"minecraft:light_gray_candle_cake":             true,
	"minecraft:light_gray_wall_banner":             true,
	"minecraft:light_weighted_pressure_plate":      true,
	"minecraft:lightning_rod":                      true,
	"minecraft:lilac":                              true,
	"minecraft:lily_of_the_valley":                 true,
	"minecraft:lily_pad":                           true,
	"minecraft:lime_banner":                        true,
	"minecraft:lime_bed":                           true,
	"minecraft:lime_candle":                        true,
	"minecraft:lime_candle_cake":                   true,
	"minecraft:lime_carpet":                        true,
	"minecraft:lime_wall_banner":                   true,
	"minecraft:magenta_banner":                     true,
	"minecraft:magenta_bed":                        true,
	"minecraft:magenta_candle":                     true,
	"minecraft:magenta_candle_cake":                true,
	"minecraft:magenta_carpet":                     true,
	"minecraft:magenta_wall_banner":                true,
	"minecraft:medium_amethyst_bud":                true,
	"minecraft:melon":                              true,
	"minecraft:melon_stem":                         true,
	"minecraft:moss_carpet":                        true,
	"minecraft:moving_piston":                      true,
	"minecraft:nether_sprouts":                     true,
	"minecraft:nether_wart":                        true,
	"minecraft:oak_button":                         true,
	"minecraft:oak_door":                           true,
	"minecraft:oak_pressure_plate":                 true,
	"minecraft:oak_sapling":                        true,
	"minecraft:oak_sign":                           true,
	"minecraft:oak_trapdoor":                       true,
	"minecraft:oak_wall_sign ":                     true,
	"minecraft:oak_wall_sign":                      true,
	"minecraft:orange_banner":                      true,
	"minecraft:orange_bed":                         true,
	"minecraft:orange_candle":                      true,
	"minecraft:orange_candle_cake":                 true,
	"minecraft:orange_carpet":                      true,
	"minecraft:orange_tulip":                       true,
	"minecraft:orange_wall_banner":                 true,
	"minecraft:oxeye_daisy":                        true,
	"minecraft:peony":                              true,
	"minecraft:pink_banner":                        true,
	"minecraft:pink_bed":                           true,
	"minecraft:pink_candle":                        true,
	"minecraft:pink_candle_cake":                   true,
	"minecraft:pink_carpet":                        true,
	"minecraft:pink_tulip":                         true,
	"minecraft:pink_wall_banner":                   true,
	"minecraft:piston":                             true,
	"minecraft:piston_head":                        true,
	"minecraft:player_head":                        true,
	"minecraft:player_wall_head":                   true,
	"minecraft:polished_blackstone_button":         true,
	"minecraft:poppy":                              true,
	"minecraft:pointed_dripstone":                  true,
	"minecraft:polished_blackstone_pressure_plate": true,
	"minecraft:potted_acacia_sapling":              true,
	"minecraft:potted_allium":                      true,
	"minecraft:potted_azure_bluet":                 true,
	"minecraft:potted_bamboo":                      true,
	"minecraft:potted_birch_sapling":               true,
	"minecraft:potted_blue_orchid":                 true,
	"minecraft:potted_brown_mushroom":              true,
	"minecraft:potted_cactus":                      true,
	"minecraft:potted_cornflower":                  true,
	"minecraft:potted_crimson_fungus":              true,
	"minecraft:potted_crimson_roots":               true,
	"minecraft:potted_dark_oak_sapling":            true,
	"minecraft:potted_dandelion":                   true,
	"minecraft:potted_dead_bush":                   true,
	"minecraft:potted_fern":                        true,
	"minecraft:potted_jungle_sapling":              true,
	"minecraft:potted_lily_of_the_valley":          true,
	"minecraft:potted_oak_sapling":                 true,
	"minecraft:potted_orange_tulip":                true,
	"minecraft:potted_oxeye_daisy":                 true,
	"minecraft:potted_pink_tulip":                  true,
	"minecraft:potted_poppy":                       true,
	"minecraft:potted_red_mushroom":                true,
	"minecraft:potted_red_tulip":                   true,
	"minecraft:potted_spruce_sapling":              true,
	"minecraft:potted_warped_fungus":               true,
	"minecraft:potted_warped_roots":                true,
	"minecraft:potted_wither_rose":                 true,
	"minecraft:potted_white_tulip":                 true,
	"minecraft:powered_rail":                       true,
	"minecraft:pumpkin_stem":                       true,
	"minecraft:purple_banner":                      true,
	"minecraft:purple_bed":                         true,
	"minecraft:purple_candle":                      true,
	"minecraft:purple_candle_cake":                 true,
	"minecraft:purple_carpet":                      true,
	"minecraft:purple_wall_banner":                 true,
	"minecraft:rail":                               true,
	"minecraft:red_banner":                         true,
	"minecraft:red_bed":                            true,
	"minecraft:red_candle":                         true,
	"minecraft:red_candle_cake":                    true,
	"minecraft:red_carpet":                         true,
	"minecraft:red_mushroom":                       true,
	"minecraft:red_tulip":                          true,
	"minecraft:red_wall_banner":                    true,
	"minecraft:redstone_lamp":                      true,
	"minecraft:redstone_torch":                     true,
	"minecraft:redstone_wall_torch":                true,
	"minecraft:redstone_wire":                      true,
	"minecraft:repeater":                           true,
	"minecraft:rose_bush":                          true,
	"minecraft:scaffolding":                        true,
	"minecraft:sculk_sensor":                       true,
	"minecraft:sea_pickle":                         true,
	"minecraft:seagrass":                           true,
	"minecraft:skeleton_skull":                     true,
	"minecraft:skeleton_wall_skull":                true,
	"minecraft:snow":                               true,
	"minecraft:soul_fire":                          true,
	"minecraft:small_amethyst_bud":                 true,
	"minecraft:small_dripleaf":                     true,
	"minecraft:soul_campfire":                      true,
	"minecraft:soul_lantern":                       true,
	"minecraft:soul_torch":                         true,
	"minecraft:soul_wall_torch":                    true,
	"minecraft:spore_blossom":                      true,
	"minecraft:spruce_button":                      true,
	"minecraft:spruce_door":                        true,
	"minecraft:spruce_pressure_plate":              true,
	"minecraft:spruce_sign":                        true,
	"minecraft:spruce_trapdoor":                    true,
	"minecraft:spruce_sapling":                     true,
	"minecraft:spruce_wall_sign":                   true,
	"minecraft:stone_button":                       true,
	"minecraft:stone_pressure_plate":               true,
	"minecraft:sticky_piston":                      true,
	"minecraft:sugar_cane":                         true,
	"minecraft:sunflower":                          true,
	"minecraft:sweet_berry_bush":                   true,
	"minecraft:tall_grass":                         true,
	"minecraft:tall_seagrass":                      true,
	"minecraft:target":                             true,
	"minecraft:torch":                              true,
	"minecraft:tripwire":                           true,
	"minecraft:tripwire_hook":                      true,
	"minecraft:tube_coral":                         true,
	"minecraft:tube_coral_fan":                     true,
	"minecraft:tube_coral_wall_fan":                true,
	"minecraft:turtle_egg":                         true,
	"minecraft:twisting_vines":                     true,
	"minecraft:twisting_vines_plant":               true,
	"minecraft:vine":                               true,
	"minecraft:wall_torch":                         true,
	"minecraft:warped_button":                      true,
	"minecraft:warped_door":                        true,
	"minecraft:warped_fungus":                      true,
	"minecraft:warped_pressure_plate":              true,
	"minecraft:warped_roots":                       true,
	"minecraft:warped_sign":                        true,
	"minecraft:warped_trapdoor":                    true,
	"minecraft:warped_wall_sign":                   true,
	"minecraft:weeping_vines":                      true,
	"minecraft:weeping_vines_plant":                true,
	"minecraft:white_banner":                       true,
	"minecraft:white_bed":                          true,
	"minecraft:white_candle":                       true,
	"minecraft:white_candle_cake":                  true,
	"minecraft:white_carpet":                       true,
	"minecraft:white_tulip":                        true,
	"minecraft:white_wall_banner":                  true,
	"minecraft:wither_rose":                        true,
	"minecraft:wither_skeleton_skull":              true,
	"minecraft:wither_skeleton_wall_skull":         true,
	"minecraft:yellow_banner":                      true,
	"minecraft:yellow_bed":                         true,
	"minecraft:yellow_candle":                      true,
	"minecraft:yellow_candle_cake":                 true,
	"minecraft:yellow_carpet":                      true,
	"minecraft:yellow_wall_banner":                 true,
	"minecraft:zombie_head":                        true,
	"minecraft:zombie_wall_head":                   true,
}

// isFoliage returns if a block is treated as leaves by the converter
func (b BlockState) isFoliage() bool {
	return foliage[b.Name]
}

// isAir returns if a block is treated as air by the converter
func (b BlockState) isAir() bool {
	return air[b.Name]
}

// isWatery returns whether or not a block can be treated as water by the converter
func (b BlockState) isWatery() bool {
	watery := water[b.Name]
	if !watery {
		waterlogged, exists := b.Properties["waterlogged"]
		if !exists || waterlogged == "false" || b.isStairs() {
			return false
		}
	}
	return true
}

// isSprite returns whether or not a block can be treated as a sprite by the converter
func (b BlockState) isSprite() bool {
	return sprite[b.Name]
}

// isSnow returns whether or not a block is treated as the thin layer of snow in gmae
// the effect snow has on blocks below it is done in post processing
func (b BlockState) isSnow() bool {
	return b.Name == "minecraft:snow"
}

// isStairs returns whether or not a block is a stair block
func (b BlockState) isStairs() bool {
	return stairs[b.Name]
}

// isSlab returns whether or not a block is a slab block
func (b BlockState) isSlab() bool {
	return slabs[b.Name]
}

// Glows returns whether or not a block is a glow block
func (b BlockState) Glows() bool {
	return glow[b.Name]
}

// isGlass returns whether or not a block is treated as glass by the converter
func (b BlockState) isGlass() bool {
	return glass[b.Name]
}

// isTransparent returns whether or not a block is treated as transparent by the converter.
// This includes all glass, foliage, stairs, and other non-water type "transparent" blocks
func (b BlockState) isTransparent() bool {
	return transparent[b.Name] || b.isGlass() || b.isFoliage() || b.isStairs()
}

// isPortal returns whether or not a block is treated as a portal by the converter.
func (b BlockState) isPortal() bool {
	return portal[b.Name]
}

// isLight returns whether or not a block is treated as a light source by the converter.
func (b BlockState) isLight() bool {
	return light[b.Name]
}

// isIce returns whether or not a block is treated as ice by the converter
func (b BlockState) isIce() bool {
	return ice[b.Name]
}

// isLava returns whether or not a block is treated as lava by the converter
func (b BlockState) isLava() bool {
	return b.Name == "minecraft:lava"
}

// isSkipped returns whether or not a block is skipped for evaluation during the conversion process.
// Blocks are skipped if they are air, or if they are non-watery sprite blocks.
func (b BlockState) isSkipped() bool {
	return b.isAir() ||
		(b.isSprite() && !b.isWatery() && !b.isSnow())
}

// isNotBlockedBy returns whether or not a block is visible if it was completely surrounded
// by another type of block
func (b BlockState) isNotBlockedBy(other BlockState) bool {
	return other.isAir() ||
		(other.isSprite() && !other.isWatery()) || // other block is a sprite that is not watery
		(!b.isWatery() && other.isWatery()) || // this block is not watery and the other block is watery
		(!b.isLava() && other.isLava()) || // this block is not lava and the other block is lava
		other.isTransparent() // other block is see-through in some other way
}
