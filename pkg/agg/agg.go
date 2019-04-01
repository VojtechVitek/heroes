package agg

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type AGG struct {
	*Header // 0x00

	Files []File // 0x02

	Filenames []Filename // EOF - 15 * length(Files)
}

func Load(r io.ReadSeeker) (*AGG, error) {
	header := &Header{}
	if err := binary.Read(r, binary.LittleEndian, header); err != nil {
		return nil, errors.Wrap(err, "failed to parse AGG header")
	}

	files := make([]File, header.NumFiles())
	if err := binary.Read(r, binary.LittleEndian, files); err != nil {
		return nil, errors.Wrap(err, "failed to parse files")
	}

	r.Seek(-int64(header.NumFiles()*FilenameLength), io.SeekEnd)
	filenames := make([]Filename, header.NumFiles())
	if err := binary.Read(r, binary.LittleEndian, filenames); err != nil {
		return nil, errors.Wrap(err, "failed to parse file names")
	}

	agg := &AGG{
		Header:    header,
		Files:     files,
		Filenames: filenames,
	}

	return agg, nil
}

func (agg *AGG) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Number of files: %v\n", agg.Header.NumFiles())

	for i := 0; i < len(agg.Files); i++ {
		fmt.Fprintf(&b, "%v: %v", agg.Filenames[i], agg.Files[i])
	}

	return b.String()
}
