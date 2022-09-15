package gamestate

import (
	kd "kingdomino/src/kingdomino"
	"math/rand"
	"sort"
	"time"
)

var HumanButtons = map[int]bool{}

type Phase int

const (
	SETUP       Phase = 1
	TILE_PLACE  Phase = 2
	TILE_SELECT Phase = 3
)

type GamePhases struct {
	SETUP       Phase
	TILE_PLACE  Phase
	TILE_SELECT Phase
}

func newGamePhases() *GamePhases {
	gp := new(GamePhases)
	gp.SETUP = SETUP
	gp.TILE_PLACE = TILE_PLACE
	gp.TILE_SELECT = TILE_SELECT

	return gp
}

var Phases = newGamePhases()

type Game struct {
	Height int
	Width  int

	Players         []*Player
	PlayerCount     int
	ActivePlayer    *Player
	PlayerTurnCount int

	TileDeck      *kd.Deck
	DrawnTiles    *kd.Deck
	RevealedTiles *kd.Deck

	Phase Phase
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

func NewGame(width, height int, players []string) *Game {
	g := new(Game)
	g.Phase = Phases.SETUP

	g.Init(width, height)

	for _, p := range players {
		g.AddPlayer(p)
	}

	return g
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
	g.NextPlayer()
}

func (g *Game) RevealTiles() {
	g.RevealedTiles = nil
	for i := 0; i < g.PlayerCount; i++ {
		g.RevealedTiles.Dominos = append(g.RevealedTiles.Dominos, g.DrawnTiles.DrawDomino(0))
	}
}

func (g *Game) DrawTiles() {
	g.DrawnTiles = nil
	for i := 0; i < g.PlayerCount; i++ {
		g.DrawnTiles.Dominos = append(g.DrawnTiles.Dominos, g.TileDeck.DrawDomino(0))
	}
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

func (g *Game) NextPlayer() {
	g.ActivePlayer = g.Players[g.PlayerTurnCount%g.PlayerCount]
}
