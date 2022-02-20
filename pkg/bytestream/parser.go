package bytestream

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

type Parser struct {
	b     *bufio.Reader
	order binary.ByteOrder
	err   error
}

func New(r io.Reader, order binary.ByteOrder) *Parser {
	return &Parser{
		b:     bufio.NewReaderSize(r, 100),
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

func (p *Parser) Byte() byte {
	if p.err != nil {
		return 0
	}

	var v uint8
	p.err = binary.Read(p.b, p.order, &v)
	return byte(v)
}

func (p *Parser) Bytes(numBytes int) []byte {
	if p.err != nil || numBytes == 0 {
		return nil
	}

	if numBytes == -1 {
		b, _ := ioutil.ReadAll(p.b)
		return b
	}

	buf := make([]byte, numBytes)
	_, err := io.ReadFull(p.b, buf)
	if err != nil {
		p.err = errors.Wrapf(err, "ReadBytes(%v)", numBytes)
		return nil
	}

	return buf
}

func (p *Parser) String(numBytes int) string {
	cString := p.Bytes(numBytes)
	if i := bytes.IndexByte(cString, 0); i > 0 {
		return string(cString[:i])
	}

	return string(cString)
}

func (p *Parser) CString() string {
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
