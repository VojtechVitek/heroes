package count

import "io"

type reader struct {
	r         io.Reader
	BytesRead int
}

func NewReader(r io.Reader) *reader {
	return &reader{
		r: r,
	}
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	r.BytesRead += n
	return n, err
}
