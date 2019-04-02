package agg

import (
	"image/png"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestLoadAGGFiles(t *testing.T) {
	t.Parallel()

	for _, file := range []string{
		"./DATA/HEROES2.AGG",
		"./DATA/HEROES2X.AGG",
	} {
		f, err := os.Open(file)
		if err != nil {
			t.Fatal(err)
		}

		agg, err := Load(f)
		if err != nil {
			t.Fatal(errors.Wrapf(err, "failed to load AGG file %v", file))
		}

		t.Logf("%v\n", agg)
	}
}

func TestLoadPalleteAndIcons(t *testing.T) {
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
		pallete := Pallete(data)

		for _, file := range []string{
			"GROUND32.TIL",
		} {
			data, err = agg.Data(file)
			if err != nil {
				t.Fatal(err)
			}

			img := NewTiles(data, pallete).Image()

			out, err := os.Create("./ground32.png")
			if err != nil {
				t.Fatal(err)
			}
			if err := png.Encode(out, img); err != nil {
				t.Fatal(err)
			}
		}
	}
}
