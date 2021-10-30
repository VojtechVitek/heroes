package lod

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

// LOD contains Heroes 3 sprites that are zipped in a custom file format.
type Lod struct {
	files map[string]*File
	data  []byte
}

func (l *Lod) String() string {
	return fmt.Sprintf("LOD: %v files", len(l.files))
}

func (l *Lod) ReadFile(filename string) ([]byte, error) {
	file, ok := l.files[filename]
	if !ok {
		return nil, errors.Errorf("can't find %q in LOD file")
	}

	r, err := gzip.NewReader(bytes.NewReader(l.data[file.offset:file.compressedSize]))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(r)
}

type File struct {
	filename       string
	offset         int
	fullSize       int
	compressedSize int
}
