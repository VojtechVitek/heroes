package agg

import (
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestLoadPallete(t *testing.T) {
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
		pallete, err := NewPalette(data)
		if err != nil {
			t.Fatal(err)
		}

		pixels := make([]uint8, 0, 4*256)
		for i := 0; i <= 255; i++ {
			r, g, b := pallete.RGB(i)
			pixels = append(pixels, r, g, b, 255)
		}

		img := &image.RGBA{pixels, 4 * 16, image.Rect(0, 0, 16, 16)}
		out, err := os.Create("pallete.png")
		if err != nil {
			t.Fatal(err)
		}
		if err := png.Encode(out, img); err != nil {
			t.Fatal(err)
		}
	}
}
