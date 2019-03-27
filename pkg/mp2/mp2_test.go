package mp2

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestLoadMap(t *testing.T) {
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

		m, err := LoadMap(f)
		if err != nil {
			t.Error(errors.Wrapf(err, "failed to parse %v", file))
			continue
		}

		t.Log(path.Base(file), "width:", m.Width(), "height:", m.Height()) //, "mapWidth:", m.MapWidth, "mapHeight:", m.MapHeight, "\nmapTiles:", m.MapTiles)
	}
}
