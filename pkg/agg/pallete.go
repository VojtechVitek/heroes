package agg

import (
	"github.com/pkg/errors"
)

type pallete []byte

func NewPallete(data []byte) (pallete, error) {
	if len(data) != 3*256 {
		return nil, errors.Errorf("failed to create color pallete: expected %v bytes, got %v bytes", 3*256, len(data))
	}
	return pallete(data), nil
}

func (p pallete) RGB(index uint8) (r, g, b uint8) {
	return p[index*3] << 2, p[index*3+1] << 2, p[index*3+2] << 2
}
