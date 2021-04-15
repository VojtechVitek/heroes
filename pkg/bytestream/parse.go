package bytestream

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type parser struct {
	b   *bytes.Buffer
	err error
}

// if err := binary.Read(bytes.NewReader(data[0:2]), binary.LittleEndian, &u16); err != nil {
//		return nil, errors.Wrap(err, "failed to read number of sprites")
// }

func New(b []byte) *parser {
	return &parser{
		b: bytes.NewBuffer(b),
	}
}

func (p *parser) Error() error {
	return p.err
}

func (p *parser) Int(numBytes int) int {
	if p.err != nil {
		return 0
	}

	switch numBytes {
	case 1:
		var v uint8
		p.err = binary.Read(p.b, binary.LittleEndian, &v)
		return int(v)
	case 2:
		var v uint16
		p.err = binary.Read(p.b, binary.LittleEndian, &v)
		return int(v)
	case 4:
		var v uint32
		p.err = binary.Read(p.b, binary.LittleEndian, &v)
		return int(v)
	case 8:
		var v uint64
		p.err = binary.Read(p.b, binary.LittleEndian, &v)
		return int(v)
	default:
		panic(fmt.Sprintf("GetInt(%v) is not implemented", numBytes))
	}
}

func (p *parser) Bool(numBytes int) bool {
	if p.err != nil {
		return false
	}

	switch numBytes {
	case 1:
		var v uint8
		p.err = binary.Read(p.b, binary.LittleEndian, &v)
		return v > 0
	case 2:
		var v uint16
		p.err = binary.Read(p.b, binary.LittleEndian, &v)
		return v > 0
	case 4:
		var v uint32
		p.err = binary.Read(p.b, binary.LittleEndian, &v)
		return v > 0
	case 8:
		var v uint64
		p.err = binary.Read(p.b, binary.LittleEndian, &v)
		return v > 0
	default:
		panic(fmt.Sprintf("GetInt(%v) is not implemented", numBytes))
	}
}

func (p *parser) ReadCString() string {
	if p.err != nil {
		return ""
	}

	str, err := p.b.ReadString(byte(0))
	if err != nil {
		p.err = err
		return ""
	}

	if len(str) <= 1 {
		return ""
	}

	return str[:len(str)-2]
}
