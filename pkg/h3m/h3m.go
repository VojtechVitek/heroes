package h3m

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
)

// The binary format sourced from
// https://github.com/potmdehex/homm3tools/blob/master/h3m/h3mlib/h3m_structures/h3m.h
// The MIT License (MIT)
// Copyright (c) 2016 John Åkerblom

// Maps commonly end with 124 bytes of null padding. Extra content at end
// is ok.
type H3M struct {
	Format  FileFormat
	MapInfo MapInfo
}

type FileFormat int

const (
	Orig             FileFormat = 0x0000000E // The Restoration of Erathia (the original HOMM III game)
	ArmageddonsBlade FileFormat = 0x00000015 // Armageddon's Blade (HOMM III expansion pack)
	ShadowOfDeath    FileFormat = 0x0000001C // The Shadow of Death (HOMM III expansion pack)
	_CHR             FileFormat = 0x0000001D // Chronicles?
	_WOG             FileFormat = 0x00000033 // In the Wake of Gods (free fan-made expansion pack)
)

type MapInfo struct {
	HasHero      bool
	MapSize      int
	HasTwoLevels bool
	Name         string
	Desc         string
	Difficulty   int
	MasteryCap   int // Only set when format is ArmageddonsBlade or ShadowOfDeath.
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

	get := bytestream.New(b.Bytes(), binary.LittleEndian)

	h3m := &H3M{}
	h3m.Format = FileFormat(get.Int(4))
	h3m.MapInfo.HasHero = get.Bool(1)
	h3m.MapInfo.MapSize = get.Int(4)
	h3m.MapInfo.HasTwoLevels = get.Bool(1)
	nameSize := get.Int(4)
	h3m.MapInfo.Name = get.ReadString(nameSize)
	descSize := get.Int(4)
	h3m.MapInfo.Desc = get.ReadString(descSize)
	h3m.MapInfo.Difficulty = get.Int(1)

	switch h3m.Format {
	case ArmageddonsBlade, ShadowOfDeath:
		h3m.MapInfo.MasteryCap = get.Int(1)
	}

	return h3m, get.Error()
}
