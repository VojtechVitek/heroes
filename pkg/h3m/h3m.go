package h3m

// The binary format sourced from
// https://github.com/potmdehex/homm3tools/blob/master/h3m/h3mlib/h3m_structures/h3m.h
// The MIT License (MIT)
// Copyright (c) 2016 John Åkerblom

// Maps commonly end with 124 bytes of null padding. Extra content at end
// is ok.
type H3M struct {
	Format  FileFormat
	MapInfo MapInfo
	Players [8]Player
}
