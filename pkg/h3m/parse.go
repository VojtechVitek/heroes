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

	h3m := &H3M{}

	// https://github.com/potmdehex/homm3tools/blob/5687f581a4eb5e7b0e8f48794d7be4e3b0a8cc8b/h3m/h3mlib/h3m_structures/h3m.h#L28
	h3m.Format = FileFormat(get.Int(4))

	// Basic map info, aka H3M_BI.
	// https://github.com/potmdehex/homm3tools/blob/5687f581a4eb5e7b0e8f48794d7be4e3b0a8cc8b/h3m/h3mlib/h3m_structures/h3m.h#L29
	h3m.MapInfo.HasHero = get.Bool(1)
	h3m.MapInfo.MapSize = get.Int(4)
	h3m.MapInfo.HasTwoLevels = get.Bool(1)
	nameLen := get.Int(4)
	h3m.MapInfo.Name = get.String(nameLen)
	descLen := get.Int(4)
	h3m.MapInfo.Desc = get.String(descLen)
	h3m.MapInfo.Difficulty = get.Int(1)

	if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
		h3m.MapInfo.MasteryCap = get.Int(1)
	}

	// Players, aka [8]H3M_PLAYER.
	// https://github.com/potmdehex/homm3tools/blob/5687f581a4eb5e7b0e8f48794d7be4e3b0a8cc8b/h3m/h3mlib/h3m_structures/h3m.h#L30
	for i := 0; i < len(h3m.Players); i++ {
		player := &h3m.Players[i]
		player.CanBeHuman = get.Bool(1)
		player.CanBeComputer = get.Bool(1)
		player.Behavior = get.Int(1)

		if h3m.Format.Is(ShadowOfDeath) {
			player.AllowedAlignments = get.Int(1)
		}

		player.TownTypes = get.Int(1)

		if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
			townConflux := get.Int(1)
			_ = townConflux // not used for now..
		}

		player.Unknown1_HasRandomTown = get.Bool(1)
		player.HasMainTown = get.Bool(1)

		if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
			// union H3M_PLAYER_EXT_ABSOD {
			//   struct H3M_PLAYER_EXT_ABSOD_DEFAULT e0;
			//   struct H3M_PLAYER_EXT_WITH_TOWN_ABSOD e1;
			//   struct H3M_PLAYER_EXT_WITH_HERO_ABSOD e2;
			//   struct H3M_PLAYER_EXT_WITH_TOWN_AND_HERO_ABSOD e3;
			// }

			if false { // H3M_PLAYER_EXT_ABSOD_DEFAULT
				player.StartingHeroIsRandom = get.Bool(1)
				player.StartingHeroType = get.Int(1)
				player.StartingHeroFace = get.Int(1)

				heroNameLen := get.Int(4)
				player.StartingHeroName = get.String(heroNameLen)
				fmt.Printf("HERO NAME: %v", player.StartingHeroName)
			}

			if false { // H3M_PLAYER_EXT_WITH_TOWN_ABSOD
				player.StartingTownCreateHero = get.Bool(1)
				player.StartingTownType = get.Int(1)
				player.StartingTownPos.X = get.Int(1)
				player.StartingTownPos.Y = get.Int(1)
				player.StartingTownPos.Z = get.Int(1)

				player.StartingHeroIsRandom = get.Bool(1)
				player.StartingHeroType = get.Int(1)
				player.StartingHeroFace = get.Int(1)

				startingHeroNameLen := get.Int(4)
				player.StartingHeroName = get.String(startingHeroNameLen)
				fmt.Printf("HERO NAME: %v", player.StartingHeroName)
			}

			if false { // H3M_PLAYER_EXT_WITH_HERO_ABSOD
				player.StartingHeroIsRandom = get.Bool(1)
				player.StartingHeroType = get.Int(1)
				player.StartingHeroFace = get.Int(1)

				startingHeroNameLen := get.Int(4)
				player.StartingHeroName = get.String(startingHeroNameLen)
				fmt.Printf("HERO NAME: %v", player.StartingHeroName)
			}

			if true { // H3M_PLAYER_EXT_WITH_TOWN_AND_HERO_ABSOD
				player.StartingTownCreateHero = get.Bool(1)
				player.StartingTownType = get.Int(1)
				player.StartingTownPos.X = get.Int(1)
				player.StartingTownPos.Y = get.Int(1)
				player.StartingTownPos.Z = get.Int(1)

				player.StartingHeroIsRandom = get.Bool(1)
				player.StartingHeroType = get.Int(1)
				player.StartingHeroFace = get.Int(1)

				startingHeroNameLen := get.Int(4)
				player.StartingHeroName = get.String(startingHeroNameLen)
				fmt.Printf("HERO NAME: %v", player.StartingHeroName)
			}

		} else {

			panic("ROE not implemented")

			// H3M_PLAYER_EXT_ROE
			startingHeroNameLen := get.Int(4)
			player.StartingHeroName = get.String(startingHeroNameLen)
			fmt.Printf("HERO NAME: %v\n", player.StartingHeroName)

			// struct H3M_PLAYER_EXT_ROE_DEFAULT e0;
			player.StartingHeroIsRandom = get.Bool(1)
			player.StartingHeroType = get.Int(1)

			// struct H3M_PLAYER_EXT_WITH_TOWN_ROE e1;
			if true /*player.HasMainTown && player.StartingHeroType == 0xFF*/ {
				player.StartingTownPos.X = get.Int(1)
				player.StartingTownPos.Y = get.Int(1)
				player.StartingTownPos.Z = get.Int(1)

				player.StartingHeroIsRandom = get.Bool(1)
				player.StartingHeroType = get.Int(1)
			}

			// struct H3M_PLAYER_EXT_WITH_HERO_ROE e2;
			if true {
				player.StartingHeroIsRandom = get.Bool(1)
				player.StartingHeroType = get.Int(1)

				player.StartingHeroFace = get.Int(1)
				startingHeroNameLen := get.Int(4)
				player.StartingHeroName = get.String(startingHeroNameLen)
			}
		}
	}

	// Player AI heroes, aka H3M_PLAYER_AI_ABSOD.
	if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {

		// TODO: 8x

		_ = get.Bytes(1) // unknown byte
		heroesCount := get.Int(4)
		fmt.Printf("Heroes count: %v\n", heroesCount)
		for i := 0; i < heroesCount; i++ {
			playerType := get.Int(1)
			nameLen := get.Int(4)
			name := get.String(nameLen)
			//fmt.Printf("AI player heroes: %v (type %v)\n", name, playerType)
			_ = name
			_ = playerType
		}
	}

	// Parse additional map information, aka H3M_AI.
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
