package def

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"strings"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
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

	data []byte
}

func (frame *Frame) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "frame %v (block %v, offset %v, format %v, size %v)", frame.Name, frame.BlockId, frame.Offset, frame.Format, frame.size)

	return b.String()
}

func (frame *Frame) Image() (image.Image, error) {
	pixels := make([]uint8, 0, frame.Width2*frame.Height2*4)

	switch frame.Format {
	case 0:
	// Collect all pixels at once.
	// NOTE: Would it be possible to use a custom image.PalettedImage and defer the .RGBA() call for later?

	// for i := 0; i < frame.Width2*frame.Height2; i++ {
	// 	r, g, b := frame.Palette.RGBA(int(frame.data[i]))
	// 	pixels = append(pixels, r, g, b, opaqueAlpha)
	// }

	case 3:
		get := bytestream.New(bytes.NewReader(frame.data), binary.LittleEndian)

		var lineOffsets []int
		for i := 0; i < frame.Height2; i++ {
			lineOffsets = append(lineOffsets, get.Int(2))
		}
		fmt.Printf("offsets: %#v\n", lineOffsets)

		// read blocks of (32 bytes)
		for _, lineOffset := range lineOffsets {
			get := bytestream.New(bytes.NewReader(frame.data[lineOffset:]), binary.LittleEndian)

			for totalRowLength := 0; totalRowLength < frame.Width2; {
				segment := get.Int(1)
				code := segment / 32
				length := (segment & 31) + 1
				switch code {
				case 7: // Raw data
					for i := 0; i < length; i++ {
						r, g, b, a := frame.Palette.RGBA(get.Int(1))
						pixels = append(pixels, r, g, b, a)
					}
					totalRowLength += length
				default: // RLE
					r, g, b, a := frame.Palette.RGBA(length * code) // TODO: chr(code) ??

					for i := 0; i < length; i++ {
						pixels = append(pixels, r, g, b, a)
					}
					totalRowLength += length
				}

				fmt.Printf("offset: %v, code: %v, length: %v\n", lineOffset, code, length)
			}
		}

	default:
		return nil, errors.Errorf("unsupported format %v", frame.Format)
	}

	// Fill in blank pixels.
	for i := 0; i < cap(pixels)-len(pixels); i++ {
		fmt.Printf("fixing.. blanks..")
		pixels = append(pixels, 255, 255, 255, opaqueAlpha) // White.
	}

	rect := image.Rect(0, 0, frame.Width2, frame.Height2)
	img := &image.RGBA{pixels, 4 * frame.Width2, rect}

	return img, nil
}

func (frame *Frame) PaletteImage() (image.Image, error) {
	pixels := make([]uint8, 0, 256*4)
	for i := 0; i < 256; i++ {
		r, g, b, a := frame.Palette.RGBA(i)
		pixels = append(pixels, r, g, b, a)
	}

	rect := image.Rect(0, 0, 16, 16)
	img := &image.RGBA{pixels, 4 * 16, rect}

	return img, nil
}
