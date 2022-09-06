package kingdomino

import "image/color"

var colours = new(ColourRegistry)

type Tile struct {
	dimensions Dimensions
	landType   LandType
	crowns     int
	colour     color.RGBA
}

func newTile(landType LandType, crowns int, colour color.RGBA) *Tile {
	t := new(Tile)
	t.landType = landType
	t.crowns = crowns
	t.colour = colour

	t.dimensions = *newDimensions(5, 50, 50)

	return t
}

var WHEAT_ZERO = newTile(WHEAT, 0, colours.WheatGold)
var WHEAT_ONE = newTile(WHEAT, 1, colours.WheatGold)

var PASTURE_ZERO = newTile(PASTURE, 0, colours.PastureGreen)
var PASTURE_ONE = newTile(PASTURE, 1, colours.PastureGreen)
var PASTURE_TWO = newTile(PASTURE, 2, colours.PastureGreen)

var FOREST_ZERO = newTile(FOREST, 0, colours.ForestGreen)
var FOREST_ONE = newTile(FOREST, 1, colours.ForestGreen)

var LAKE_ZERO = newTile(LAKE, 0, colours.LakeBlue)
var LAKE_ONE = newTile(LAKE, 1, colours.LakeBlue)

var SWAMP_ZERO = newTile(SWAMP, 0, colours.SwampBrown)
var SWAMP_ONE = newTile(SWAMP, 1, colours.SwampBrown)
var SWAMP_TWO = newTile(SWAMP, 2, colours.SwampBrown)

var MINE_ZERO = newTile(MINE, 0, colours.MineGrey)
var MINE_ONE = newTile(MINE, 1, colours.MineGrey)
var MINE_TWO = newTile(MINE, 2, colours.MineGrey)
var MINE_THREE = newTile(MINE, 3, colours.MineGrey)
