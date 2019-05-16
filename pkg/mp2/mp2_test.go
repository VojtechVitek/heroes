package mp2

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/VojtechVitek/heroes/pkg/agg"
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
	var tileWidth, tileHeight int
	var tiles []*image.RGBA

	for _, file := range []string{
		"../agg/DATA/HEROES2.AGG",
	} {
		f, err := os.Open(file)
		if err != nil {
			t.Fatal(err)
		}

		aggFile, err := agg.Load(f)
		if err != nil {
			t.Fatal(errors.Wrapf(err, "failed to load AGG file %v", file))
		}

		data, err := aggFile.Data("KB.PAL")
		if err != nil {
			t.Fatal(err)
		}
		palette, err := agg.NewPalette(data)
		if err != nil {
			t.Fatal(err)
		}

		for _, file := range []string{
			"GROUND32.TIL",
		} {
			data, err = aggFile.Data(file)
			if err != nil {
				t.Fatal(err)
			}

			allTiles := agg.NewTiles(data, palette)
			tiles = allTiles.Images()
			tileWidth, tileHeight = allTiles.TileWidth(), allTiles.TileHeight()
		}
	}

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
		if !strings.HasSuffix(file, "Europe.mp2") {
			continue
		}
		f, err := os.Open(file)
		if err != nil {
			t.Fatal(err)
		}

		m, err := LoadMap(f)
		if err != nil {
			t.Fatal(errors.Wrapf(err, "failed to load map %v", file))
		}

		//t.Logf("%v\n%v", file, m.Header)
		//t.Log("tiles:", m.Tiles)

		mapWidth, mapHeight := m.Width(), m.Height()
		rect := image.Rect(0, 0, mapWidth*tileWidth, mapHeight*tileHeight)
		img := image.NewRGBA(rect)

		for x := 0; x < mapWidth; x++ {
			for y := 0; y < mapHeight; y++ {
				tileIndex := m.Tiles[x*mapWidth+y].TileIndex
				draw.Draw(img, image.Rect(y*tileWidth, x*tileHeight, (y+1)*tileWidth, (x+1)*tileHeight), tiles[tileIndex], image.Point{0, 0}, draw.Src)
			}
		}

		out, err := os.Create(fmt.Sprintf("out/%v.png", filepath.Base(file)))
		if err != nil {
			t.Fatal(err)
		}
		if err := png.Encode(out, img); err != nil {
			t.Fatal(err)
		}
	}
}
