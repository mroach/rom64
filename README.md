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

* `-o`, `--output` Defaults to `table` but can also be `text`, `json`, `csv`, `tab`, `xml`
* `-c`, `--columns` Defaults to most useful columns. Can be a comma-separated list, or specified multiple times.

#### Checksums

When using `table`, `csv`, or `tab` format, checksums are calculated if the column is requested with `-c | --columns`.

When using `text`, `json`, or `xml` format, by default *no checksums are calculated* because of the CPU overhead.
To calculate checksums in these modes, request them using the `-c | --columns` flag.

```
rom64 ls . --output text --columns file_md5,file_sha1,file_crc1,file_crc2
```

**Available columns**

| Column ID        | Description |
| ---------------- | ----------- |
| cic              | CIC chip type. example: 6102                                     |
| crc1             | Expected CRC1 checksum of ROM internals. Also known as 'CRC HI'  |
| crc2             | Expected CRC2 checksum of ROM internals. Also known as 'CRC LO'  |
| image_name       | Image name / game title embedded in the ROM.                     |
| region           | Region description of the ROM derived from the ROM ID.           |
| rom_id           | ROM ID / serial. example: *NSME* for Super Mario 64 (USA)        |
| version          | Version of the ROM. One of: 1.0, 1.1, 1.2, or 1.3.               |
| video_system     | Video system derived from the ROM region. NTSC or PAL.           |

Columns prefixed with `file_` are information about the file itself rather than ROM header metadata.

| Column ID        | Description |
| ---------------- | ----------- |
| file_crc1        | Actual calculated CRC1 of the file's first 1MB of data           |
| file_crc2        | Actual calculated CRC2 of the file's first 1MB of data           |
| file_format      | File format code. One of: z64, v64, n64                          |
| file_format_desc | File format description. example: *Big-endian*                   |
| file_md5         | MD5 hash/checksum of the file on disk. Lower-case hexadecimal.   |
| file_name        | File name on disk                                                |
| file_sha1        | SHA-1 hash/checksum of the file on disk. Lower-case hexadecimal. |
| file_size_mbits  | File size in megabits. Always a whole number. example: *256*     |
| file_size_mbytes | File size in megabytes. Always a whole number. example: *32*     |

```
$ rom64 ls ~/Downloads/n64 -c image_name,rom_id,region,video_system,cic,file_size_mbits,file_format_desc,crc1,file_name
+----------------------+--------+---------------+-------+------+-----------+--------------+----------+-----------------------------------------------------------+
|      Image Name      | Rom ID |    Region     | Video | CIC  | Size (Mb) | File Format  |  CRC-1   |                         File Name                         |
+----------------------+--------+---------------+-------+------+-----------+--------------+----------+-----------------------------------------------------------+
| 1080 SNOWBOARDING    | NTEA   | Japan and USA | NTSC  | 6103 |       128 | Big-endian   | 1FBAF161 | 1080 Snowboarding (Japan, USA) (En,Ja).z64                |
| 40 WINKS             | N4WX   | PAL Regions   | PAL   | 6102 |       256 | Byte-swapped | ABA51D09 | 40 Winks (Europe) (En,Es,It) (Proto).n64                  |
| 40 WINKS             | N4WX   | PAL Regions   | PAL   | 6102 |       256 | Big-endian   | ABA51D09 | 40 Winks (Europe) (En,Es,It) (Proto).z64                  |
| Banjo-Kazooie        | NBKE   | USA           | NTSC  | 6103 |       128 | Big-endian   | CD7559AC | Banjo-Kazooie (USA) (Rev A).z64                           |
| BANJO TOOIE          | NB7E   | USA           | NTSC  | 6105 |       256 | Big-endian   | C2E9AA9A | Banjo-Tooie (USA).z64                                     |
| Blast Corps          | NBCE   | USA           | NTSC  | 6102 |        64 | Big-endian   | 7C647E65 | Blast Corps (USA) (Rev A).z64                             |
| BOMBERMAN64U         | NBME   | USA           | NTSC  | 6102 |        64 | Big-endian   | F568D51E | Bomberman 64 (USA).z64                                    |
| CONKER BFD           | NFUE   | USA           | NTSC  | 6105 |       512 | Big-endian   | 30C7AC50 | Conker's Bad Fur Day (USA).z64                            |
| Diddy Kong Racing    | NDYE   | USA           | NTSC  | 6103 |        96 | Big-endian   | E402430D | Diddy Kong Racing (USA) (En,Fr) (Rev A).z64               |
| DONKEY KONG 64       | NDOE   | USA           | NTSC  | 6105 |       256 | Big-endian   | EC58EABF | Donkey Kong 64 (USA).z64                                  |
| F1 WORLD GRAND PRIX2 | NF2P   | Europe        | PAL   | 6102 |        96 | Byte-swapped | 874965A3 | F-1 World Grand Prix II (Europe) (En,Fr,De,Es).n64        |
| F-ZERO X             | NFZP   | Europe        | PAL   | 6106 |       128 | Big-endian   | 776646F6 | F-ZERO X (E) [!].z64                                      |
| F-ZERO X             | NFZP   | Europe        | PAL   | 6106 |       128 | Byte-swapped | 776646F6 | F-Zero X (Europe).n64                                     |
| F-ZERO X             | NFZP   | Europe        | PAL   | 6106 |       128 | Big-endian   | 776646F6 | F-Zero X (Europe).z64                                     |
| F-ZERO X             | CFZE   | USA           | NTSC  | 6106 |       128 | Big-endian   | B30ED978 | F-Zero X (USA).z64                                        |
| GOLDENEYE            | NGEE   | USA           | NTSC  | 6102 |       128 | Byte-swapped | DCBC50D1 | Goldeneye.v64                                             |
| HSV ADVENTURE RACING | NNSX   | PAL Regions   | PAL   | 6102 |       128 | Byte-swapped | 72611D7D | HSV Adventure Racing! (Australia).n64                     |
| HSV ADVENTURE RACING | NNSX   | PAL Regions   | PAL   | 6102 |       128 | Big-endian   | 72611D7D | HSV Adventure Racing! (Australia).z64                     |
| JET FORCE GEMINI     | NJFE   | USA           | NTSC  | 6105 |       256 | Big-endian   | 8A6009B6 | Jet Force Gemini (USA).z64                                |
| Killer Instinct Gold | NKIE   | USA           | NTSC  | 6102 |        96 | Big-endian   | F908CA4C | Killer Instinct Gold (USA) (Rev B).z64                    |
| LEGORacers           | NLGE   | USA           | NTSC  | 6102 |       128 | Big-endian   | 096A40EA | LEGO Racers (USA) (En,Fr,De,Es,It,Nl,Sv,No,Da,Fi).z64     |
| ZELDA MAJORA'S MASK  | NZSE   | USA           | NTSC  | 6105 |       256 | Big-endian   | 5354631C | Legend of Zelda, The - Majora's Mask (USA).z64            |
| THE LEGEND OF ZELDA  | CZLE   | USA           | NTSC  | 6105 |       256 | Big-endian   | 693BA2AE | Legend of Zelda, The - Ocarina of Time (U) (V1.2) [!].z64 |
| THE LEGEND OF ZELDA  | CZLE   | USA           | NTSC  | 6105 |       256 | Big-endian   | 693BA2AE | Legend of Zelda, The - Ocarina of Time (USA) (Rev B).z64  |
| STARFOX64            | NFXP   | Europe        | PAL   | 7102 |        96 | Big-endian   | F4CBE92C | Lylat Wars (E) (M3) [!].z64                               |
| STARFOX64            | NFXP   | Europe        | PAL   | 7102 |        96 | Byte-swapped | F4CBE92C | Lylat Wars (Europe) (En,Fr,De).n64                        |
| STARFOX64            | NFXP   | Europe        | PAL   | 7102 |        96 | Big-endian   | F4CBE92C | Lylat Wars (Europe) (En,Fr,De).z64                        |
| MarioGolf64          | NMFE   | USA           | NTSC  | 6102 |       192 | Big-endian   | 664BA3D4 | Mario Golf (USA).z64                                      |
| MARIOKART64          | NKTE   | USA           | NTSC  | 6102 |        96 | Big-endian   | 3E5055B6 | Mario Kart 64 (USA).z64                                   |
| MarioParty2          | NMWE   | USA           | NTSC  | 6102 |       256 | Big-endian   | 9EA95858 | Mario Party 2 (USA).z64                                   |
| MarioTennis          | NM8E   | USA           | NTSC  | 6102 |       128 | Big-endian   | 5001CF4F | Mario Tennis (USA).z64                                    |
| Mega Man 64          | NM6E   | USA           | NTSC  | 6102 |       256 | Big-endian   | 0EC158F5 | Mega Man 64 (USA).z64                                     |
| PAPER MARIO          | NMQE   | USA           | NTSC  | 6103 |       320 | Big-endian   | 65EEE53A | Paper Mario (USA).z64                                     |
| Perfect Dark         | NPDE   | USA           | NTSC  | 6105 |       256 | Big-endian   | 41F2B98F | Perfect Dark (USA) (Rev A).z64                            |
| Pilot Wings64        | NPWE   | USA           | NTSC  | 6102 |        64 | Big-endian   | C851961C | Pilotwings 64 (USA).z64                                   |
| PYORO64              | BPYB   | Brazil        | NTSC  | 6102 |        16 | Big-endian   | 03D73956 | Pyoro64 MPAL.n64                                          |
| PYORO64              | BPYE   | USA           | NTSC  | 6102 |        16 | Big-endian   | 03D607B8 | Pyoro64 NTSC.n64                                          |
| PYORO64              | BPYI   | Italy         | PAL   | 6102 |        16 | Big-endian   | 58ABD009 | Pyoro64 PAL.n64                                           |
| TSUMI TO BATSU       | NGUJ   | Japan         | NTSC  | 6102 |       256 | Byte-swapped | B6BC0FB0 | Sin and Punishment.v64                                    |
| TSUMI TO BATSU       | NGUJ   | Japan         | NTSC  | 6102 |       256 | Big-endian   | B6BC0FB0 | Sin and Punishment.z64                                    |
| STARFOX64            | NFXE   | USA           | NTSC  | 6101 |        96 | Big-endian   | BA780BA0 | Star Fox 64 (USA) (Rev A).z64                             |
| Rogue Squadron       | NRSE   | USA           | NTSC  | 6102 |       128 | Big-endian   | 66A24BEC | Star Wars - Rogue Squadron (USA) (Rev A).z64              |
| SUPER MARIO 64       | NSME   | USA           | NTSC  | 6102 |        64 | Big-endian   | 635A2BFF | Super Mario 64 (USA).z64                                  |
| SUPERMARIO64         | NSMJ   | Japan         | NTSC  | 6102 |        64 | Byte-swapped | D6FBA4A8 | Super Mario 64 - Shindou Edition (J) [!].n64              |
| SUPERMARIO64         | NSMJ   | Japan         | NTSC  | 6102 |        64 | Big-endian   | D6FBA4A8 | Super Mario 64 - Shindou Edition (J) [!].z64              |
| SMASH BROTHERS       | NALE   | USA           | NTSC  | 6103 |       128 | Byte-swapped | 916B8B5B | Super Smash Bros. (USA).n64                               |
| SMASH BROTHERS       | NALE   | USA           | NTSC  | 6103 |       128 | Big-endian   | 916B8B5B | Super Smash Bros. (USA).z64                               |
| TSUMI TO BATSU       | NGUJ   | Japan         | NTSC  | 6102 |       256 | Byte-swapped | B6BC0FB0 | Tsumi to Batsu - Hoshi no Keishousha (Japan).n64          |
| TSUMI TO BATSU       | NGUJ   | Japan         | NTSC  | 6102 |       256 | Big-endian   | B6BC0FB0 | Tsumi to Batsu - Hoshi no Keishousha (Japan).z64          |
| WAVE RACE 64         | NWRE   | USA           | NTSC  | 6102 |        64 | Big-endian   | 492F4B61 | Wave Race 64 (USA) (Rev A).z64                            |
+----------------------+--------+---------------+-------+------+-----------+--------------+----------+-----------------------------------------------------------+

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
  Checksums:
    MD5:     00e2920665f2329b95797a7eaabc2390
    SHA1:    4cbadd3c4e0729dec46af64ad018050eada4f47a
    CRC 1:   30C7AC50
    CRC 2:   7704072D

ROM:
  ID:        NFUE
  Title:     CONKER BFD
  Media:     Cartridge
  Region:    USA
  Video:     NTSC
  Version:   1.0
  CIC:       6105
  CRC 1:     30C7AC50
  CRC 2:     7704072D
```

Same ROM but with `--output json`

```json
{
  "crc1": "30C7AC50",
  "crc2": "7704072D",
  "image_name": "CONKER BFD",
  "media_format": {
    "code": "N",
    "description": "Cartridge"
  },
  "cartridge_id": "FU",
  "region": {
    "id": "E",
    "short_name": "USA",
    "description": "USA",
    "video_system": "NTSC"
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
    "sha1": "4cbadd3c4e0729dec46af64ad018050eada4f47a",
    "crc1": "30C7AC50",
    "crc2": "7704072D"
  }
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
SHA-1 MATCH  581297B9D5C3A4C33169AE0AAE218C742CD9CBCF "Tsumi to Batsu - Hoshi no Keishousha (Japan).z64"
```

#### `--rename-validated`

With this flag, after a ROM's hash has been validated against the datfile, the ROM file will
be renamed to match what's in the datfile, if it doesn't already.

```
$ rom64 validate --rename-validated sm64.z64
Found 1 datfile entries for ROM serial 'NSME'
SHA-1 MATCH  9BEF1128717F958171A4AFAC3ED78EE2BB4E86CE "Super Mario 64 (USA).z64"
Renaming "sm64.z64" => "Super Mario 64 (USA).z64"
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
