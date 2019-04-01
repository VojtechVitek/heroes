package agg

import (
	"bytes"
	"encoding/binary"
)

type Header [2]byte

func (h *Header) NumFiles() int {
	var numFiles uint16
	_ = binary.Read(bytes.NewReader(h[0:2]), binary.LittleEndian, &numFiles)
	return int(numFiles)
}
