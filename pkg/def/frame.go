package def

import (
	"fmt"
	"image"
	"strings"
)

const opaqueAlpha = 255

type Frame struct {
	BlockId int
	Name    string
	Offset  int
	Data    []byte

	Palette *Palette
	Width   int
	Height  int
}

func (frame *Frame) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "frame %v (block %v)", frame.Name, frame.BlockId)

	return b.String()
}

func (f *Frame) Image() image.Image {
	data := f.Data[:]

	// Collect all pixels at once.
	// TODO: Why don't we use a custom image.PalettedImage and defer the .RGBA() call for later?
	pixels := make([]uint8, 0, f.Width*f.Height*4)
	for i := 0; i < f.Width*f.Height; i++ {
		r, g, b := f.Palette.RGB(int(data[i]))
		pixels = append(pixels, r, g, b, opaqueAlpha)
	}

	rect := image.Rect(0, 0, f.Width, f.Height)
	return &image.RGBA{pixels[:f.Width*f.Height*4], 4 * f.Width, rect}
}
