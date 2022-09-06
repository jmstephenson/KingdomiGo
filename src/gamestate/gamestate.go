package gamestate

import (
	kd "kingdomino/src/kingdomino"
	"math/rand"
	"sort"
	"time"
)

var HumanButtons = map[int]bool{}

const (
	SETUP       = iota
	TILE_PLACE  = iota
	TILE_SELECT = iota
)

type Game struct {
	Height int
	Width  int

	Players         []*Player
	PlayerCount     int
	ActivePlayer    *Player
	PlayerTurnCount int

	TileDeck *kd.Deck

	Phase GamePhase
}

type Player struct {
	Name          string
	TileToPlace   *kd.Domino
	Board         *kd.PlayerBoard
	LastTileValue int // determines turn order
	TurnTaken     bool
}

func NewPlayer(name string) *Player {
	p := new(Player)
	p.Name = name

	p.Board = new(kd.PlayerBoard)

	return p
}

type GamePhase struct {
}

func (g *Game) Init(width, height int) {
	g.Height = height
	g.Width = width

	g.Players = make([]*Player, 0)
	g.TileDeck = kd.NewDeck()

}

func (g *Game) InitGame() {
	// set up deck
	g.TileDeck.ShuffleDeck()

	// set up players
	g.AddPlayer("Alice")
	g.AddPlayer("Bob")

	g.PlayerCount = len(g.Players)
	g.PlayerTurnCount = 0

	g.RandomisePlayerOrder()
	g.SetActivePlayer()

}

func (g *Game) AddPlayer(name string) {
	g.Players = append(g.Players, NewPlayer(name))
}

func (g *Game) Update() {

}

func (g *Game) RandomisePlayerOrder() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for _, player := range g.Players {
		player.LastTileValue = r1.Intn(kd.TOTAL_TILES)
	}

	g.SortPlayerOrder()
}

func (g *Game) SortPlayerOrder() {
	sort.Slice(g.Players, func(i, j int) bool {
		return g.Players[i].LastTileValue < g.Players[j].LastTileValue
	})
}

func (g *Game) SetActivePlayer() {
	g.ActivePlayer = g.Players[g.PlayerTurnCount%g.PlayerCount]
}
