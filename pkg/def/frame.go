package def

import (
	"image"

	"github.com/pkg/errors"
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

func (frame *Frame) Image() (image.Image, error) {
	pixels := make([]uint8, 0, frame.Width*frame.Height*4)

	switch frame.Format {
	case 0:
		// Collect all pixels at once.
		// NOTE: Would it be possible to use a custom image.PalettedImage and defer the .RGBA() call for later?
		for i := 0; i < frame.Width*frame.Height; i++ {
			r, g, b := frame.Palette.RGB(int(frame.Data[i]))
			pixels = append(pixels, r, g, b, opaqueAlpha)
		}

	case 3:

	default:
		return nil, errors.Errorf("unsupported format %v", frame.Format)
	}

	rect := image.Rect(0, 0, frame.Width, frame.Height)
	img := &image.RGBA{pixels[:frame.Width*frame.Height*4], 4 * frame.Width, rect}

	return img, nil
}
