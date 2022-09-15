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
// faces
var Dom1Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom2Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom3Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom4Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom5Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom6Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom7Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom8Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom9Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom10Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom11Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom12Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom13Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom14Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom15Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom16Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom17Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom18Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom19Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom20Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom21Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom22Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom23Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom24Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom25Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom26Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom27Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom28Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom29Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom30Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom31Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom32Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom33Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom34Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom35Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom36Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom37Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom38Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom39Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom40Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom41Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom42Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom43Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom44Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom45Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom46Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom47Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)
var Dom48Face = NewImageWeb(DOMINO_FACES, 28, 21, 108, 54, false)

// backs
var Dom1Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom2Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom3Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom4Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom5Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom6Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom7Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom8Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom9Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom10Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom11Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom12Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom13Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom14Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom15Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom16Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom17Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom18Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom19Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom20Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom21Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom22Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom23Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom24Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom25Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom26Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom27Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom28Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom29Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom30Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom31Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom32Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom33Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom34Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom35Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom36Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom37Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom38Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom39Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom40Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom41Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom42Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom43Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom44Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom45Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom46Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom47Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
var Dom48Back = NewImageWeb(DOMINO_BACKS, 28, 21, 108, 54, false)
