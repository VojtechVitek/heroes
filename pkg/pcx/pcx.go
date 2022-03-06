package pcx

import (
	"encoding/binary"
	"image"
	"io"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
	"github.com/VojtechVitek/heroes/pkg/palette"
	"github.com/pkg/errors"
)

type Pcx struct {
	Size   int
	Width  int
	Height int

	data []byte

	isPaletted bool
	palette    palette.Palette
}

func Parse(r io.ReadSeeker) (*Pcx, error) {
	get := bytestream.New(r, binary.LittleEndian)
	pcx := &Pcx{
		Size:   get.Int(4),
		Width:  get.Int(4),
		Height: get.Int(4),
	}
	pcx.data = get.Bytes(pcx.Size)

	if pcx.Size == pcx.Width*pcx.Height {
		// 8bit format (PCX8B) has palette end the
		pcx.palette = get.Bytes(256 * 3)
	}

	return pcx, get.Error()
}

func (pcx *Pcx) Image() (image.Image, error) {
	pixels := make([]uint8, 0, pcx.Width*pcx.Height*4)

	if len(pcx.palette) == 0 {

		// Raw 24bit RGB pixels.
		for i := 0; i < pcx.Width*pcx.Height; i++ {
			pixels = append(pixels, pcx.data[i*3:i*3+2]...)
			pixels = append(pixels, palette.OpaqueAlpha)
		}

	} else {

		// Paletted 8bit RGB pixels.
		for i := 0; i < pcx.Width*pcx.Height; i++ {
			r, g, b, a := pcx.palette.RGBA(int(pcx.data[i]))
			pixels = append(pixels, r, g, b, a)
		}
	}

	// Fill in blank pixels.
	if cap(pixels)-len(pixels) > 0 {
		return nil, errors.Errorf("failed to parse PCX image (format %v): missing %v pixels", cap(pixels)-len(pixels))
	}
	// for i := 0; i < cap(pixels)-len(pixels); i++ {
	// 	pixels = append(pixels, 255, 255, 255, palette.OpaqueAlpha) // White.
	// }

	rect := image.Rect(0, 0, pcx.Width, pcx.Height)
	img := &image.RGBA{pixels, 4 * pcx.Width, rect}

	return img, nil
}
