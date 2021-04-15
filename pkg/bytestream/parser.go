package bytestream

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
)

type Parser struct {
	b     *bytes.Buffer
	order binary.ByteOrder
	err   error
}

func New(b []byte, order binary.ByteOrder) *Parser {
	return &Parser{
		b:     bytes.NewBuffer(b),
		order: order,
	}
}

func (p *Parser) Error() error {
	return p.err
}

func (p *Parser) Int(numBytes int) int {
	if p.err != nil {
		return 0
	}

	switch numBytes {
	case 1:
		var v uint8
		p.err = binary.Read(p.b, p.order, &v)
		return int(v)
	case 2:
		var v uint16
		p.err = binary.Read(p.b, p.order, &v)
		return int(v)
	case 4:
		var v uint32
		p.err = binary.Read(p.b, p.order, &v)
		return int(v)
	case 8:
		var v uint64
		p.err = binary.Read(p.b, p.order, &v)
		return int(v)
	default:
		panic(fmt.Sprintf("GetInt(%v) is not implemented", numBytes))
	}
}

func (p *Parser) Bool(numBytes int) bool {
	if p.err != nil {
		return false
	}

	switch numBytes {
	case 1:
		var v uint8
		p.err = binary.Read(p.b, p.order, &v)
		return v > 0
	case 2:
		var v uint16
		p.err = binary.Read(p.b, p.order, &v)
		return v > 0
	case 4:
		var v uint32
		p.err = binary.Read(p.b, p.order, &v)
		return v > 0
	case 8:
		var v uint64
		p.err = binary.Read(p.b, p.order, &v)
		return v > 0
	default:
		panic(fmt.Sprintf("GetInt(%v) is not implemented", numBytes))
	}
}

func (p *Parser) ReadBytes(numBytes int) []byte {
	if p.err != nil || numBytes == 0 {
		return nil
	}

	buf := make([]byte, numBytes)
	n, err := p.b.Read(buf)
	if err != nil {
		p.err = errors.Wrapf(err, "ReadBytes(%v)", numBytes)
		return nil
	}
	if n != numBytes {
		p.err = errors.Errorf("ReadBytes(%v): failed to read all bytes - only %v bytes were read", numBytes, n)
		return nil
	}

	return buf
}

func (p *Parser) ReadString(numBytes int) string {
	return string(p.ReadBytes(numBytes))
}

func (p *Parser) ReadCString() string {
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
