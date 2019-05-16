package agg

import (
	"bytes"
	"encoding/binary"
	"image"
)

// Tiles implement image.Image interface.
//var tiles image.Image = Tiles{}

func NewTiles(data []byte, palette palette) *Tiles {
	return &Tiles{
		data:    data,
		palette: palette,
	}
}

type Tiles struct {
	// 0x00  Number of tiles  2 bytes
	// 0x02  Tile width       2 bytes
	// 0x04  Tile height      2 bytes
	// 0x06  Data
	data []byte

	palette palette
}

func (t *Tiles) NumTiles() int   { return t.uint16ToInt(0, 2) }
func (t *Tiles) TileWidth() int  { return t.uint16ToInt(2, 4) }
func (t *Tiles) TileHeight() int { return t.uint16ToInt(4, 6) }

func (t *Tiles) uint16ToInt(from, to int) int {
	var v uint16
	_ = binary.Read(bytes.NewReader(t.data[from:to]), binary.LittleEndian, &v)
	return int(v)
}

const opaqueAlpha = uint8(255)

func (t *Tiles) Images() []*image.RGBA {
	data := t.data[6:] // Pixels only, strip off the header.

	numTiles := t.NumTiles()
	width := t.TileWidth()
	height := t.TileHeight()

	// Collect all pixels at once.
	// TODO: Why don't we use a custom image.PalettedImage and defer the .RGBA() call for later?
	pixels := make([]uint8, 0, numTiles*width*height*4)
	for i := 0; i < numTiles*width*height; i++ {
		r, g, b := t.palette.RGB(int(data[i]))
		pixels = append(pixels, r, g, b, opaqueAlpha)
	}

	rect := image.Rect(0, 0, width, height)

	imgs := make([]*image.RGBA, 0, numTiles)
	for i := 0; i < numTiles; i++ {
		img := &image.RGBA{pixels[i*width*height*4 : (i+1)*width*height*4], 4 * width, rect}
		imgs = append(imgs, img)
	}

	return imgs
}
