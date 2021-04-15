package h3m

import "fmt"

type FileFormat int

const (
	Orig             FileFormat = 0x0000000E // The Restoration of Erathia (the original HOMM III game)
	ArmageddonsBlade FileFormat = 0x00000015 // Armageddon's Blade (HOMM III expansion pack)
	ShadowOfDeath    FileFormat = 0x0000001C // The Shadow of Death (HOMM III expansion pack)
	_CHR             FileFormat = 0x0000001D // Chronicles?
	_WOG             FileFormat = 0x00000033 // In the Wake of Gods (free fan-made expansion pack)
)

func (f FileFormat) String() string {
	switch f {
	case Orig:
		return "The Restoration of Erathia (original map)"
	case ArmageddonsBlade:
		return "Armageddon's Blade"
	case ShadowOfDeath:
		return "The Shadow of Death"
	default:
		return fmt.Sprintf("unknown (%x)", f)
	}
}

func (f FileFormat) Is(formats ...FileFormat) bool {
	for _, format := range formats {
		if f == format {
			return true
		}
	}
	return false
}
