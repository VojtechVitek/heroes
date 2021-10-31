package lod

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

// LOD contains Heroes 3 sprites that are zipped in a custom file format.
type Lod struct {
	files map[string]*File
	data  []byte
}

func (l *Lod) String() string {
	var b strings.Builder
	for name, f := range l.files {
		fmt.Fprintf(&b, "%v (offset: 0x%x, fullSize: %v, compressedSize: %v\n", name, f.offset, f.fullSize, f.compressedSize)
	}
	return b.String()
}

func (l *Lod) ReadFile(filename string) ([]byte, error) {
	file, ok := l.files[filename]
	if !ok {
		return nil, errors.Errorf("%q not found in LOD file", filename)
	}

	if file.compressedSize == 0 {
		plainData := l.data[file.offset : file.offset+file.fullSize]
		return plainData, nil
	}

	compressedData := l.data[file.offset : file.offset+file.compressedSize]
	r, err := gzip.NewReader(bytes.NewReader(compressedData)) //, int64(len(compressedData)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read compressed data")
	}

	// for _, f := range r.File {
	// 	fmt.Println(f.Name)
	// }

	return ioutil.ReadAll(r)
	//return nil, nil
}

type File struct {
	filename       string
	offset         int
	fullSize       int
	compressedSize int
}
