package mp2

import (
	"fmt"
	"log"
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
type AllowColors [6]Bool

func (ac AllowColors) Colors() []Color {
	colors := make([]Color, 0, 6)
	for i, enable := range ac {
		if enable.Bool() {
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
	log.Println(fmt.Sprintf("%v", colors))
	return colors
}

func (ac AllowColors) String() string {
	return fmt.Sprintf("%v", ac.Colors())
}
