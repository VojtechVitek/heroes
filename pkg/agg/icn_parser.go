package agg

type spriteParserState int

const (
	start spriteParserState = iota
	eol
	eof
)

type spriteParser struct {
	data      []byte
	lineWidth int
	pallete   pallete
	state     spriteParserState
}

func (p *spriteParser) next() bool {
	switch p.state {
	case start:
		p.state = eof
	case eof:
		return false
	}
	return true
}

func (p *spriteParser) getPixels() []uint8 {
	return []uint8{0, 0, 0, 0}
}
