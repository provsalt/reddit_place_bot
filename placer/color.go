package placer

import "image/color"

type Color struct {
	color.RGBA
	color int
}

var colors = map[int]color.Palette{
	0: []color.Color{
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
	},
}

const (
	ColorDarkRed = iota
	ColorRed
	ColorOrange
	ColorYellow

	ColorDarkGreen = iota + 6
	ColorGreen
	ColorLightGreen
	ColorDarkTeal
	ColorTeal

	ColorDarkBlue = iota + 12
	ColorBlue
	ColorLightBlue
	ColorIndigo
	ColorPeriwinkle

	ColorDarkPurple = 18
	ColorPurple     = 19
	ColorPink       = 22
	ColorLightPink  = 23
	ColorDarkBrown  = 24
	ColorBrown      = 25

	ColorBlack     = 27
	ColorGray      = 29
	ColorLightGray = 30
	ColorWhite     = 31
)

// hexToColor converts hex color string to rgba
func hexToColor(hex string) color.RGBA {
	var r, g, b, a uint8
	if len(hex) == 7 {
		r = uint8(hex[1])<<4 | uint8(hex[2])
		g = uint8(hex[3])<<4 | uint8(hex[4])
		b = uint8(hex[5])<<4 | uint8(hex[6])
		a = 255
	} else if len(hex) == 9 {
		r = uint8(hex[1])<<4 | uint8(hex[2])
		g = uint8(hex[3])<<4 | uint8(hex[4])
		b = uint8(hex[5])<<4 | uint8(hex[6])
		a = uint8(hex[7])<<4 | uint8(hex[8])
	}
	return color.RGBA{R: r, G: g, B: b, A: a}
}
