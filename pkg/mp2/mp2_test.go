package mp2

import (
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
	//t.Log("tiles:", m.Tiles)

	width, height := m.Width(), m.Height()
	rect := image.Rect(0, 0, width*tileWidth, height*tileHeight)
	img := image.NewRGBA(rect)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			tileIndex := m.Tiles[x*width+y].TileIndex
			t.Logf("drawing tile %4v", tileIndex)
			draw.Draw(img, image.Rect(x*width, y*height, (x+1)*width, (y+1)*height), tiles[tileIndex], image.Point{0, 0}, draw.Src)
		}
		t.Log()
	}

	out, err := os.Create("map.png")
	if err != nil {
		t.Fatal(err)
	}
	if err := png.Encode(out, img); err != nil {
		t.Fatal(err)
	}
}
