package agg

import (
	"fmt"
	"image/png"
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
		pallete, err := NewPallete(data)
		if err != nil {
			t.Fatal(err)
		}

		for _, file := range agg.Files("ICN") {
			data, err = agg.Data(file)
			if err != nil {
				t.Fatal(err)
			}

			icn, err := NewICN(data, pallete)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("%+v", icn)

			for i, sprite := range icn.Sprites() {
				img, err := sprite.RenderImage(pallete)
				if err != nil {
					t.Fatal(errors.Wrap(err, "failed to render image"))
				}
				out, err := os.Create(fmt.Sprintf("out/%v-%v.png", file, i))
				if err != nil {
					t.Fatal(err)
				}
				if err := png.Encode(out, img); err != nil {
					t.Fatal(errors.Wrap(err, "failed to render into PNG"))
				}
			}
		}
	}
}
