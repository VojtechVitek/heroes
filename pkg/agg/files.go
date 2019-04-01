package agg

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/VojtechVitek/heroes/pkg/cstring"
)

// 0x00    CRC      4 bytes
// 0x04    Offset   4 bytes // TODO: This is 2 bytes only in Heroes 1.
// 0x08    Size     4 bytes
type File [12]byte

func (f File) CRC() int {
	var crc uint32
	_ = binary.Read(bytes.NewReader(f[0:4]), binary.LittleEndian, &crc)
	return int(crc)
}
func (f File) Offset() int {
	var offset uint32
	_ = binary.Read(bytes.NewReader(f[4:8]), binary.LittleEndian, &offset)
	return int(offset)
}
func (f File) Size() int {
	var size uint32
	_ = binary.Read(bytes.NewReader(f[8:12]), binary.LittleEndian, &size)
	return int(size)
}

func (f File) String() string {
	return fmt.Sprintf("CRC: %v, Offset: %v, Size: %v\n", f.CRC(), f.Offset(), f.Size())
}

const FilenameLength = 15

type Filename [FilenameLength]byte

func (f Filename) String() string {
	return cstring.String(f[:])
}
