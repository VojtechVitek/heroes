package lod

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

// LOD contains Heroes 3 sprites that are zipped in a custom file format.
type LOD struct {
	files map[string]*File
	data  []byte
}

func (l *LOD) String() string {
	var b strings.Builder
	for name, f := range l.files {
		fmt.Fprintf(&b, "%v (offset: 0x%x, fullSize: %v, compressedSize: %v\n", name, f.dataOffset, f.fullSize, f.compressedSize)
	}
	return b.String()
}

func (l *LOD) ReadFile(filename string) ([]byte, error) {
	file, ok := l.files[filename]
	if !ok {
		return nil, errors.Errorf("%q not found in LOD file", filename)
	}

	if file.compressedSize == 0 {
		plainData := l.data[file.dataOffset : file.dataOffset+file.fullSize]
		return plainData, nil
	}

	compressedData := l.data[file.dataOffset : file.dataOffset+file.compressedSize]

	r, err := zlib.NewReader(bytes.NewReader(compressedData)) //, int64(len(compressedData)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read compressed data")
	}

	return ioutil.ReadAll(r)
}

type File struct {
	filename       string
	dataOffset     int // Offset within LOD data section.
	fileOffset     int // Offset within LOD file.
	fullSize       int
	compressedSize int // If 0, the data is not compressed.
}
