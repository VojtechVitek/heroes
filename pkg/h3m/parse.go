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

	//spew.Dump(b.Bytes())

	get := bytestream.New(b.Bytes(), binary.LittleEndian)

	h3m := &H3M{}

	// https://github.com/potmdehex/homm3tools/blob/5687f581a4eb5e7b0e8f48794d7be4e3b0a8cc8b/h3m/h3mlib/h3m_structures/h3m.h#L28
	h3m.Format = FileFormat(get.Int(4))
	fmt.Println("format:", h3m.Format)

	// Basic map info, aka H3M_BI.
	// https://github.com/potmdehex/homm3tools/blob/5687f581a4eb5e7b0e8f48794d7be4e3b0a8cc8b/h3m/h3mlib/h3m_structures/h3m.h#L29
	h3m.MapInfo.HasHero = get.Bool(1)
	h3m.MapInfo.MapSize = get.Int(4)
	h3m.MapInfo.HasTwoLevels = get.Bool(1)

	nameLen := get.Int(4)
	h3m.MapInfo.Name = get.String(nameLen)
	fmt.Println("name:", h3m.MapInfo.Name)

	descLen := get.Int(4)
	h3m.MapInfo.Desc = get.String(descLen)
	fmt.Println("desc:", h3m.MapInfo.Desc)

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

			if player.HasMainTown {
				player.StartingTownCreateHero = get.Bool(1)
				player.StartingTown = Town(get.Int(1))
				player.StartingTownPos.X = get.Int(1)
				player.StartingTownPos.Y = get.Int(1)
				player.StartingTownPos.Z = get.Int(1)

				//  meta->player_ext_types[idx] = 1;
			}

			// When type != 0xFF
			// player_has_ai[idx] = 1;

			player.StartingHeroIsRandom = get.Bool(1)
			player.StartingHeroType = get.Int(1)
			player.StartingHeroFace = get.Int(1)

			len := get.Int(4)
			player.StartingHeroName = get.String(len)
			fmt.Println("Player", i, "name (=", len, "):", player.StartingHeroName, ", StartingHeroType:", player.StartingHeroType)

			if player.StartingHeroType != 0xFF {
				_ = get.Bytes(1) // unknown byte
				heroesCount := get.Int(4)
				fmt.Println("heroesCount:", heroesCount)
				if heroesCount > 10 {
					heroesCount = 10
				}
				for i := 0; i < heroesCount; i++ {
					heroesType := get.Int(1)
					len := get.Int(4)
					heroesName := get.String(len)
					fmt.Println(i, ": ", heroesName, "(", len, ")", heroesType)
				}
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

	// Parse additional map information, aka H3M_AI.
	h3m.MapInfo.WinCondition = WinCondition(get.Int(1))
	h3m.MapInfo.WinConditionAllowNormalWin = get.Bool(1) // Allow normal win. Supposedly doesn't work.
	h3m.MapInfo.WinConditionAppliesToComputer = get.Bool(1)

	if h3m.MapInfo.WinCondition.Is(WIN_ACQUIRE_ARTIFACT, WIN_ACCUMULATE_CREATURES, WIN_BUILD_GRAIL, WIN_DEFEAT_HERO) {
		if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
			h3m.MapInfo.WinConditionType = get.Int(2)
		} else {
			h3m.MapInfo.WinConditionType = get.Int(1)
		}
	}

	if h3m.MapInfo.WinCondition.Is(WIN_ACCUMULATE_CREATURES, WIN_ACCUMULATE_RESOURCES) {
		h3m.MapInfo.WinConditionAmount = get.Int(4)
	}

	if h3m.MapInfo.WinCondition.Is(WIN_TRANSPORT_ARTIFACT, WIN_UPGRADE_TOWN) {
		h3m.MapInfo.WinConditionPos.X = get.Int(1)
		h3m.MapInfo.WinConditionPos.Y = get.Int(1)
		h3m.MapInfo.WinConditionPos.Z = get.Int(1)
	}

	if h3m.MapInfo.WinCondition.Is(WIN_UPGRADE_TOWN) {
		h3m.MapInfo.WinConditionUpgradeHallLevel = get.Int(1)
		h3m.MapInfo.WinConditionUpgradeCastleLevel = get.Int(1)
	}

	h3m.MapInfo.LoseCondition = LoseCondition(get.Int(1))
	switch h3m.MapInfo.LoseCondition {
	case LOSE_TOWN, LOSE_HERO:
		h3m.MapInfo.LoseConditionPos.X = get.Int(1)
		h3m.MapInfo.LoseConditionPos.Y = get.Int(1)
		h3m.MapInfo.LoseConditionPos.Z = get.Int(1)
	case LOSE_TIME:
		h3m.MapInfo.LoseConditionDays = get.Int(2)
	}

	h3m.TeamsCount = get.Int(1)
	h3m.Teams[0] = get.Int(1)
	h3m.Teams[1] = get.Int(1)
	h3m.Teams[2] = get.Int(1)
	h3m.Teams[3] = get.Int(1)
	h3m.Teams[4] = get.Int(1)
	h3m.Teams[5] = get.Int(1)
	h3m.Teams[6] = get.Int(1)
	h3m.Teams[7] = get.Int(1)

	if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
		h3m.AvailableHeroes = get.Bytes(20) // AB/SOD
		_ = get.Bytes(4)                    // empty; AB/SOD
	}

	if h3m.Format.Is(ShadowOfDeath) {
		h3m.CustomHeroesCount = get.Int(1) // SOD

		for i := 0; i < h3m.CustomHeroesCount; i++ {
			h3m.CustomHeroes[i] = CustomHeroes{
				Type:           get.Int(1),
				Face:           get.Int(1),
				Name:           get.String(get.Int(4)),
				AllowedPlayers: get.Int(1),
			}
		}
	}

	_ = get.Bytes(31) // reserved

	if h3m.Format.Is(ShadowOfDeath) {
		availableArtifacts := get.Bytes(18)
		_ = availableArtifacts // TODO

		availableSpells := get.Bytes(9)
		_ = availableSpells // TODO

		availableSkills := get.Bytes(4)
		_ = availableSkills // TODO
	}

	rumorsCount := get.Int(4)
	for i := 0; i < rumorsCount; i++ {
		if i > 10 {
			break
		}
		nameLen := get.Int(4)
		name := get.String(nameLen)
		descLen := get.Int(4)
		desc := get.String(descLen)
		fmt.Println("rumor:", name, ":", desc)
	}

	if h3m.Format.Is(ShadowOfDeath) {
		heroSettings := get.Bytes(156) // SOD
		_ = heroSettings
	}

	return h3m, get.Error()
}
