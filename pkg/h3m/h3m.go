package h3m

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
	"github.com/pkg/errors"
)

// The binary format sourced from
// https://github.com/potmdehex/homm3tools/blob/master/h3m/h3mlib/h3m_structures/h3m.h
// The MIT License (MIT)
// Copyright (c) 2016 John Åkerblom

// Maps commonly end with 124 bytes of null padding. Extra content at end is ok.
type H3M struct {
	Format FileFormat
	MapInfo
	Players []Player // 8 players: Red, Blue, Tan, Green, Orange, Purple, Teal, Pink
	Tiles   []*Tile
}

func Parse(r io.Reader) (*H3M, error) {
	r, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	get := bytestream.New(r, binary.LittleEndian)

	h3m := &H3M{}

	// https://github.com/potmdehex/homm3tools/blob/5687f581a4eb5e7b0e8f48794d7be4e3b0a8cc8b/h3m/h3mlib/h3m_structures/h3m.h#L28
	h3m.Format = FileFormat(get.Int(4))
	fmt.Println("format:", h3m.Format)

	// Basic map info, aka H3M_BI.
	// https://github.com/potmdehex/homm3tools/blob/5687f581a4eb5e7b0e8f48794d7be4e3b0a8cc8b/h3m/h3mlib/h3m_structures/h3m.h#L29
	h3m.MapInfo.HasHero = get.Bool(1)
	h3m.MapInfo.MapSize = MapSize(get.Int(4))
	h3m.MapInfo.HasTwoLevels = get.Bool(1)

	nameLen := get.Int(4)
	h3m.MapInfo.Name = get.String(nameLen)
	fmt.Println("name:", h3m.MapInfo.Name)

	descLen := get.Int(4)
	h3m.MapInfo.Desc = get.String(descLen)
	fmt.Println("desc:", h3m.MapInfo.Desc)

	h3m.MapInfo.Difficulty = get.Int(1)

	// Players, aka 8x H3M_PLAYER.
	// https://github.com/potmdehex/homm3tools/blob/5687f581a4eb5e7b0e8f48794d7be4e3b0a8cc8b/h3m/h3mlib/h3m_structures/h3m.h#L30
	for i := 0; i < 8; i++ {
		player := Player{}

		if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
			player.MasteryCap = get.Int(1)
		}

		player.CanBeHuman = get.Bool(1)
		player.CanBeComputer = get.Bool(1)
		player.Behavior = get.Int(1) // 0-Random, 1-Warrior, 2-Builder, 3-Explorer

		if h3m.Format.Is(ShadowOfDeath) {
			player.AllowedAlignments = get.Int(1) // Bool? .. whether it is set which towns the player owns
		}

		player.TownTypes = Town(get.Int(1))

		if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {
			townTypeConflux := get.Int(1)
			_ = townTypeConflux // not used for now..
		}

		player.HasRandomTown = get.Bool(1)
		player.HasMainTown = get.Bool(1)

		if h3m.Format.Is(ROE) {

			// H3M_PLAYER_EXT_ROE
			// https://github.com/potmdehex/homm3tools/blob/master/h3m/h3mlib/h3m_structures/players/h3m_player.h#L65-L70

			if player.HasMainTown {
				player.StartingTownPos.X = get.Int(1)
				player.StartingTownPos.Y = get.Int(1)
				player.StartingTownPos.Z = get.Int(1)
			}

			player.StartingHeroIsRandom = get.Bool(1)
			player.StartingHeroType = get.Int(1)

			if player.StartingHeroType != 0xFF {
				player.StartingHeroFace = get.Int(1)
				startingHeroNameLen := get.Int(4)
				player.StartingHeroName = get.String(startingHeroNameLen)
			}

		} else if h3m.Format.Is(ArmageddonsBlade, ShadowOfDeath) {

			if player.HasMainTown {
				player.StartingTownCreateHero = get.Bool(1)
				player.StartingTown = Town(get.Int(1))
				player.StartingTownPos.X = get.Int(1)
				player.StartingTownPos.Y = get.Int(1)
				player.StartingTownPos.Z = get.Int(1)
			}

			player.StartingHeroIsRandom = get.Bool(1)
			player.StartingHeroType = get.Int(1)
			player.StartingHeroFace = get.Int(1)

			startingHeroNameLen := get.Int(4)
			player.StartingHeroName = get.String(startingHeroNameLen)
			//fmt.Println("Player", i, "name (=", len, "):", player.StartingHeroName, ", StartingHeroType:", player.StartingHeroType)

			if player.StartingHeroType != 0xFF {
				_ = get.Bytes(1) // Number of the hero's face. Standard face if 0xFF.
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
			// https://github.com/potmdehex/homm3tools/commit/d8b9f48b2567e5094aaed95e2205a71f279b9685
			return nil, errors.Errorf("TODO: Map format %v", h3m.Format)
		}

		if player.CanBeHuman || player.CanBeComputer {
			h3m.Players = append(h3m.Players, player)
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
	h3m.Teams = make([]int, 8)
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

	// ---- Map Tiles ----

	x, y := h3m.MapSize.Size()

	h3m.Tiles = make([]*Tile, x*y)

	for i := 0; i < x*y; i++ {
		h3m.Tiles[i] = &Tile{
			TerrainType:   Terrain(get.Int(1)),
			TerrainSprite: get.Int(1),
			RiverType:     get.Int(1),
			RiverSprite:   get.Int(1),
			RoadType:      get.Int(1),
			RoadSprite:    get.Int(1),
			Mirroring:     get.Int(1),
		}
	}

	return h3m, get.Error()
}
