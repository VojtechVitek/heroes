package agg

import (
	"image/png"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestLoadMapTiles(t *testing.T) {
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
		pallete, err := NewPallete(data)
		if err != nil {
			t.Fatal(err)
		}

		for _, file := range []string{
			"GROUND32.TIL",
			"CLOF32.TIL",
			"STON.TIL",
		} {
			data, err = agg.Data(file)
			if err != nil {
				t.Fatal(err)
			}

			img := NewTiles(data, pallete).Image()

			out, err := os.Create(file + ".png")
			if err != nil {
				t.Fatal(err)
			}
			if err := png.Encode(out, img); err != nil {
				t.Fatal(err)
			}
		}
	}
}
