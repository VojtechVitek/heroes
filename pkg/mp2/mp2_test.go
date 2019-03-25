package mp2

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestParsingMp2Maps(t *testing.T) {
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
			t.Error(err)
		}

		m, err := LoadMap(f)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v:\n%+v", path.Base(file), m)
	}
}
