package def

import (
	"image"
)

const opaqueAlpha = 255

type Frame struct {
	BlockId int
	Name    string
	Offset  int

	Palette *Palette
	Width   int
	Height  int

	size        int // unused
	Format      int
	FullWidth   int
	FullHeight  int
	Width2      int
	Height2     int
	LeftMargin  int
	RightMargin int

	Data []byte
}

// func (frame *Frame) String() string {
// 	var b strings.Builder

// 	fmt.Fprintf(&b, "frame %v (block %v, offset %v)", frame.Name, frame.BlockId, frame.Offset)

// 	return b.String()
// }

func (f *Frame) Image() image.Image {
	// Collect all pixels at once.
	// NOTE: Would it be possible to use a custom image.PalettedImage and defer the .RGBA() call for later?
	pixels := make([]uint8, 0, f.Width*f.Height*4)
	for i := 0; i < f.Width*f.Height; i++ {
		r, g, b := f.Palette.RGB(int(f.Data[i]))
		pixels = append(pixels, r, g, b, opaqueAlpha)
	}

	rect := image.Rect(0, 0, f.Width, f.Height)
	return &image.RGBA{pixels[:f.Width*f.Height*4], 4 * f.Width, rect}
}
