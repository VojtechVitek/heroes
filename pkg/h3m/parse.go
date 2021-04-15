package h3m

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
)

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

	h3m := &H3M{
		Format: FileFormat(get.Int(4)),
	}

	h3m.MapInfo.HasHero = get.Bool(1)
	h3m.MapInfo.MapSize = get.Int(4)
	h3m.MapInfo.HasTwoLevels = get.Bool(1)
	nameSize := get.Int(4)
	h3m.MapInfo.Name = get.String(nameSize)
	descSize := get.Int(4)
	h3m.MapInfo.Desc = get.String(descSize)
	h3m.MapInfo.Difficulty = get.Int(1)

	if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
		h3m.MapInfo.MasteryCap = get.Int(1)
	}

	// Parse players.
	for i := 0; i < len(h3m.Players); i++ {
		player := &h3m.Players[i]
		player.CanBeHuman = get.Bool(1)
		player.CanBeHuman = get.Bool(1)
		player.Behavior = get.Int(1)

		if h3m.Format == ShadowOfDeath {
			player.AllowedAlignments = get.Int(1)
		}

		player.TownTypes = get.Int(1)

		if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
			townConflux := get.Int(1)
			_ = townConflux // not used for now..
		}

		player.HasRandomTown = get.Bool(1)
		player.HasMainTown = get.Bool(1)
	}

	if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
		// Parse AI player heroes.

		_ = get.Bytes(1) // unknown byte
		heroesCount := get.Int(4)
		fmt.Printf("Heroes count: %v\n", heroesCount)
		for i := 0; i < heroesCount; i++ {
			playerType := get.Int(1)
			nameSize := get.Int(4)
			name := get.String(nameSize)
			fmt.Printf("AI player heroes: %v (type %v)\n", name, playerType)
		}
	}

	// Parse additional map information.
	h3m.MapInfo.WinCondition = Condition(get.Int(1))
	h3m.MapInfo.WinConditionAllowNormalWin = get.Bool(1) // Allow normal win. Supposedly doesn't work.
	h3m.MapInfo.WinConditionAppliesToComputer = get.Bool(1)

	if h3m.MapInfo.WinCondition.Is(ACQUIRE_ARTIFACT, ACCUMULATE_CREATURES, BUILD_GRAIL, DEFEAT_HERO) {
		if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
			h3m.MapInfo.WinConditionType = get.Int(2)
		} else {
			h3m.MapInfo.WinConditionType = get.Int(1)
		}
	}

	if h3m.MapInfo.WinCondition.Is(ACCUMULATE_CREATURES, ACCUMULATE_RESOURCES) {
		h3m.MapInfo.WinConditionAmount = get.Int(4)
	}

	if h3m.MapInfo.WinCondition.Is(TRANSPORT_ARTIFACT, UPGRADE_TOWN) {
		h3m.MapInfo.WinConditionPos.X = get.Int(1)
		h3m.MapInfo.WinConditionPos.Y = get.Int(1)
		h3m.MapInfo.WinConditionPos.Z = get.Int(1)
	}

	if h3m.MapInfo.WinCondition.Is(UPGRADE_TOWN) {
		h3m.MapInfo.WinConditionUpgradeHallLevel = get.Int(1)
		h3m.MapInfo.WinConditionUpgradeCastleLevel = get.Int(1)
	}

	//h3m.MapInfo.LoseCondition = Condition(get.Int(1))

	return h3m, get.Error()
}
