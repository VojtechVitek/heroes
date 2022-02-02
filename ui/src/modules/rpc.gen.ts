/* eslint-disable */
//   300bea23083605662962866b3b2cc7c35624ecb9
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/typescript
// Do not edit by hand. Update your webrpc schema and re-generate.

// WebRPC description and code-gen version
export const WebRPCVersion = ""

// Schema version of your RIDL schema
export const WebRPCSchemaVersion = ""

// Schema hash generated from your RIDL schema
export const WebRPCSchemaHash = "300bea23083605662962866b3b2cc7c35624ecb9"


//
// Types
//
export interface h3mPosition {
  X: number
  Y: number
  Z: number
}

export interface h3mCustomHeroes {
  Type: number
  Face: number
  Name: string
  AllowedPlayers: number
}

export interface h3mMapInfo {
  HasHero: boolean
  MapSize: number
  HasTwoLevels: boolean
  Name: string
  Desc: string
  Difficulty: number
  MasteryCap: number
  WinCondition: number
  WinConditionAllowNormalWin: boolean
  WinConditionAppliesToComputer: boolean
  WinConditionType: number
  WinConditionAmount: number
  WinConditionUpgradeHallLevel: number
  WinConditionUpgradeCastleLevel: number
  WinConditionPos: h3mPosition
  LoseCondition: number
  LoseConditionPos: h3mPosition
  LoseConditionDays: number
  TeamsCount: number
  Teams: Array<number>
  AvailableHeroes: Array<string>
  CustomHeroesCount: number
  CustomHeroes: Array<h3mCustomHeroes>
}

export interface h3mPlayer {
  CanBeHuman: boolean
  CanBeComputer: boolean
  Behavior: number
  AllowedAlignments: number
  TownTypes: number
  Unknown1_HasRandomTown: boolean
  HasMainTown: boolean
  StartingTownCreateHero: boolean
  StartingTown: number
  StartingTownPos: h3mPosition
  StartingHeroIsRandom: boolean
  StartingHeroType: number
  StartingHeroFace: number
  StartingHeroName: string
}

export interface h3mTile {
  TerrainType: number
  TerrainSprite: number
  RiverType: number
  RiverSprite: number
  RoadType: number
  RoadSprite: number
  Mirroring: number
}

export interface h3mH3M {
  Format: number
  HasHero: boolean
  MapSize: number
  HasTwoLevels: boolean
  Name: string
  Desc: string
  Difficulty: number
  MasteryCap: number
  WinCondition: number
  WinConditionAllowNormalWin: boolean
  WinConditionAppliesToComputer: boolean
  WinConditionType: number
  WinConditionAmount: number
  WinConditionUpgradeHallLevel: number
  WinConditionUpgradeCastleLevel: number
  WinConditionPos: h3mPosition
  LoseCondition: number
  LoseConditionPos: h3mPosition
  LoseConditionDays: number
  TeamsCount: number
  Teams: Array<number>
  AvailableHeroes: Array<string>
  CustomHeroesCount: number
  CustomHeroes: Array<h3mCustomHeroes>
  Players: Array<h3mPlayer>
  Tiles: Array<h3mTile>
}

export interface Map {
  Format: number
  HasHero: boolean
  MapSize: number
  HasTwoLevels: boolean
  Name: string
  Desc: string
  Difficulty: number
  MasteryCap: number
  WinCondition: number
  WinConditionAllowNormalWin: boolean
  WinConditionAppliesToComputer: boolean
  WinConditionType: number
  WinConditionAmount: number
  WinConditionUpgradeHallLevel: number
  WinConditionUpgradeCastleLevel: number
  WinConditionPos: h3mPosition
  LoseCondition: number
  LoseConditionPos: h3mPosition
  LoseConditionDays: number
  TeamsCount: number
  Teams: Array<number>
  AvailableHeroes: Array<string>
  CustomHeroesCount: number
  CustomHeroes: Array<h3mCustomHeroes>
  Players: Array<h3mPlayer>
  Tiles: Array<h3mTile>
}

export interface API {
  getMap(headers?: object): Promise<GetMapReturn>
}

export interface GetMapArgs {
}

export interface GetMapReturn {
  m: Map  
}


  
//
// Client
//
export class API implements API {
  protected hostname: string
  protected fetch: Fetch
  protected path = '/rpc/API/'

  constructor(hostname: string, fetch: Fetch) {
    this.hostname = hostname
    this.fetch = fetch
  }

  private url(name: string): string {
    return this.hostname + this.path + name
  }
  
  getMap = (headers?: object): Promise<GetMapReturn> => {
    return this.fetch(
      this.url('GetMap'),
      createHTTPRequest({}, headers)
      ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          m: <Map>(_data.m)
        }
      })
    })
  }
  
}

  
export interface WebRPCError extends Error {
  code: string
  msg: string
	status: number
}

const createHTTPRequest = (body: object = {}, headers: object = {}): object => {
  return {
    method: 'POST',
    headers: { ...headers, 'Content-Type': 'application/json' },
    body: JSON.stringify(body || {})
  }
}

const buildResponse = (res: Response): Promise<any> => {
  return res.text().then(text => {
    let data
    try {
      data = JSON.parse(text)
    } catch(err) {
      throw { code: 'unknown', msg: `expecting JSON, got: ${text}`, status: res.status } as WebRPCError
    }
    if (!res.ok) {
      throw data // webrpc error response
    }
    return data
  })
}

export type Fetch = (input: RequestInfo, init?: RequestInit) => Promise<Response>
