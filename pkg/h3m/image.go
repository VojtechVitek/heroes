package h3m

import (
	"image"

	"github.com/VojtechVitek/heroes/pkg/palette"
)

func (h3m *H3M) Image() (image.Image, error) {
	width, height := h3m.MapInfo.MapSize.Size()
	pixels := make([]uint8, 0, len(h3m.Tiles)*4)

	for _, tile := range h3m.Tiles {
		r, g, b := tile.TerrainType.RGB()
		pixels = append(pixels, r, g, b, palette.OpaqueAlpha)
	}

	// // Fill in blank pixels.
	// if cap(pixels)-len(pixels) > 0 {
	// 	return nil, errors.Errorf("failed to parse PCX image (format %v): missing %v pixels", cap(pixels)-len(pixels))
	// }
	// for i := 0; i < cap(pixels)-len(pixels); i++ {
	// 	pixels = append(pixels, 255, 255, 255, palette.OpaqueAlpha) // White.
	// }

	rect := image.Rect(0, 0, width, height)
	img := &image.RGBA{pixels, 4 * width, rect}

	return img, nil
}
