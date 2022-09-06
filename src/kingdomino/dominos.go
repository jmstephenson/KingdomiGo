package kingdomino

import (
	"math/rand"
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

	d = append(d, newDomino(*WHEAT_ZERO, *WHEAT_ZERO, 1))
	d = append(d, newDomino(*WHEAT_ZERO, *WHEAT_ZERO, 2))
	d = append(d, newDomino(*FOREST_ZERO, *FOREST_ZERO, 3))
	d = append(d, newDomino(*FOREST_ZERO, *FOREST_ZERO, 4))
	d = append(d, newDomino(*FOREST_ZERO, *FOREST_ZERO, 5))
	d = append(d, newDomino(*FOREST_ZERO, *FOREST_ZERO, 6))
	d = append(d, newDomino(*LAKE_ZERO, *LAKE_ZERO, 7))
	d = append(d, newDomino(*LAKE_ZERO, *LAKE_ZERO, 8))

	d = append(d, newDomino(*LAKE_ZERO, *LAKE_ZERO, 9))
	d = append(d, newDomino(*PASTURE_ZERO, *PASTURE_ZERO, 10))
	d = append(d, newDomino(*PASTURE_ZERO, *PASTURE_ZERO, 11))
	d = append(d, newDomino(*SWAMP_ZERO, *SWAMP_ZERO, 12))
	d = append(d, newDomino(*WHEAT_ZERO, *FOREST_ZERO, 13))
	d = append(d, newDomino(*WHEAT_ZERO, *LAKE_ZERO, 14))
	d = append(d, newDomino(*WHEAT_ZERO, *PASTURE_ZERO, 15))
	d = append(d, newDomino(*WHEAT_ZERO, *SWAMP_ZERO, 16))

	d = append(d, newDomino(*FOREST_ZERO, *LAKE_ZERO, 17))
	d = append(d, newDomino(*FOREST_ZERO, *PASTURE_ZERO, 18))
	d = append(d, newDomino(*WHEAT_ONE, *FOREST_ZERO, 19))
	d = append(d, newDomino(*WHEAT_ONE, *LAKE_ZERO, 20))
	d = append(d, newDomino(*WHEAT_ONE, *PASTURE_ZERO, 21))
	d = append(d, newDomino(*WHEAT_ONE, *SWAMP_ZERO, 22))
	d = append(d, newDomino(*WHEAT_ONE, *MINE_ZERO, 23))
	d = append(d, newDomino(*FOREST_ONE, *WHEAT_ZERO, 24))

	d = append(d, newDomino(*FOREST_ONE, *WHEAT_ZERO, 25))
	d = append(d, newDomino(*FOREST_ONE, *WHEAT_ZERO, 26))
	d = append(d, newDomino(*FOREST_ONE, *WHEAT_ZERO, 27))
	d = append(d, newDomino(*FOREST_ONE, *LAKE_ZERO, 28))
	d = append(d, newDomino(*FOREST_ONE, *PASTURE_ZERO, 29))
	d = append(d, newDomino(*LAKE_ONE, *WHEAT_ZERO, 30))
	d = append(d, newDomino(*LAKE_ONE, *WHEAT_ZERO, 31))
	d = append(d, newDomino(*LAKE_ONE, *FOREST_ZERO, 32))

	d = append(d, newDomino(*LAKE_ONE, *FOREST_ZERO, 33))
	d = append(d, newDomino(*LAKE_ONE, *FOREST_ZERO, 34))
	d = append(d, newDomino(*LAKE_ONE, *FOREST_ZERO, 35))
	d = append(d, newDomino(*WHEAT_ZERO, *PASTURE_ONE, 36))
	d = append(d, newDomino(*LAKE_ZERO, *PASTURE_ONE, 37))
	d = append(d, newDomino(*WHEAT_ZERO, *SWAMP_ONE, 38))
	d = append(d, newDomino(*PASTURE_ZERO, *SWAMP_ONE, 39))
	d = append(d, newDomino(*MINE_ONE, *WHEAT_ZERO, 40))

	d = append(d, newDomino(*WHEAT_ZERO, *PASTURE_TWO, 41))
	d = append(d, newDomino(*LAKE_ZERO, *PASTURE_TWO, 42))
	d = append(d, newDomino(*WHEAT_ZERO, *SWAMP_TWO, 43))
	d = append(d, newDomino(*PASTURE_ZERO, *SWAMP_TWO, 44))
	d = append(d, newDomino(*MINE_TWO, *WHEAT_ZERO, 45))
	d = append(d, newDomino(*SWAMP_ZERO, *MINE_TWO, 46))
	d = append(d, newDomino(*SWAMP_ZERO, *MINE_TWO, 47))
	d = append(d, newDomino(*WHEAT_ZERO, *MINE_THREE, 48))

	deck := new(Deck)
	deck.Dominos = d
	return deck
}

func (d *Deck) ShuffleDeck() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	shuffled := make([]*Domino, TOTAL_TILES)

	for len(d.Dominos) > 0 {
		position := r1.Intn(len(d.Dominos))
		domino := d.Dominos[position]
		d.Dominos = remove(d.Dominos, position)
		shuffled[TOTAL_TILES-1-len(d.Dominos)] = domino
	}

	d.Dominos = shuffled
}

func remove(deck []*Domino, pos int) []*Domino {
	return append(deck[:pos], deck[pos+1:]...)
}
