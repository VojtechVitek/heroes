package h3m

import "fmt"

type FileFormat int

const (
	ROE              FileFormat = 0x0000000E // The Restoration of Erathia (the original HOMM III game)
	ArmageddonsBlade FileFormat = 0x00000015 // Armageddon's Blade (HOMM III expansion pack)
	ShadowOfDeath    FileFormat = 0x0000001C // The Shadow of Death (HOMM III expansion pack)
	_CHR             FileFormat = 0x0000001D // Chronicles?
	HOTA             FileFormat = 0x00000020 // Horn of the Abyss (HotA mod)
	_WOG             FileFormat = 0x00000033 // In the Wake of Gods (free fan-made expansion pack)
)

func (f FileFormat) String() string {
	switch f {
	case ROE:
		return "The Restoration of Erathia (original map)"
	case ArmageddonsBlade:
		return "Armageddon's Blade"
	case ShadowOfDeath:
		return "The Shadow of Death"
	case HOTA:
		return "Horn of the Abyss"
	default:
		return fmt.Sprintf("unknown H3M file format (%x)", int(f))
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
