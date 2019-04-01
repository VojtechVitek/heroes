package agg

import (
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestLoadSingleMap(t *testing.T) {
	file := "./DATA/HEROES2.AGG"

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
