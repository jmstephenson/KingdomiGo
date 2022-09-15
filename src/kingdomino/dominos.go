package kingdomino

import (
	"math/rand"
	"sort"
	"time"
)

type Domino struct {
	dimensions Dimensions
	face_a     Tile
	face_b     Tile
	Value      int
	rotation   int
	FaceImage  ImageWeb
	BackImage  ImageWeb
}

func newDomino(face_a, face_b Tile, value int, face ImageWeb, back ImageWeb) *Domino {
	d := new(Domino)
	d.face_a = face_a
	d.face_b = face_b
	d.Value = value
	d.FaceImage = face
	d.BackImage = back

	return d
}

type Deck struct {
	Dominos []*Domino
}

func NewDeck() *Deck {
	d := make([]*Domino, 0)

	d = append(d, newDomino(*WHEAT_ZERO, *WHEAT_ZERO, 1, *Dom1Face, *Dom1Back))
	d = append(d, newDomino(*WHEAT_ZERO, *WHEAT_ZERO, 2, *Dom2Face, *Dom2Back))
	d = append(d, newDomino(*FOREST_ZERO, *FOREST_ZERO, 3, *Dom3Face, *Dom3Back))
	d = append(d, newDomino(*FOREST_ZERO, *FOREST_ZERO, 4, *Dom4Face, *Dom4Back))
	d = append(d, newDomino(*FOREST_ZERO, *FOREST_ZERO, 5, *Dom5Face, *Dom5Back))
	d = append(d, newDomino(*FOREST_ZERO, *FOREST_ZERO, 6, *Dom6Face, *Dom6Back))
	d = append(d, newDomino(*LAKE_ZERO, *LAKE_ZERO, 7, *Dom7Face, *Dom7Back))
	d = append(d, newDomino(*LAKE_ZERO, *LAKE_ZERO, 8, *Dom8Face, *Dom8Back))

	d = append(d, newDomino(*LAKE_ZERO, *LAKE_ZERO, 9, *Dom9Face, *Dom9Back))
	d = append(d, newDomino(*PASTURE_ZERO, *PASTURE_ZERO, 10, *Dom10Face, *Dom10Back))
	d = append(d, newDomino(*PASTURE_ZERO, *PASTURE_ZERO, 11, *Dom11Face, *Dom11Back))
	d = append(d, newDomino(*SWAMP_ZERO, *SWAMP_ZERO, 12, *Dom12Face, *Dom12Back))
	d = append(d, newDomino(*WHEAT_ZERO, *FOREST_ZERO, 13, *Dom13Face, *Dom13Back))
	d = append(d, newDomino(*WHEAT_ZERO, *LAKE_ZERO, 14, *Dom14Face, *Dom14Back))
	d = append(d, newDomino(*WHEAT_ZERO, *PASTURE_ZERO, 15, *Dom15Face, *Dom15Back))
	d = append(d, newDomino(*WHEAT_ZERO, *SWAMP_ZERO, 16, *Dom16Face, *Dom16Back))

	d = append(d, newDomino(*FOREST_ZERO, *LAKE_ZERO, 17, *Dom17Face, *Dom17Back))
	d = append(d, newDomino(*FOREST_ZERO, *PASTURE_ZERO, 18, *Dom18Face, *Dom18Back))
	d = append(d, newDomino(*WHEAT_ONE, *FOREST_ZERO, 19, *Dom19Face, *Dom19Back))
	d = append(d, newDomino(*WHEAT_ONE, *LAKE_ZERO, 20, *Dom20Face, *Dom20Back))
	d = append(d, newDomino(*WHEAT_ONE, *PASTURE_ZERO, 21, *Dom21Face, *Dom21Back))
	d = append(d, newDomino(*WHEAT_ONE, *SWAMP_ZERO, 22, *Dom22Face, *Dom22Back))
	d = append(d, newDomino(*WHEAT_ONE, *MINE_ZERO, 23, *Dom23Face, *Dom23Back))
	d = append(d, newDomino(*FOREST_ONE, *WHEAT_ZERO, 24, *Dom24Face, *Dom24Back))

	d = append(d, newDomino(*FOREST_ONE, *WHEAT_ZERO, 25, *Dom25Face, *Dom25Back))
	d = append(d, newDomino(*FOREST_ONE, *WHEAT_ZERO, 26, *Dom26Face, *Dom26Back))
	d = append(d, newDomino(*FOREST_ONE, *WHEAT_ZERO, 27, *Dom27Face, *Dom27Back))
	d = append(d, newDomino(*FOREST_ONE, *LAKE_ZERO, 28, *Dom28Face, *Dom28Back))
	d = append(d, newDomino(*FOREST_ONE, *PASTURE_ZERO, 29, *Dom29Face, *Dom29Back))
	d = append(d, newDomino(*LAKE_ONE, *WHEAT_ZERO, 30, *Dom30Face, *Dom30Back))
	d = append(d, newDomino(*LAKE_ONE, *WHEAT_ZERO, 31, *Dom31Face, *Dom31Back))
	d = append(d, newDomino(*LAKE_ONE, *FOREST_ZERO, 32, *Dom32Face, *Dom32Back))

	d = append(d, newDomino(*LAKE_ONE, *FOREST_ZERO, 33, *Dom33Face, *Dom33Back))
	d = append(d, newDomino(*LAKE_ONE, *FOREST_ZERO, 34, *Dom34Face, *Dom34Back))
	d = append(d, newDomino(*LAKE_ONE, *FOREST_ZERO, 35, *Dom35Face, *Dom35Back))
	d = append(d, newDomino(*WHEAT_ZERO, *PASTURE_ONE, 36, *Dom36Face, *Dom36Back))
	d = append(d, newDomino(*LAKE_ZERO, *PASTURE_ONE, 37, *Dom37Face, *Dom37Back))
	d = append(d, newDomino(*WHEAT_ZERO, *SWAMP_ONE, 38, *Dom38Face, *Dom38Back))
	d = append(d, newDomino(*PASTURE_ZERO, *SWAMP_ONE, 39, *Dom39Face, *Dom39Back))
	d = append(d, newDomino(*MINE_ONE, *WHEAT_ZERO, 40, *Dom40Face, *Dom40Back))

	d = append(d, newDomino(*WHEAT_ZERO, *PASTURE_TWO, 41, *Dom41Face, *Dom41Back))
	d = append(d, newDomino(*LAKE_ZERO, *PASTURE_TWO, 42, *Dom42Face, *Dom42Back))
	d = append(d, newDomino(*WHEAT_ZERO, *SWAMP_TWO, 43, *Dom43Face, *Dom43Back))
	d = append(d, newDomino(*PASTURE_ZERO, *SWAMP_TWO, 44, *Dom44Face, *Dom44Back))
	d = append(d, newDomino(*MINE_TWO, *WHEAT_ZERO, 45, *Dom45Face, *Dom45Back))
	d = append(d, newDomino(*SWAMP_ZERO, *MINE_TWO, 46, *Dom46Face, *Dom46Back))
	d = append(d, newDomino(*SWAMP_ZERO, *MINE_TWO, 47, *Dom47Face, *Dom47Back))
	d = append(d, newDomino(*WHEAT_ZERO, *MINE_THREE, 48, *Dom47Face, *Dom48Back))

	deck := new(Deck)
	deck.Dominos = d
	return deck
}

func NewDeckSelection(dominos []*Domino) *Deck {
	d := new(Deck)
	d.Dominos = dominos

	return d
}

func (d *Deck) ShuffleDeck() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	shuffled := make([]*Domino, TOTAL_TILES)

	for len(d.Dominos) > 0 {
		position := r1.Intn(len(d.Dominos))
		domino := d.Dominos[position]
		_ = d.DrawDomino(position)
		shuffled[TOTAL_TILES-1-len(d.Dominos)] = domino
	}

	d.Dominos = shuffled
}

func (d *Deck) DrawDomino(pos int) *Domino {
	drawnTile := d.Dominos[pos]
	d.Dominos = append(d.Dominos[:pos], d.Dominos[pos+1:]...)
	return drawnTile
}

func (d *Deck) SortByValue(deck []*Domino) {
	sort.Slice(d.Dominos, func(i, j int) bool {
		return d.Dominos[i].Value < d.Dominos[j].Value
	})
}

func (d *Deck) TopCard() *Domino {
	return d.Dominos[0]
}
