package agg

import (
	"bytes"
	"encoding/binary"
	"image"
)

// Tiles implement image.Image interface.
//var tiles image.Image = Tiles{}
const alpha = uint8(255)

func NewTiles(data []byte, pallete Pallete) *Tiles {
	return &Tiles{
		data:    data,
		pallete: pallete,
	}
}

type Tiles struct {
	// 0x00  Number of tiles  2 bytes
	// 0x02  Tile width       2 bytes
	// 0x04  Tile height      2 bytes
	// 0x06  Data
	data []byte

	pallete Pallete
}

func (t *Tiles) NumTiles() int   { return t.uint16ToInt(0, 2) }
func (t *Tiles) TileWidth() int  { return t.uint16ToInt(2, 4) }
func (t *Tiles) TileHeight() int { return t.uint16ToInt(4, 6) }

func (t *Tiles) uint16ToInt(from, to int) int {
	var v uint16
	_ = binary.Read(bytes.NewReader(t.data[from:to]), binary.LittleEndian, &v)
	return int(v)
}

func (t *Tiles) Image() *image.RGBA {
	data := t.data[6:] // Pixels only, strip off the header.

	//numTiles := t.NumTiles()
	width := t.TileWidth()
	height := t.TileHeight()
	rect := image.Rect(0, 0, width, height)

	pixels := make([]uint8, 0, 4*width*height)

	i3, i4 := 0, 0
	for i3 < 3*width*height {
		// Pixel is represented by 3 bytes pointing at the pallete, representing RGB.
		r, g, b := t.pallete.RGB(data[i3])
		pixels = append(pixels, r, g, b, alpha)
		i3 += 3
		i4 += 4
	}

	return &image.RGBA{pixels, 4 * width, rect}

	//img := &image.RGBA{data[6 : numTiles*width*height], 4 * numTiles * width, numTiles * height}

	// img := image.NewRGBA(rect)
	// for tile := 0; tile < numTiles; tile++ {
	// 	for y := 0; y < height; y++ {
	// 		for x := 0; x < width; x++ {
	// 			img.Set(x, y, color.RGBA{
	// 				t.pallete[t.data[6+tile*numTiles+y*height+x*3]] << 2,   // R
	// 				t.pallete[t.data[6+tile*numTiles+y*height+x*3+1]] << 2, // G
	// 				t.pallete[t.data[6+tile*numTiles+y*height+x*3+2]] << 2, // B
	// 				255, // Alpha
	// 			})
	// 		}
	// 	}
	// }

	//panic(fmt.Sprintf("len(t.data)=%v, numTiles*tileWidth*tileHeight=%v", len(t.data), 6+numTiles*tileWidth*tileHeight))
	// i3, i4 := 0, 0
	// for i3 < numTiles*tileWidth*tileHeight {
	// 	// img.Pix[i4+0] = t.pallete[t.data[6+i3+0]] << 2 // r
	// 	// img.Pix[i4+1] = t.pallete[t.data[6+i3+1]] << 2 // g
	// 	// img.Pix[i4+2] = t.pallete[t.data[6+i3+2]] << 2 // b
	// 	img.Pix[i4+0] = uint8(rand.Int31n(255))
	// 	img.Pix[i4+1] = uint8(rand.Int31n(255))
	// 	img.Pix[i4+2] = uint8(rand.Int31n(255))
	// 	img.Pix[i4+3] = 0xFF // alpha
	// 	i3 += 3
	// 	i4 += 4
	// }
	//panic(fmt.Sprintf("%v:%v ... max %v... %v:%v", numTiles*tileWidth, numTiles*tileHeight, numTiles*tileWidth*tileHeight, i3, i4))
	//return img
}
