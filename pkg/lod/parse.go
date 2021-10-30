package lod

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
	"github.com/pkg/errors"
)

func Parse(r io.Reader) (*Lod, error) {
	var b bytes.Buffer
	if _, err := b.ReadFrom(r); err != nil {
		return nil, err
	}

	lod := &Lod{
		files: map[string]*File{},
		data:  b.Bytes(),
	}

	get := bytestream.New(b.Bytes(), binary.LittleEndian)

	if get.String(3) != "LOD" {
		return nil, errors.Errorf("unknown file format, not a LOD file")
	}
	_ = get.Bytes(5) // Unknown
	numFiles := get.Int(4)
	_ = get.Bytes(80) // Unknown

	for i := 0; i < numFiles; i++ {
		file := &File{}
		file.filename = get.String(16)
		file.offset = get.Int(4)
		file.fullSize = get.Int(4)
		_ = get.Bytes(4) // Unknown
		file.compressedSize = get.Int(4)

		lod.files[file.filename] = file
	}

	return lod, get.Error()
}
