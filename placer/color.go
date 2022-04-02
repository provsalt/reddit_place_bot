package placer

import (
	"fmt"
	"image/color"
)

type Color struct {
	color.RGBA
	color int
}

var Colors = color.Palette{
	hexToColor("BE0039"),
	hexToColor("FF4500"),
	hexToColor("FFA800"),
	hexToColor("FFD635"),
	hexToColor("00A368"),
	hexToColor("00CC78"),
	hexToColor("7EED56"),
	hexToColor("00756F"),
	hexToColor("009EAA"),
	hexToColor("2450A4"),
	hexToColor("3690EA"),
	hexToColor("51E9F4"),
	hexToColor("493AC1"),
	hexToColor("6A5CFF"),
	hexToColor("811E9F"),
	hexToColor("B44AC0"),
	hexToColor("FF3881"),
	hexToColor("FF99AA"),
	hexToColor("6D482F"),
	hexToColor("9C6926"),
	hexToColor("000000"),
	hexToColor("898D90"),
	hexToColor("D4D7D9"),
	hexToColor("FFFFFF"),
}

var colorMapping = map[int]color.Color{
	ColorDarkRed:    hexToColor("BE0039"),
	ColorRed:        hexToColor("FF4500"),
	ColorOrange:     hexToColor("FFA800"),
	ColorYellow:     hexToColor("FFD635"),
	ColorDarkGreen:  hexToColor("00A368"),
	ColorGreen:      hexToColor("7EED56"),
	ColorLightGreen: hexToColor("00CC78"),
	ColorDarkTeal:   hexToColor("00756F"),
	ColorTeal:       hexToColor("009EAA"),
	ColorDarkBlue:   hexToColor("2450A4"),
	ColorBlue:       hexToColor("3690EA"),
	ColorLightBlue:  hexToColor("51E9F4"),
	ColorIndigo:     hexToColor("493AC1"),
	ColorPeriwinkle: hexToColor("6A5CFF"),
	ColorDarkPurple: hexToColor("811E9F"),
	ColorPurple:     hexToColor("B44AC0"),
	ColorPink:       hexToColor("FF3881"),
	ColorLightPink:  hexToColor("FF99AA"),
	ColorDarkBrown:  hexToColor("6D482F"),
	ColorBrown:      hexToColor("9C6926"),
	ColorBlack:      hexToColor("000000"),
	ColorGray:       hexToColor("898D90"),
	ColorLightGray:  hexToColor("D4D7D9"),
	ColorWhite:      hexToColor("FFFFFF"),
}

const (
	ColorDarkRed = iota
	ColorRed
	ColorOrange
	ColorYellow

	ColorDarkGreen  = 6
	ColorGreen      = 7
	ColorLightGreen = 8
	ColorDarkTeal   = 9
	ColorTeal       = 10

	ColorDarkBlue   = 12
	ColorBlue       = 13
	ColorLightBlue  = 14
	ColorIndigo     = 15
	ColorPeriwinkle = 16

	ColorDarkPurple = 18
	ColorPurple     = 19

	ColorPink      = 22
	ColorLightPink = 23
	ColorDarkBrown = 24
	ColorBrown     = 25

	ColorBlack     = 27
	ColorGray      = 29
	ColorLightGray = 30
	ColorWhite     = 31
)

// hexToColor converts hex color string to rgba
func hexToColor(s string) (c color.RGBA) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, _ = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, _ = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	}
	return
}
