package kingdomino

type LandType int

const (
	PASTURE LandType = iota
	WHEAT
	LAKE
	MINE
	FOREST
	SWAMP

	TOTAL_TILES = 48

	PASTURE_GREEN = "#6ab051"
	WHEAT_GOLD    = "#f0d351"
	LAKE_BLUE     = "#218fd9"
	MINE_GREY     = "#a1998f"
	FOREST_GREEN  = "#266912"
	SWAMP_BROWN   = "#45392b"
)

type Dimensions struct {
	height int
	width  int
	length int
}

func newDimensions(h, w, l int) *Dimensions {
	d := new(Dimensions)
	d.height = h
	d.width = w
	d.length = l

	return d
}

type Table struct {
	dimensions Dimensions
}

type PlayerBoard struct {
	spaces     [25]Tile
	score      int
	isComplete bool
}
