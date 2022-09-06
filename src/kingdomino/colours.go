package kingdomino

import (
	"errors"
	"image/color"
	"log"
)

var errInvalidFormat = errors.New("invalid format")
var errPurple = color.RGBA{R: 145, G: 2, B: 227, A: 1}
var Colours = newColourRegistry()

func newColourRegistry() *ColourRegistry {
	return &ColourRegistry{
		ErrorPurple:  errPurple,
		PastureGreen: ParseHexColorSafe(PASTURE_GREEN),
		WheatGold:    ParseHexColorSafe(WHEAT_GOLD),
		LakeBlue:     ParseHexColorSafe(LAKE_BLUE),
		MineGrey:     ParseHexColorSafe(MINE_GREY),
		ForestGreen:  ParseHexColorSafe(FOREST_GREEN),
		SwampBrown:   ParseHexColorSafe(SWAMP_BROWN),
	}
}

type ColourRegistry struct {
	ErrorPurple  color.RGBA
	PastureGreen color.RGBA
	WheatGold    color.RGBA
	LakeBlue     color.RGBA
	MineGrey     color.RGBA
	ForestGreen  color.RGBA
	SwampBrown   color.RGBA
}

func ParseHexColorSafe(s string) (c color.RGBA) {
	hex, err := ParseHexColor(s)
	if err != nil {
		log.Print(err)
		return errPurple
	} else {
		return hex
	}
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
