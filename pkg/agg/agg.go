package agg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

// 0x00   (0)    Length             2 bytes
type AGG [4]byte

func (agg *AGG) Length() int {
	var length uint16
	_ = binary.Read(bytes.NewReader(agg[0:2]), binary.LittleEndian, &length)
	return int(length)
}

func Load(r io.Reader) (*AGG, error) {
	agg := &AGG{}

	if err := binary.Read(r, binary.LittleEndian, agg); err != nil {
		return nil, errors.Wrap(err, "failed to parse AGG header")
	}

	// tiles := make(Tiles, header.Width()*header.Height())
	// if err := binary.Read(r, binary.LittleEndian, tiles); err != nil {
	// 	return nil, errors.Wrap(err, "failed to parse tiles")
	// }

	return agg, nil
}

func (agg *AGG) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Length: %v\n", agg.Length())

	return b.String()
}
