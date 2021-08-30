![rom64](doc/logo.png)

--------------------------------------------------------------------------------

Nintendo 64 ROM utility written in Go.

Commands
--------------------------------------------------------------------------------

* [ls](#rom64-ls) - List information about all ROMs in a directory
* [info](#rom64-info) - Show information about a single ROM
* [convert](#rom64-convert) - Convert a ROM file to the native (Z64, Big-endian) format
* [validate](#rom64-validate) - Validate the ROM's SHA-1 checksum against a list of known-good ROM dumps.

### `rom64 ls`

List ROM files in a given directory. By default, output is in a human-readable table.

#### Options

* `-f`, `--format` Defaults to `table` but can also be `text`, `json`, `csv`, `tab`, `xml`
* `-c`, `--columns` Defaults to most useful columns. Can be a comma-separated list, or specified multiple times.

**Available columns**

| Column ID | Description |
| --------- | ----------- |
| cic              | CIC chip type. example: 6102                                     |
| crc1             | CRC1 checksum of ROM internals. Also known as 'CRC HI'           |
| crc2             | CRC2 checksum of ROM internals. Also known as 'CRC LO'           |
| file_format      | File format code. One of: z64, v64, n64                          |
| file_format_desc | File format description. example: Big-endian                     |
| file_name        | File name on disk                                                |
| file_size_mbits  | File size in megabits. Always a whole number. example: 256       |
| file_size_mbytes | File size in megabytes. Always a whole number. example: 32       |
| image_name       | Image name / game title embedded in the ROM.                     |
| md5              | MD5 hash/checksum of the file on disk. Lower-case hexadecimal.   |
| region           | Region description of the ROM derived from the ROM ID.           |
| rom_id           | ROM ID / serial. example: NSME for Super Mario 64 (USA)          |
| sha1             | SHA-1 hash/checksum of the file on disk. Lower-case hexadecimal. |
| version          | Version of the ROM. One of: 1.0, 1.1, 1.2, or 1.3.               |
| video_system     | Video system derived from the ROM region. NTSC or PAL.           |

```
$ rom64 ls ~/Downloads/n64 -c image_name,file_format_desc,rom_id,region,video_system,cic,file_size_mbits,md5,file_name
+----------------------+--------------+--------+--------+-------+------+-----------+----------------------------------+-----------------------------------------------------------+
|      Image Name      | File Format  | Rom ID | Region | Video | CIC  | Size (Mb) |               MD5                |                         File Name                         |
+----------------------+--------------+--------+--------+-------+------+-----------+----------------------------------+-----------------------------------------------------------+
| 1080 SNOWBOARDING    | Big-endian   | NTEA   | JP/US  | NTSC  | 6103 |       128 | fa27089c425dbab99f19245c5c997613 | 1080 Snowboarding (Japan, USA) (En,Ja).z64                |
| BANJO TOOIE          | Big-endian   | NB7E   | US     | NTSC  | 6105 |       256 | 40e98faa24ac3ebe1d25cb5e5ddf49e4 | Banjo-Tooie (USA).z64                                     |
| CONKER BFD           | Big-endian   | NFUE   | US     | NTSC  | 6105 |       512 | 00e2920665f2329b95797a7eaabc2390 | Conker's Bad Fur Day (USA).z64                            |
| Diddy Kong Racing    | Big-endian   | NDYE   | US     | NTSC  | 6103 |        96 | b31f8cca50f31acc9b999ed5b779d6ed | Diddy Kong Racing (USA) (En,Fr) (Rev A).z64               |
| DONKEY KONG 64       | Big-endian   | NDOE   | US     | NTSC  | 6105 |       256 | 9ec41abf2519fc386cadd0731f6e868c | Donkey Kong 64 (USA).z64                                  |
| F1 WORLD GRAND PRIX2 | Byte-swapped | NF2P   | EU     | PAL   | 6102 |        96 | 2eb7766279c6b84d912c80aaa725bb6d | F-1 World Grand Prix II (Europe) (En,Fr,De,Es).n64        |
| F-ZERO X             | Big-endian   | NFZP   | EU     | PAL   | 6106 |       128 | ee79a8fe287b5dcaea584439363342fc | F-ZERO X (E) [!].z64                                      |
| F-ZERO X             | Byte-swapped | NFZP   | EU     | PAL   | 6106 |       128 | 7b2d9e24e6535be213ba036efe618e31 | F-Zero X (Europe).n64                                     |
| F-ZERO X             | Big-endian   | NFZP   | EU     | PAL   | 6106 |       128 | ee79a8fe287b5dcaea584439363342fc | F-Zero X (Europe).z64                                     |
| F-ZERO X             | Big-endian   | CFZE   | US     | NTSC  | 6106 |       128 | 753437d0d8ada1d12f3f9cf0f0a5171f | F-Zero X (USA).z64                                        |
| GOLDENEYE            | Byte-swapped | NGEE   | US     | NTSC  | 6102 |       128 | 08becb418039bcaf948dff73b9fe177b | Goldeneye.v64                                             |
| HSV ADVENTURE RACING | Byte-swapped | NNSX   | EU     | PAL   | 6102 |       128 | 1d8f8ca47fe02d1950e7a9725822414e | HSV Adventure Racing! (Australia).n64                     |
| HSV ADVENTURE RACING | Big-endian   | NNSX   | EU     | PAL   | 6102 |       128 | 26f7d8f4640ebdfa823f84e5f89d62bf | HSV Adventure Racing! (Australia).z64                     |
| JET FORCE GEMINI     | Big-endian   | NJFE   | US     | NTSC  | 6105 |       256 | 772cc6eab2620d2d3cdc17bbc26c4f68 | Jet Force Gemini (USA).z64                                |
| Killer Instinct Gold | Big-endian   | NKIE   | US     | NTSC  | 6102 |        96 | dd0a82fcc10397afb37f12bb7f94e67a | Killer Instinct Gold (USA) (Rev B).z64                    |
| ZELDA MAJORA'S MASK  | Big-endian   | NZSE   | US     | NTSC  | 6105 |       256 | 2a0a8acb61538235bc1094d297fb6556 | Legend of Zelda, The - Majora's Mask (USA).z64            |
| THE LEGEND OF ZELDA  | Big-endian   | CZLE   | US     | NTSC  | 6105 |       256 | 57a9719ad547c516342e1a15d5c28c3d | Legend of Zelda, The - Ocarina of Time (U) (V1.2) [!].z64 |
| THE LEGEND OF ZELDA  | Big-endian   | CZLE   | US     | NTSC  | 6105 |       256 | 57a9719ad547c516342e1a15d5c28c3d | Legend of Zelda, The - Ocarina of Time (USA) (Rev B).z64  |
| STARFOX64            | Big-endian   | NFXP   | EU     | PAL   | 7102 |        96 | 884ccca35cbeedb8ed288326f9662100 | Lylat Wars (E) (M3) [!].z64                               |
| STARFOX64            | Byte-swapped | NFXP   | EU     | PAL   | 7102 |        96 | 204a14c2ac815afee74b58ef9394708d | Lylat Wars (Europe) (En,Fr,De).n64                        |
| STARFOX64            | Big-endian   | NFXP   | EU     | PAL   | 7102 |        96 | 884ccca35cbeedb8ed288326f9662100 | Lylat Wars (Europe) (En,Fr,De).z64                        |
| Mega Man 64          | Big-endian   | NM6E   | US     | NTSC  | 6102 |       256 | 3620674acb51e436d5150738ac1c0969 | Mega Man 64 (USA).z64                                     |
| PAPER MARIO          | Big-endian   | NMQE   | US     | NTSC  | 6103 |       320 | a722f8161ff489943191330bf8416496 | Paper Mario (USA).z64                                     |
| Perfect Dark         | Big-endian   | NPDE   | US     | NTSC  | 6105 |       256 | e03b088b6ac9e0080440efed07c1e40f | Perfect Dark (USA) (Rev A).z64                            |
| Pilot Wings64        | Big-endian   | NPWE   | US     | NTSC  | 6102 |        64 | 8b346182730ceaffe5e2ccf6d223c5ef | Pilotwings 64 (USA).z64                                   |
| STARFOX64            | Big-endian   | NFXE   | US     | NTSC  | 6101 |        96 | 741a94eee093c4c8684e66b89f8685e8 | Star Fox 64 (USA) (Rev A).z64                             |
| Rogue Squadron       | Big-endian   | NRSE   | US     | NTSC  | 6102 |       128 | 47cac4e2a6309458342f21a9018ffbf0 | Star Wars - Rogue Squadron (USA) (Rev A).z64              |
| SUPER MARIO 64       | Big-endian   | NSME   | US     | NTSC  | 6102 |        64 | 20b854b239203baf6c961b850a4a51a2 | Super Mario 64 (USA).z64                                  |
| SUPERMARIO64         | Byte-swapped | NSMJ   | JP     | NTSC  | 6102 |        64 | f3a6e533de52fb84d4e2dfae6167c8b1 | Super Mario 64 - Shindou Edition (J) [!].n64              |
| SUPERMARIO64         | Big-endian   | NSMJ   | JP     | NTSC  | 6102 |        64 | 2d727c3278aa232d94f2fb45aec4d303 | Super Mario 64 - Shindou Edition (J) [!].z64              |
| SMASH BROTHERS       | Byte-swapped | NALE   | US     | NTSC  | 6103 |       128 | 3c4e65d9b7bf55338496108f04da7f41 | Super Smash Bros. (USA).n64                               |
| SMASH BROTHERS       | Big-endian   | NALE   | US     | NTSC  | 6103 |       128 | f7c52568a31aadf26e14dc2b6416b2ed | Super Smash Bros. (USA).z64                               |
| TSUMI TO BATSU       | Byte-swapped | NGUJ   | JP     | NTSC  | 6102 |       256 | 3b8638452b46deba9edffbcd6cdb6966 | Tsumi to Batsu - Hoshi no Keishousha (Japan).n64          |
| TSUMI TO BATSU       | Big-endian   | NGUJ   | JP     | NTSC  | 6102 |       256 | a0657bc99e169153fd46aeccfde748f3 | Tsumi to Batsu - Hoshi no Keishousha (Japan).z64          |
| WAVE RACE 64         | Big-endian   | NWRE   | US     | NTSC  | 6102 |        64 | 2048a640c12d1cf2052ba1629937d2ff | Wave Race 64 (USA) (Rev A).z64                            |
+----------------------+--------------+--------+--------+-------+------+-----------+----------------------------------+-----------------------------------------------------------+
```

### `rom64 info`

Show information about a single file.

Also supports the same output options as `ls`

```
$ rom64 info ~/Downloads/n64/Conker\'s\ Bad\ Fur\ Day\ \(USA\).z64
File:
  Name:    Conker's Bad Fur Day (USA).z64
  Size:    64 MB
  Format:  z64 (Big-endian)
  MD5:     00e2920665f2329b95797a7eaabc2390
  SHA1:    4cbadd3c4e0729dec46af64ad018050eada4f47a

ROM:
  ID:        NFUE
  Title:     CONKER BFD
  Media:     Cartridge
  Region:    US
  Video:     NTSC
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
    "format": {
      "code": "z64",
      "description": "Big-endian"
    },
    "size": 64,
    "md5": "00e2920665f2329b95797a7eaabc2390",
    "sha1": "4cbadd3c4e0729dec46af64ad018050eada4f47a"
  },
  "video_system": "NTSC"
}
```

### `rom64 convert`

Converts a ROM from a non-native format to the native big-endian Z64 format.

After conversion, the new ROM's SHA-1 checksum is validated against a known
list of good checksums, same as in the `validate` command.


### `rom64 validate`

Computes the ROM file's SHA-1 checksum and validates it against a list of known-good
checksums from a "datfile".

Checksums only work on files in the native Big-endian (Z64) format.

The binary includes a recent version of the datile from [dat-o-matic].
If you want to use your own, specify it with the `--datfile` flag.

```
$ rom64 validate ~/Downloads/n64/Tsumi\ to\ Batsu\ -\ Hoshi\ no\ Keishousha\ \(Japan\).z64
Found 1 datfile entries for ROM serial 'NGUJ'
  SHA-1 MATCH  581297B9D5C3A4C33169AE0AAE218C742CD9CBCF Tsumi to Batsu - Hoshi no Keishousha (Japan).z64
```

[dat-o-matic]: https://datomatic.no-intro.org/index.php?page=download&s=24&op=dat


Buidling
--------------------------------------------------------------------------------

1. [Install Go] for your platform.
2. Fetch module dependencies: `go get -d .`
3. `make all`

[Install Go]: https://golang.org/doc/install


Development
--------------------------------------------------------------------------------
This project follows [Semantic Versioning].

Git commit message should use [Conventional Commits] to have a clearly
readable history to the point where changelogs/release notes can be generated automatically.

[Semantic Versioning]: https://semver.org/
[Conventional Commits]: https://www.conventionalcommits.org/en/v1.0.0/
