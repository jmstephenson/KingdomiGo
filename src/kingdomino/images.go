package kingdomino

type ImageWeb struct {
	FilePath  string
	StartPosX int
	StartPosY int
	Length    int
	Height    int
	Visible   bool
}

func NewImageWeb(filePath string, startPosX, startPosY, height, length int, visible bool) *ImageWeb {
	i := new(ImageWeb)
	i.FilePath = filePath
	i.StartPosX = startPosX
	i.StartPosY = startPosY
	i.Height = height
	i.Length = length
	i.Visible = true

	return i
}

const DOMINO_FACES string = "images/kingdomino_tile_faces.webp"
const DOMINO_BACKS string = "images/kingdomino_tile_backs.webp"
const PLAYER_CASTLES string = "images/kingdomino_player_castles.webp"

// Dominos
var Dom1Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
