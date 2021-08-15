# rom64

Nintendo 64 ROM utility written in Go.

ROM file headers are parsed to detect information:

* File format
  * `z64` - Big-endian (Native format of N64)
  * `n64` - Little Endian
  * `v64` - Big-endian, byte-swapped
* Media format
  * `N` - Cartridge
  * `D` - 64DD Disk
  * `C` - Cartridge for expandable game
  * `E` - 64DD Expansion
  * `Z` - Aleck64 Cartridge
* Game ID - Each title has a unique two-character ID such as `SM` for *Super Mario 64*
* Game region
  * `E` - United States
  * `J` - Japan
  * `P` - Europe and PAL regions
  * There are more, but those are the most common
* CIC - Version of the copy protection chip on the ROM such as `6102` or `6105`
* CRC1 and CRC2 - CRC checks built into the ROM header to validate the rest of the ROM.
* Title - Game titles are stored in the ROM in a limited format such as *SUPERMARIO64*

## Commands

* [ls](#rom64-ls)
* [stat](#rom64-stat)
* [convert](#rom64-convert)

### `rom64 ls`

List ROM files in a given directory. By default, output is in a human-readable table.

#### Options

* `--with-md5` calculates the file MD5s. It's opt-in since it can be a little slow.
* `--format` Defaults to `table` but can also be `text`, `json`, `csv`, `tab`

```
rom64 ls ~/Downloads/n64 --with-md5
+----------------------+--------+------+--------+---------+--------+------+----------+----------+----------------------------------+-----------------------------------------------------------+
|        Title         | Format | Size | ROM ID | Version | Region | CIC  |   CRC1   |   CRC2   |               MD5                |                         File Name                         |
+----------------------+--------+------+--------+---------+--------+------+----------+----------+----------------------------------+-----------------------------------------------------------+
| 1080 SNOWBOARDING    | z64    |   16 | NTEA   |     1.0 | JP/US  | 6103 | 1FBAF161 | 2C1C54F1 | fa27089c425dbab99f19245c5c997613 | 1080 Snowboarding (Japan, USA) (En,Ja).z64                |
| Banjo-Kazooie        | z64    |   16 | NBKE   |     1.1 | US     | 6103 | CD7559AC | B26CF5AE | b11f476d4bc8e039355241e871dc08cf | Banjo-Kazooie (USA) (Rev A).z64                           |
| BANJO TOOIE          | z64    |   32 | NB7E   |     1.0 | US     | 6105 | C2E9AA9A | 475D70AA | 40e98faa24ac3ebe1d25cb5e5ddf49e4 | Banjo-Tooie (USA).z64                                     |
| Blast Corps          | z64    |    8 | NBCE   |     1.1 | US     | 6102 | 7C647E65 | 1948D305 | 5875fc73069077c93e214233b60f0bdc | Blast Corps (USA) (Rev A).z64                             |
| BOMBERMAN64U         | z64    |    8 | NBME   |     1.0 | US     | 6102 | F568D51E | 7E49BA1E | 093058ece14c8cc1a887b2087eb5cfe9 | Bomberman 64 (USA).z64                                    |
| CONKER BFD           | z64    |   64 | NFUE   |     1.0 | US     | 6105 | 30C7AC50 | 7704072D | 00e2920665f2329b95797a7eaabc2390 | Conker's Bad Fur Day (USA).z64                            |
| Diddy Kong Racing    | z64    |   12 | NDYE   |     1.1 | US     | 6103 | E402430D | D2FCFC9D | b31f8cca50f31acc9b999ed5b779d6ed | Diddy Kong Racing (USA) (En,Fr) (Rev A).z64               |
| DONKEY KONG 64       | z64    |   32 | NDOE   |     1.0 | US     | 6105 | EC58EABF | AD7C7169 | 9ec41abf2519fc386cadd0731f6e868c | Donkey Kong 64 (USA).z64                                  |
| F1 WORLD GRAND PRIX2 | v64    |   12 | NF2P   |     1.0 | EU     | 6102 | 874965A3 | 1B147D88 | 2eb7766279c6b84d912c80aaa725bb6d | F-1 World Grand Prix II (Europe) (En,Fr,De,Es).n64        |
| F-ZERO X             | z64    |   16 | NFZP   |     1.0 | EU     | 6106 | 776646F6 | 06B9AC2B | ee79a8fe287b5dcaea584439363342fc | F-ZERO X (E) [!].z64                                      |
| F-ZERO X             | v64    |   16 | NFZP   |     1.0 | EU     | 6106 | 776646F6 | 06B9AC2B | 7b2d9e24e6535be213ba036efe618e31 | F-Zero X (Europe).n64                                     |
| F-ZERO X             | z64    |   16 | NFZP   |     1.0 | EU     | 6106 | 776646F6 | 06B9AC2B | ee79a8fe287b5dcaea584439363342fc | F-Zero X (Europe).z64                                     |
| F-ZERO X             | z64    |   16 | CFZE   |     1.0 | US     | 6106 | B30ED978 | 3003C9F9 | 753437d0d8ada1d12f3f9cf0f0a5171f | F-Zero X (USA).z64                                        |
| GOLDENEYE            | v64    |   16 | NGEE   |     1.0 | US     | 6102 | DCBC50D1 | 09FD1AA3 | 08becb418039bcaf948dff73b9fe177b | Goldeneye.v64                                             |
| JET FORCE GEMINI     | z64    |   32 | NJFE   |     1.0 | US     | 6105 | 8A6009B6 | 94ACE150 | 772cc6eab2620d2d3cdc17bbc26c4f68 | Jet Force Gemini (USA).z64                                |
| Killer Instinct Gold | z64    |   12 | NKIE   |     1.2 | US     | 6102 | F908CA4C | 36464327 | dd0a82fcc10397afb37f12bb7f94e67a | Killer Instinct Gold (USA) (Rev B).z64                    |
| LEGORacers           | z64    |   16 | NLGE   |     1.0 | US     | 6102 | 096A40EA | 8ABE0A10 | 97c4cae584f427ec44266e9b98fbf7b6 | LEGO Racers (USA) (En,Fr,De,Es,It,Nl,Sv,No,Da,Fi).z64     |
| ZELDA MAJORA'S MASK  | z64    |   32 | NZSE   |     1.0 | US     | 6105 | 5354631C | 03A2DEF0 | 2a0a8acb61538235bc1094d297fb6556 | Legend of Zelda, The - Majora's Mask (USA).z64            |
| THE LEGEND OF ZELDA  | z64    |   32 | CZLE   |     1.2 | US     | 6105 | 693BA2AE | B7F14E9F | 57a9719ad547c516342e1a15d5c28c3d | Legend of Zelda, The - Ocarina of Time (U) (V1.2) [!].z64 |
| THE LEGEND OF ZELDA  | z64    |   32 | CZLE   |     1.2 | US     | 6105 | 693BA2AE | B7F14E9F | 57a9719ad547c516342e1a15d5c28c3d | Legend of Zelda, The - Ocarina of Time (USA) (Rev B).z64  |
| STARFOX64            | z64    |   12 | NFXP   |     1.0 | EU     | 7102 | F4CBE92C | B392ED12 | 884ccca35cbeedb8ed288326f9662100 | Lylat Wars (Europe) (En,Fr,De).z64                        |
| MarioGolf64          | z64    |   24 | NMFE   |     1.0 | US     | 6102 | 664BA3D4 | 678A80B7 | 7a5d0d77a462b5a7c372fb19efde1a5f | Mario Golf (USA).z64                                      |
| MARIOKART64          | z64    |   12 | NKTE   |     1.0 | US     | 6102 | 3E5055B6 | 2E92DA52 | 3a67d9986f54eb282924fca4cd5f6dff | Mario Kart 64 (USA).z64                                   |
| MarioParty2          | z64    |   32 | NMWE   |     1.0 | US     | 6102 | 9EA95858 | AF72B618 | 04840612a35ece222afdb2dfbf926409 | Mario Party 2 (USA).z64                                   |
| MarioTennis          | z64    |   16 | NM8E   |     1.0 | US     | 6102 | 5001CF4F | F30CB3BD | 759358fad1ed5ae31dcb2001a07f2fe5 | Mario Tennis (USA).z64                                    |
| Mega Man 64          | z64    |   32 | NM6E   |     1.0 | US     | 6102 | 0EC158F5 | FB3E6896 | 3620674acb51e436d5150738ac1c0969 | Mega Man 64 (USA).z64                                     |
| PAPER MARIO          | z64    |   40 | NMQE   |     1.0 | US     | 6103 | 65EEE53A | ED7D733C | a722f8161ff489943191330bf8416496 | Paper Mario (USA).z64                                     |
| Perfect Dark         | z64    |   32 | NPDE   |     1.1 | US     | 6105 | 41F2B98F | B458B466 | e03b088b6ac9e0080440efed07c1e40f | Perfect Dark (USA) (Rev A).z64                            |
| Pilot Wings64        | z64    |    8 | NPWE   |     1.0 | US     | 6102 | C851961C | 78FCAAFA | 8b346182730ceaffe5e2ccf6d223c5ef | Pilotwings 64 (USA).z64                                   |
| TSUMI TO BATSU       | v64    |   32 | NGUJ   |     1.0 | JP     | 6102 | B6BC0FB0 | E3812198 | 3b8638452b46deba9edffbcd6cdb6966 | Sin and Punishment.v64                                    |
| TSUMI TO BATSU       | z64    |   32 | NGUJ   |     1.0 | JP     | 6102 | B6BC0FB0 | E3812198 | a0657bc99e169153fd46aeccfde748f3 | Sin and Punishment.z64                                    |
| STARFOX64            | z64    |   12 | NFXE   |     1.1 | US     | 6101 | BA780BA0 | 0F21DB34 | 741a94eee093c4c8684e66b89f8685e8 | Star Fox 64 (USA) (Rev A).z64                             |
| Rogue Squadron       | z64    |   16 | NRSE   |     1.0 | US     | 6102 | 66A24BEC | 2EADD94F | 47cac4e2a6309458342f21a9018ffbf0 | Star Wars - Rogue Squadron (USA) (Rev A).z64              |
| SUPER MARIO 64       | z64    |    8 | NSME   |     1.0 | US     | 6102 | 635A2BFF | 8B022326 | 20b854b239203baf6c961b850a4a51a2 | Super Mario 64 (USA).z64                                  |
| SUPERMARIO64         | v64    |    8 | NSMJ   |     1.3 | JP     | 6102 | D6FBA4A8 | 6326AA2C | f3a6e533de52fb84d4e2dfae6167c8b1 | Super Mario 64 - Shindou Edition (J) [!].n64              |
| SUPERMARIO64         | z64    |    8 | NSMJ   |     1.3 | JP     | 6102 | D6FBA4A8 | 6326AA2C | 2d727c3278aa232d94f2fb45aec4d303 | Super Mario 64 - Shindou Edition (J) [!].z64              |
| SMASH BROTHERS       | v64    |   16 | NALE   |     1.0 | US     | 6103 | 916B8B5B | 780B85A4 | 3c4e65d9b7bf55338496108f04da7f41 | Super Smash Bros. (USA).n64                               |
| SMASH BROTHERS       | z64    |   16 | NALE   |     1.0 | US     | 6103 | 916B8B5B | 780B85A4 | f7c52568a31aadf26e14dc2b6416b2ed | Super Smash Bros. (USA).z64                               |
| TSUMI TO BATSU       | v64    |   32 | NGUJ   |     1.0 | JP     | 6102 | B6BC0FB0 | E3812198 | 3b8638452b46deba9edffbcd6cdb6966 | Tsumi to Batsu - Hoshi no Keishousha (Japan).n64          |
| TSUMI TO BATSU       | z64    |   32 | NGUJ   |     1.0 | JP     | 6102 | B6BC0FB0 | E3812198 | a0657bc99e169153fd46aeccfde748f3 | Tsumi to Batsu - Hoshi no Keishousha (Japan).z64          |
| WAVE RACE 64         | z64    |    8 | NWRE   |     1.1 | US     | 6102 | 492F4B61 | 04E5146A | 2048a640c12d1cf2052ba1629937d2ff | Wave Race 64 (USA) (Rev A).z64                            |
+----------------------+--------+------+--------+---------+--------+------+----------+----------+----------------------------------+-----------------------------------------------------------+
```

## `rom64 stat`

List information about a single file.

Also supports the same output options as `ls`

```
rom64 stat ~/Downloads/n64/Conker\'s\ Bad\ Fur\ Day\ \(USA\).z64
File:
  Name:    Conker's Bad Fur Day (USA).z64
  Size:    64 MB
  Format:  z64
  MD5:     00e2920665f2329b95797a7eaabc2390

ROM:
  ID:        NFUE
  Title:     CONKER BFD
  Media:     Cartridge
  Region:    US
  Version:   1.0
  CIC:       6105
  CRC 1:     30C7AC50
  CRC 2:     7704072D
```

Same ROM but with `--output json`

```json
{
  "crc_1": "30C7AC50",
  "crc_2": "7704072D",
  "image_name": "CONKER BFD",
  "media_format": {
    "code": "N",
    "description": "Cartridge"
  },
  "cartridge_id": "FU",
  "region": {
    "code": "E",
    "description": "US"
  },
  "version": 0,
  "cic": "6105",
  "file": {
    "path": "/home/mroach/Downloads/n64/Conker's Bad Fur Day (USA).z64",
    "name": "Conker's Bad Fur Day (USA).z64",
    "format": "z64",
    "size": 64,
    "md5": "00e2920665f2329b95797a7eaabc2390"
  }
}
```

## `rom64 convert`

Converts a ROM from a non-native format to the native big-endian Z64 format.
