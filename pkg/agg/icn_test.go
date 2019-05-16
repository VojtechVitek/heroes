package agg

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestLoadICNs(t *testing.T) {
	t.Parallel()

	for _, file := range []string{
		"./DATA/HEROES2.AGG",
	} {
		f, err := os.Open(file)
		if err != nil {
			t.Fatal(err)
		}

		agg, err := Load(f)
		if err != nil {
			t.Fatal(errors.Wrapf(err, "failed to load AGG file %v", file))
		}

		data, err := agg.Data("KB.PAL")
		if err != nil {
			t.Fatal(err)
		}
		palette, err := NewPalette(data)
		if err != nil {
			t.Fatal(err)
		}

		gifPalette := palette.GifPalette()

		for _, file := range agg.Files("ICN") {
			data, err = agg.Data(file)
			if err != nil {
				t.Fatal(err)
			}

			icn, err := NewICN(data, palette)
			if err != nil {
				t.Fatal(err)
			}

			outGif := &gif.GIF{}
			sprites := icn.Sprites()

			var maxW, maxH, negX, negY int
			for _, sprite := range sprites {
				if sprite.Width > maxW {
					maxW = sprite.Width
				}
				if sprite.Width > maxH {
					maxH = sprite.Width
				}
				if sprite.X < negX {
					negX = sprite.X
				}
				if sprite.Y < negY {
					negY = sprite.Y
				}
			}

			for _, sprite := range sprites {
				if sprite.Width <= 1 && sprite.Height <= 1 {
					continue
				}

				img, err := sprite.RenderImage(palette)
				if err != nil {
					t.Fatal(errors.Wrap(err, "failed to render image"))
				}

				palettedImg := image.NewPaletted(image.Rect(0, 0, maxW-negX, maxH-negY), gifPalette)
				draw.FloydSteinberg.Draw(palettedImg, image.Rect(sprite.X-negX, sprite.Y-negY, sprite.X-negX+sprite.Width, sprite.Y-negY+sprite.Height), img, image.ZP)

				outGif.Image = append(outGif.Image, palettedImg)
				outGif.Delay = append(outGif.Delay, 15)
			}

			if len(outGif.Image) <= 1 {
				continue
			}

			out, err := os.Create(fmt.Sprintf("out/%v.gif", file))
			if err != nil {
				t.Fatal(err)
			}

			if err := gif.EncodeAll(out, outGif); err != nil {
				t.Fatal(err)
			}
		}
	}
}
