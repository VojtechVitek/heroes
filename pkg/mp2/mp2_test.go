package mp2

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestLoadMapsHeader(t *testing.T) {
	var mapFiles []string

	dir, _ := os.Getwd()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		filenameLower := strings.ToLower(filepath.Base(path))
		if strings.HasSuffix(filenameLower, ".mp2") || strings.HasSuffix(filenameLower, ".mx2") {
			mapFiles = append(mapFiles, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range mapFiles {
		f, err := os.Open(file)
		if err != nil {
			t.Fatal(err)
		}

		h, err := LoadHeader(f)
		if err != nil {
			t.Fatal(errors.Wrapf(err, "failed to load map %v", file))
		}

		t.Logf("%v\n%v", path.Base(file), h)
	}
}

func TestLoadSingleMap(t *testing.T) {
	//file := "./maps/THEOTHER.MP2"
	//file := "./maps/PANDAMON.MP2"
	file := "./maps/SLUGFEST.MP2"

	f, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}

	m, err := LoadMap(f)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "failed to load map %v", file))
	}

	//t.Logf("%v\n%v", file, m.Header)
	t.Log("tiles:", m.Tiles)
}
