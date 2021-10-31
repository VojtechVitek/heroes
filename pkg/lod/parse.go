package lod

import (
	"encoding/binary"
	"io"
	"io/ioutil"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
	"github.com/pkg/errors"
)

func Parse(r io.Reader) (*Lod, error) {
	lod := &Lod{
		files: map[string]*File{},
	}

	countingReader := &countingReader{r: r}

	get := bytestream.New(countingReader, binary.LittleEndian)

	if get.String(3) != "LOD" {
		return nil, errors.Errorf("failed to parse LOD file header")
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

	if err := get.Error(); err != nil {
		return nil, errors.Wrap(err, "failed to parse LOD file header data")
	}

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read LOD data")
	}

	// Data starts exactly at 0x4E25C position. But we already read some data
	// from the original io.Reader above this, so we need to "seek" correctly.
	lod.data = buf[0x4E25C-countingReader.bytesRead:]

	return lod, nil
}

type countingReader struct {
	r         io.Reader
	bytesRead int
}

func (r *countingReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	r.bytesRead += n
	return n, err
}
