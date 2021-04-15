package h3m

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
)

// Ported from C/C++ https://github.com/potmdehex/homm3toolsby.
// Original file: https://github.com/potmdehex/homm3tools/blob/master/h3m/h3mlib/h3m_structures/h3m.h
// The MIT License (MIT)
// Copyright (c) 2016 John Åkerblom

// Maps commonly end with 124 bytes of null padding. Extra content at end
// is ok.
type H3M struct {
	Format int
	Roe1   Roe
}

type Roe struct {
	HasHero      bool
	MapSize      int
	HasTwoLevels bool
	NameSize     int
	Name         string
	DescSize     int
	Desc         string
}

func Parse(r io.Reader) (*H3M, error) {
	r, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	_, err = b.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	get := bytestream.New(b.Bytes())

	h3m := &H3M{}
	h3m.Format = get.Int(4)

	h3m.Roe1.HasHero = get.Bool(1)
	h3m.Roe1.MapSize = get.Int(4)
	h3m.Roe1.HasTwoLevels = get.Bool(1)
	h3m.Roe1.NameSize = get.Int(4)
	h3m.Roe1.Name = get.ReadCString()
	h3m.Roe1.DescSize = get.Int(4)
	h3m.Roe1.Desc = get.ReadCString()

	return h3m, get.Error()
}
