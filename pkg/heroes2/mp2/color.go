package mp2

import (
	"fmt"
)

type Color int

const (
	Blue Color = iota
	Green
	Red
	Yellow
	Orange
	Purple
)

var colorString = []string{
	"blue",
	"green",
	"red",
	"yellow",
	"orange",
	"purple",
}

func (c Color) String() string { return colorString[c] }

// 6 bytes, each representing "enable" flag for the following colors:
// blue, green, red, yellow, orange, purple
type AllowColors [6]uint8

func (ac AllowColors) Colors() []Color {
	colors := make([]Color, 0, 6)
	for i, boolFlag := range ac {
		if boolFlag > 0 {
			switch i {
			case 0:
				colors = append(colors, Blue)
			case 1:
				colors = append(colors, Green)
			case 2:
				colors = append(colors, Red)
			case 3:
				colors = append(colors, Yellow)
			case 4:
				colors = append(colors, Orange)
			case 5:
				colors = append(colors, Purple)
			}
		}
	}

	return colors
}

func (ac AllowColors) String() string {
	return fmt.Sprintf("%v", ac.Colors())
}
