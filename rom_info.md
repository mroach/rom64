# ROM Information

A collection of my notes and gathered information about N64 ROM files and headers.


ROM ID / Serial
--------------------------------------------------------------------------------
Each ROM has a four-character ID, such as `NSME` for *Super Mario 64*.
Using that as an example:

```
N SM E
│ │┘ │
│ │  └ Region
│ └ Software ID
└ Media Type
```

### Media Type - `0x38 - 0x3B` - 4 bytes, NUL-padded

By far the most common is `N` for a standard cartridge.

There are a handful of games with `C` such as *Legend of Zelda: Ocarina of Time*,
*Mario Party*, and *F-Zero X*.

* `N` - Cartridge
* `D` - 64DD Disk
* `C` - Cartridge for expandable game
* `E` - 64DD Expansion
* `Z` - Aleck64 Cartridge

### Software ID - `0x3C - 0x3D` - 2 bytes

Each game gets its own two-character alphanumeric ID that is the same across all regions.

Examples:

* `SM` - Super Mario 64
* `FU` - Conker's Bad Fur Day
* `B5` - Biohazard 2

### ROM Region IDs - `0x3E` - 1 byte

The final character of the ROM ID is the region.

By far the most common are:
*E* for the US, *J* for Japan, and *P* for PAL regions.

* `7` - Beta
* `A` - Japan + US
* `B` - Brazil
* `C` - China
* `D` - Germany
* `E` - US
* `F` - France
* `G` - Gateway 64 (NTSC)
* `H` - Netherlands
* `I` - Italy
* `J` - Japan
* `K` - South Korea
* `L` - Gateway 64 (PAL)
* `N` - Canada
* `P` - PAL
* `S` - Spain
* `U` - Australia
* `W` - Scandinavia
* `X` - PAL (Uncommon)
* `Y` - PAL (Uncommon)


Package Regions
--------------------------------------------------------------------------------

Package regions do not appear in ROM data. These only appear on the cartridge itself
and on the box, perhaps to denote the target market such as box art and language of the manuals.

There's not always a clear mapping to/from ROM region IDs.

| Code   | Description                    |
| ------ | ------------------------------ |
| `ASI`  | Singapore, Malaysia, Indonesia |
| `ASM`  | Asia (excluding Japan)         |
| `AUS`  | Australia                      |
| `ESP`  | Spain                          |
| `EUR`  | Europe                         |
| `EUU`  | Europe                         |
| `FAH`  | France and Netherlands         |
| `FRA`  | France                         |
| `FRG`  | France and Germany             |
| `HKG`  | Hong Kong                      |
| `HOL`  | Netherlands                    |
| `ITA`  | Italy                          |
| `JPN`  | Japan                          |
| `KOR`  | South Korea                    |
| `MSA`  | Mexico                         |
| `NOE`  | Germany                        |
| `ROC`  | Taiwan                         |
| `SCN`  | Scandinavia                    |
| `UKV`  | United Kingdom                 |
| `USA`  | United States                  |

Brazilian cartridges don't have package region codes, though it's common to use `BRA`
when referring to a Brazil-specific game or cartridge.


Software Title - `0x20` - 20 bytes, NUL-padded
--------------------------------------------------------------------------------

Each ROM includes up to 20 characters for the software/game title.
Some examples:

* `CONKER BFD`
* `Diddy Kong Racing`
* `F-ZERO X`
* `JET FORCE GEMINI`
* `LEGORacers`
* `THE LEGEND OF ZELDA`
* `MarioParty2`
* `EVANGELION`
* `PAPER MARIO`
* `Perfect Dark`
* `TSUMI TO BATSU`


CRC1 - `0x10 - 0x13` - 4 bytes
--------------------------------------------------------------------------------

Also known as **CRC HI** in the EverDrive 64.

Calculated against `0x100000` bytes starting at `0x1000`

See [http://n64dev.org/n64crc.html](http://n64dev.org/n64crc.html)


CRC2 - `0x14 - 0x18` - 4 bytes
--------------------------------------------------------------------------------

Also known as **CRC LO** in the EverDrive 64.

See [http://n64dev.org/n64crc.html](http://n64dev.org/n64crc.html)


References
--------------------------------------------------------------------------------
* https://github.com/n64dev/cen64/blob/master/device/cart_db.c
* http://en64.shoutwiki.com/wiki/ROM#Cartridge_ROM_Header
