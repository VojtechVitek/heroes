package mp2

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParsingMp2Maps(t *testing.T) {
	var mp2Files []string

	dir, _ := os.Getwd()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".mp2") {
			mp2Files = append(mp2Files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range mp2Files {
		f, err := os.Open(file)
		if err != nil {
			t.Error(err)
		}

		m, err := LoadMap(f)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%+v", m)
	}
}
