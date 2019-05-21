package mp2

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/VojtechVitek/heroes/pkg/agg"
	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
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

func TestRenderMap(t *testing.T) {
	var tileWidth, tileHeight int
	var tiles []image.Image
	var windmill image.Image

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

		{
			data, err = aggFile.Data("GROUND32.TIL")
			if err != nil {
				t.Fatal(err)
			}

			allTiles := agg.NewTiles(data)
			tiles = allTiles.Images(palette)
			tileWidth, tileHeight = allTiles.TileWidth(), allTiles.TileHeight()
		}

		{
			data, err := aggFile.Data("OBJNGRA2.ICN")
			if err != nil {
				t.Fatal(err)
			}

			icn, err := agg.NewICN(data)
			if err != nil {
				t.Fatal(err)
			}

			windmill, err = icn.Sprites()[39+16].RenderImage(palette)
			if err != nil {
				t.Fatal(err)
			}
			_ = windmill
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

		mapWidth, mapHeight := m.Width(), m.Height()
		rect := image.Rect(0, 0, mapWidth*tileWidth, mapHeight*tileHeight)
		mapImg := image.NewRGBA(rect)

		// var b strings.Builder
		for y := 0; y < mapHeight; y++ {
			for x := 0; x < mapWidth; x++ {
				drawRect := image.Rect(x*tileWidth, y*tileHeight, (x+1)*tileWidth, (y+1)*tileHeight)
				tile := m.Tiles[y*mapWidth+x]

				img := tiles[tile.TileIndex]

				switch tile.Shape % 4 {
				case 1: // vertical flip
					img = imaging.FlipV(img)
				case 2: // horizontal flip
					img = imaging.FlipH(img)
				case 3: // vertical+horizontal flip
					img = imaging.FlipV(img)
					img = imaging.FlipH(img)
				}
				draw.Draw(mapImg, drawRect, img, image.Point{0, 0}, draw.Src)
			}
		}

		for y := 0; y < mapHeight; y++ {
			for x := 0; x < mapWidth; x++ {
				tile := m.Tiles[y*mapWidth+x]
				_ = tile
				// switch obj {
				// case Windmill:
				addLabel(mapImg, x*tileWidth, y*tileHeight+tileHeight, fmt.Sprintf("%v", tile.UniqueNumber1))
				//				}
			}
			// fmt.Fprintln(&b)
		}

		// t.Log(b.String())

		out, err := os.Create(fmt.Sprintf("out/%v.png", filepath.Base(file)))
		if err != nil {
			t.Fatal(err)
		}
		if err := png.Encode(out, mapImg); err != nil {
			t.Fatal(err)
		}
	}
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
