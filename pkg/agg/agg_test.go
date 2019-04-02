package agg

import (
	"io/ioutil"
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

		r, err := agg.Open("KB.PAL")
		if err != nil {
			t.Fatal(err)
		}

		data, _ := ioutil.ReadAll(r)
		t.Logf("size of KB.PAL: %v", len(data))
		t.Log(data)

		_ = Pallete(data)

	}
}
