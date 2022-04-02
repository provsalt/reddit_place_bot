package placer

import "image/color"

type Color struct {
	color.RGBA
	color int
}

var colourw map[string]Color = map[string]Color{
	"black":       Color{RGBA: color.RGBA{A: 255}, color: ColorBlack},
	"white":       Color{RGBA: color.RGBA{R: 255, G: 255, B: 255, A: 255}, color: ColorWhite},
	"red":         Color{RGBA: color.RGBA{R: 255, G: 69, A: 255}, color: ColorRed},
	"light_green": Color{RGBA: color.RGBA{R: 126, G: 237, B: 86, A: 255}, color: ColorLightGreen},
	"dark_green":  Color{RGBA: color.RGBA{G: 163, B: 104, A: 255}, color: ColorDarkGreen},
	"light_blue":  Color{color.RGBA{B: 255, A: 255}, 4},
	"blue":        Color{color.RGBA{B: 255, A: 255}, 4},
	"dark_blue":   Color{color.RGBA{B: 255, A: 255}, 4},
	"yellow":      Color{color.RGBA{R: 255, G: 255, A: 255}, 5},
	"orange":      Color{color.RGBA{R: 255, G: 128, A: 255}, 6},
	"purple":      Color{color.RGBA{R: 128, B: 255, A: 255}, 7},
	"pink":        Color{color.RGBA{R: 255, B: 255, A: 255}, 8},
	"brown":       Color{color.RGBA{R: 128, G: 64, A: 255}, 9},
}

const (
	ColorRed = iota + 2
	ColorOrange
	ColorYellow
	ColorDarkGreen  = 6
	ColorLightGreen = 8
	ColorDarkBlue   = 12
	ColorBlue       = 13
	ColorLightBlue  = 14
	ColorDarkPurple = 18
	ColorPurple     = 19
	ColorLightPink  = 23
	ColorBrown      = 25
	ColorBlack      = 27
	ColorGray       = 29
	ColorWhite      = 31
)
