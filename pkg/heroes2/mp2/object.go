package mp2

//go:generate stringer -type=Object
type Object uint8

const (
	_                     Object = iota // 0
	_alchemyLab                         // 1
	_                                   // 2
	_                                   // 3
	Skeleton                            // 4
	DaemonCave                          // 5
	_                                   // 6
	FaerieRing                          // 7
	_                                   // 8
	_                                   // 9
	Gazebo                              // 10
	_                                   // 11
	_graveyard                          // 12
	ArchersHouse                        // 13
	_                                   // 14
	DwarfCottage                        // 15
	PeasantHut                          // 16
	_                                   // 17
	_                                   // 18
	_                                   // 19
	DragonCity                          // 20
	_lightHouse                         // 21
	WaterWheel                          // 22
	_mines                              // 23
	_                                   // 24
	_obelisk                            // 25
	Oasis                               // 26
	_                                   // 27
	k8s                                 // 28 - ship can land
	_sawmill                            // 29
	Oracle                              // 30
	_                                   // 31
	_shipWreck                          // 32
	_                                   // 33
	DesertTent                          // 34
	_castleBottom                       // 35
	StoneLights                         // 36
	_wagonCamp                          // 37
	_                                   // 38
	_                                   // 39
	_windmill                           // 40
	_                                   // 41
	_                                   // 42
	_                                   // 43
	_                                   // 44
	_                                   // 45
	_                                   // 46
	_                                   // 47
	RandTown                            // 48
	RandCastle                          // 49
	_                                   // 50
	_                                   // 51
	_                                   // 52
	_                                   // 53
	_                                   // 54
	_                                   // 55
	_bush                               // 56 - rosti
	_                                   // 57
	WatchTower                          // 58
	TreeHouse                           // 59
	TreeCity                            // 60
	Ruins                               // 61
	Fort                                // 62
	_tradingPost                        // 63
	AbandonedMine                       // 64
	_                                   // 65
	_                                   // 66
	_                                   // 67
	TreeOfKnowledge                     // 68
	DoctorHut                           // 69
	Temple                              // 70
	HillFort                            // 71
	HalflingHole                        // 72
	_mercenaryCamp                      // 73
	_                                   // 74
	_                                   // 75
	Pyramid                             // 76
	CityOfTheDead                       // 77
	Excavation                          // 78
	Sphinx                              // 79
	_                                   // 80
	_                                   // 81
	ArtesianSpring                      // 82
	TrollBridge                         // 83
	_pondUpper                          // 84
	_witchsHutMiddle                    // 85
	Xanadu                              // 86
	Cave                                // 87
	_                                   // 88
	MagellansMaps                       // 89
	_                                   // 90
	DerelictShip                        // 91
	_                                   // 92
	_                                   // 93
	MagicWell                           // 94
	_                                   // 95
	ObservationTower                    // 96
	_golemFreemansFoundry               // 97
	_                                   // 98
	Forest                              // 99
	Mountains                           // 100
	_                                   // 101
	flower                              // 102
	stoneInWater                        // 103
	_                                   // 104
	_                                   // 105
	DeadTree                            // 106
	_                                   // 107
	_                                   // 108
	_                                   // 109
	_                                   // 110
	_                                   // 111
	_                                   // 112
	_                                   // 113
	Arena                               // 114
	BarrowMounds                        // 115
	Mermaid                             // 116
	Sirens                              // 117
	HutMagi                             // 118
	EyeMagi                             // 119
	TravellerTent                       // 120
	_                                   // 121
	_syrens                             // 122
	Jail                                // 123
	FireAltar                           // 124
	AirAltar                            // 125
	EarthAltar                          // 126
	WaterAltar                          // 127
	_                                   // 128
	AlchemyLab                          // 129
	_                                   // 130
	_                                   // 131
	_                                   // 132
	_                                   // 133
	Chest                               // 134
	_                                   // 135
	FirePit                             // 136
	Fountain                            // 137
	_                                   // 138
	_                                   // 139
	Graveyard                           // 140
	_                                   // 141
	slam                                // 142 chatrc
	_                                   // 143
	_                                   // 144
	_                                   // 145
	_                                   // 146
	_                                   // 147
	_                                   // 148
	LightHouse                          // 149
	_                                   // 150
	Mines                               // 151
	_                                   // 152
	Obelisk                             // 153
	_                                   // 154
	RandResource                        // 155
	_                                   // 156
	Sawmill                             // 157
	_                                   // 158
	_                                   // 159
	ShipWreck                           // 160
	_                                   // 161
	_                                   // 162
	_                                   // 163
	_                                   // 164
	WagonCamp                           // 165
	_                                   // 166
	_                                   // 167
	Windmill                            // 168
	_                                   // 169
	_                                   // 170
	Ship                                // 171
	_                                   // 172
	_                                   // 173
	_                                   // 174
	_                                   // 175
	_                                   // 176
	TownGate                            // 177
	_                                   // 178
	RandC1                              // 179
	RandC2                              // 180
	RandC3Nature                        // 181
	RandDragon                          // 182
	_                                   // 183
	_                                   // 184
	_                                   // 185
	_                                   // 186 - wooden structure/tower
	_                                   // 187
	_                                   // 188
	_                                   // 189
	_                                   // 190
	TradingPost                         // 191
	_                                   // 192
	_                                   // 193
	Stonehenge                          // 194
	_                                   // 195
	_                                   // 196
	_                                   // 197
	_                                   // 198
	_                                   // 199
	_                                   // 200
	MercenaryCamp                       // 201
	GinHouse                            // 202 blue house with red roof
	_                                   // 203
	_                                   // 204
	_                                   // 205
	_                                   // 206
	_                                   // 207
	_                                   // 208
	_                                   // 209
	_                                   // 210
	_                                   // 211
	Pond                                // 212
	WitchsHut                           // 213
	_                                   // 214
	_                                   // 215
	_                                   // 216
	_                                   // 217
	Wood                                // 218
	_                                   // 219
	drawnManHoldingWood                 // 220
	_                                   // 221
	Well                                // 222
	Mushroom                            // 223
	_                                   // 224
	GolemHouse                          // 225
	_                                   // 226
	_                                   // 227
	_                                   // 228
	_                                   // 229
	_                                   // 230
	_                                   // 231
	_                                   // 232
	_                                   // 233
	AlchemyTower                        // 234
	Stables                             // 235
	_                                   // 236
	_                                   // 237
	_                                   // 238
	_                                   // 239
	_                                   // 240
	_                                   // 241
	_                                   // 242
	_                                   // 243
	RandChest                           // 244
	_                                   // 245
	_                                   // 246
	_                                   // 247
	_                                   // 248
	_                                   // 249
	Syrens                              // 250
	_                                   // 251
	_                                   // 252
	_                                   // 253
	_                                   // 254
	_                                   // 255
)
