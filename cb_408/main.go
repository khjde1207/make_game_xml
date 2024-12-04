package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var fns = template.FuncMap{
	"isLast": func(index int, len int) bool {
		return index+1 != len
	},
}

// export GOOS=windows
// export GOARCH=amd64
// go build -o cb_408maker.exe && zip cb_408maker.zip cb_408maker.exe

func main() {

	t := template.New("").Funcs(fns)
	t1, _ := t.Parse(string(rootData))

	for key, coll := range collectionMap {
		path := fmt.Sprint("./", key)
		_, err := os.Stat(path)
		if err != nil {
			os.Mkdir(path, os.ModePerm)
		}

		dirList, _ := os.ReadDir(path)
		roms := []RomData{}
		for _, gameFile := range dirList {
			if gameFile.Name() == "metadata.json" || gameFile.IsDir() {
				continue
			}

			fullName := gameFile.Name()
			arr := strings.Split(fullName, ".")
			name := strings.Join(arr[0:len(arr)-1], "")
			rom := RomData{
				FullName: fullName,
				Name:     name,
			}
			roms = append(roms, rom)

		}
		coll.Roms = roms

		ofile, _ := os.Create(fmt.Sprint(path, "/metadata.json"))

		t1.Execute(ofile, coll)
	}

}

var collectionMap = map[string]CollectionData{
	"ARCADE": CollectionData{
		Name:    "ARCADE",
		Dirname: "ARCADE",
		Core:    "RetroArch FB Alpha 2012",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FB Alpha 2012",
				DisplayName: "Arcade (FB Alpha 2012)",
				Core:        "fbalpha2012_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME 2010 (0.139)",
				DisplayName: "Arcade (MAME 2010)",
				Core:        "mame2010_libretro",
			},
		},
	},
	"ATARI2600": CollectionData{
		Name:    "A2600",
		Dirname: "A2600\\nATARI 2600\\nATARI2600",
		Core:    "RetroArch Stella 2014",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Stella",
				DisplayName: "Atari - 2600 (Stella)",
				Core:        "stella_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Stella 2014",
				DisplayName: "Atari - 2600 (Stella 2014)",
				Core:        "stella2014_libretro",
			},
		},
	},
	"ATARI5200": CollectionData{
		Name:    "A5200",
		Dirname: "A5200\\nATARI 5200\\nATARI5200",
		Core:    "RetroArch a5200",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Atari800",
				DisplayName: "Atari - 5200 (Atari800)",
				Core:        "atari800_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch a5200",
				DisplayName: "Atari - 5200 (a5200)",
				Core:        "a5200_libretro",
			},
		},
	},
	"ATARI7800": CollectionData{
		Name:    "A7800",
		Dirname: "A7800\\nATARI 7800\\nATARI7800",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch ProSystem",
				DisplayName: "Atari - 7800 (ProSystem)",
				Core:        "prosystem_libretro",
			},
		},
	},
	"ATARILYNX": CollectionData{
		Name:    "LYNX",
		Dirname: "LYNX\\nATARI LYNX\\nATARILYNX",
		Core:    "RetroArch Beetle Lynx",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Handy",
				DisplayName: "Atari - Lynx (Handy)",
				Core:        "handy_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Beetle Lynx",
				DisplayName: "Atari - Lynx (Beetle Lynx)",
				Core:        "mednafen_lynx_libretro",
			},
		},
	},
	"CPS1": CollectionData{
		Name:    "CPS1",
		Dirname: "CPS1",
		Core:    "RetroArch FB Alpha 2012",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Alpha",
				DisplayName: "Arcade (FinalBurn Alpha)",
				Core:        "fbalpha_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FB Alpha 2012",
				DisplayName: "Arcade (FB Alpha 2012)",
				Core:        "fbalpha2012_libretro",
			},
		},
	},
	"CPS2": CollectionData{
		Name:    "CPS2",
		Dirname: "CPS2",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
		},
	},
	"CPS3": CollectionData{
		Name:    "CPS3",
		Dirname: "CPS3",
		Core:    "RetroArch FinalBurn Alpha",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Alpha",
				DisplayName: "Arcade (FinalBurn Alpha)",
				Core:        "fbalpha_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FB Alpha 2012",
				DisplayName: "Arcade (FB Alpha 2012)",
				Core:        "fbalpha2012_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME (Git)",
				DisplayName: "Arcade (MAME - Current)",
				Core:        "mamearcade_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME 2010 (0.139)",
				DisplayName: "Arcade (MAME 2010)",
				Core:        "mame2010_libretro",
			},
		},
	},
	"DC": CollectionData{
		Name:    "DC",
		Dirname: "DC\\nDC HACK\\nDREAMCAST",
		Core:    "Flycast",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Flycast",
				DisplayName: "Sega - Dreamcast/Naomi (Flycast)",
				Core:        "flycast_libretro",
			},
			EmulatorData{
				CoreName:    "Flycast",
				DisplayName: "Flycast Emulator - Sega DC/NAOMI",
				Core:        "",
			},
		},
	},
	"FBNEO": CollectionData{
		Name:    "FBNEO",
		Dirname: "FBNEO\\nFBNEO ARCADE\\nFBNEOARCADE",
		Core:    "RetroArch FinalBurn Alpha",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Alpha",
				DisplayName: "Arcade (FinalBurn Alpha)",
				Core:        "fbalpha_libretro",
			},
		},
	},
	"FBNEO ACT": CollectionData{
		Name:    "AACT",
		Dirname: "AACT\\nFBNEO ACT\\nFBNEO ACT V\\nFBNEO ACT HACK\\nMAME ACT",
		Core:    "RetroArch FinalBurn Alpha",

		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Alpha",
				DisplayName: "Arcade (FinalBurn Alpha)",
				Core:        "fbalpha_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME (Git)",
				DisplayName: "Arcade (MAME - Current)",
				Core:        "mamearcade_libretro",
			},
		},
	},
	"FBNEO ETC": CollectionData{
		Name:    "AETC",
		Dirname: "AETC\\nFBNEO ETC\\nFBNEO ETC V\\nMAME ETC",
		Core:    "RetroArch MAME (Git)",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME (Git)",
				DisplayName: "Arcade (MAME - Current)",
				Core:        "mamearcade_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME 2003-Plus",
				DisplayName: "Arcade (MAME 2003-Plus)",
				Core:        "mame2003_plus_libretro",
			},
		},
	},
	"FBNEO FLY": CollectionData{
		Name:       "AFLY",
		Dirname:    "AFLY\\nFBNEO FLY\\nFBNEO FLY V\\nMAME FLY\\nMAME FLY V",
		Extensions: "7z, zip",
		Core:       "RetroArch MAME (Git)",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FB Alpha 2012",
				DisplayName: "Arcade (FB Alpha 2012)",
				Core:        "fbalpha2012_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME (Git)",
				DisplayName: "Arcade (MAME - Current)",
				Core:        "mamearcade_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME 2003-Plus",
				DisplayName: "Arcade (MAME 2003-Plus)",
				Core:        "mame2003_plus_libretro",
			},
		},
	},
	"FBNEO FTG": CollectionData{
		Name:       "AFTG",
		Dirname:    "AFTG\\nFBNEO FTG\\nFBNEO FTG HACK\\nMAME FTG\\nMAME FTG HACK",
		Extensions: "7z, zip",
		Core:       "RetroArch FB Alpha 2012",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Alpha",
				DisplayName: "Arcade (FinalBurn Alpha)",
				Core:        "fbalpha_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FB Alpha 2012",
				DisplayName: "Arcade (FB Alpha 2012)",
				Core:        "fbalpha2012_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME (Git)",
				DisplayName: "Arcade (MAME - Current)",
				Core:        "mamearcade_libretro",
			},
		},
	},
	"FBNEO STG": CollectionData{
		Name:       "ASTG",
		Dirname:    "ASTG\\nFBNEO STG\\nFBNEO STG V\\nFBNEO STG HACK\\nMAME STG\\nMAME STG V",
		Extensions: "7z, zip",
		Core:       "RetroArch MAME (Git)",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FB Alpha 2012",
				DisplayName: "Arcade (FB Alpha 2012)",
				Core:        "fbalpha2012_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME 2003 (0.78)",
				DisplayName: "Arcade (MAME 2003)",
				Core:        "mame2003_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME (Git)",
				DisplayName: "Arcade (MAME - Current)",
				Core:        "mamearcade_libretro",
			},
		},
	},
	"FC": CollectionData{
		Name:       "FC",
		Dirname:    "FC",
		Extensions: "nes, 7z, zip",
		Core:       "RetroArch FCEUmm",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Nestopia",
				DisplayName: "Nintendo - NES / Famicom (Nestopia UE)",
				Core:        "nestopia_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FCEUmm",
				DisplayName: "Nintendo - NES / Famicom (FCEUmm)",
				Core:        "fceumm_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch QuickNES",
				DisplayName: "Nintendo - NES / Famicom (QuickNES)",
				Core:        "quicknes_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Mesen",
				DisplayName: "Nintendo - NES / Famicom (Mesen)",
				Core:        "mesen_libretro",
			},
		},
	},
	"GAMEGEAR": CollectionData{
		Name:    "GG",
		Dirname: "GG\\nGAME GEAR\\nGAMEGEAR",
		Core:    "RetroArch PicoDrive",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch SMS Plus GX",
				DisplayName: "Sega - MS/GG (SMS Plus GX)",
				Core:        "smsplus_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Gearsystem",
				DisplayName: "Sega - MS/GG/SG-1000 (Gearsystem)",
				Core:        "gearsystem_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX)",
				Core:        "genesis_plus_gx_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX Wide",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX Wide)",
				Core:        "genesis_plus_gx_wide_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch PicoDrive",
				DisplayName: "Sega - MS/GG/MD/CD/32X (PicoDrive)",
				Core:        "picodrive_libretro",
			},
		},
	},
	"GB": CollectionData{
		Name:    "GB",
		Dirname: "GB",
		Core:    "RetroArch mGBA",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Gambatte",
				DisplayName: "Nintendo - Game Boy / Color (Gambatte)",
				Core:        "gambatte_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch SameBoy",
				DisplayName: "Nintendo - Game Boy / Color (SameBoy)",
				Core:        "sameboy_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch TGB Dual",
				DisplayName: "Nintendo - Game Boy / Color (TGB Dual)",
				Core:        "tgbdual_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Gearboy",
				DisplayName: "Nintendo - Game Boy / Color (Gearboy)",
				Core:        "gearboy_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch mGBA",
				DisplayName: "Nintendo - Game Boy Advance (mGBA)",
				Core:        "mgba_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch VBA-M",
				DisplayName: "Nintendo - Game Boy Advance (VBA-M)",
				Core:        "vbam_libretro",
			},
		},
	},
	"GBA": CollectionData{
		Name:    "GBA",
		Dirname: "GBA",
		Core:    "RetroArch mGBA",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch gpSP",
				DisplayName: "Nintendo - Game Boy Advance (gpSP)",
				Core:        "gpsp_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch mGBA",
				DisplayName: "Nintendo - Game Boy Advance (mGBA)",
				Core:        "mgba_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch VBA Next",
				DisplayName: "Nintendo - Game Boy Advance (VBA Next)",
				Core:        "vba_next_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch VBA-M",
				DisplayName: "Nintendo - Game Boy Advance (VBA-M)",
				Core:        "vbam_libretro",
			},
		},
	},
	"GBC": CollectionData{
		Name:    "GBC",
		Dirname: "GBC",
		Core:    "RetroArch Mesen-S",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Gambatte",
				DisplayName: "Nintendo - Game Boy / Color (Gambatte)",
				Core:        "gambatte_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch SameBoy",
				DisplayName: "Nintendo - Game Boy / Color (SameBoy)",
				Core:        "sameboy_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch TGB Dual",
				DisplayName: "Nintendo - Game Boy / Color (TGB Dual)",
				Core:        "tgbdual_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Gearboy",
				DisplayName: "Nintendo - Game Boy / Color (Gearboy)",
				Core:        "gearboy_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch mGBA",
				DisplayName: "Nintendo - Game Boy Advance (mGBA)",
				Core:        "mgba_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch VBA-M",
				DisplayName: "Nintendo - Game Boy Advance (VBA-M)",
				Core:        "vbam_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Mesen-S",
				DisplayName: "Nintendo - SNES / SFC / Game Boy / Color (Mesen-S)",
				Core:        "mesen-s_libretro",
			},
		},
	},
	"GENESIS": CollectionData{
		Name:    "GENESIS",
		Dirname: "GENESIS",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch BlastEm",
				DisplayName: "Sega - Mega Drive - Genesis (BlastEm)",
				Core:        "blastem_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX)",
				Core:        "genesis_plus_gx_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX Wide",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX Wide)",
				Core:        "genesis_plus_gx_wide_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch PicoDrive",
				DisplayName: "Sega - MS/GG/MD/CD/32X (PicoDrive)",
				Core:        "picodrive_libretro",
			},
		},
	},
	"GG": CollectionData{
		Name:    "GG",
		Dirname: "GG\\nGAME GEAR\\nGAMEGEAR",
		Core:    "RetroArch PicoDrive",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch SMS Plus GX",
				DisplayName: "Sega - MS/GG (SMS Plus GX)",
				Core:        "smsplus_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Gearsystem",
				DisplayName: "Sega - MS/GG/SG-1000 (Gearsystem)",
				Core:        "gearsystem_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX)",
				Core:        "genesis_plus_gx_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX Wide",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX Wide)",
				Core:        "genesis_plus_gx_wide_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch PicoDrive",
				DisplayName: "Sega - MS/GG/MD/CD/32X (PicoDrive)",
				Core:        "picodrive_libretro",
			},
		},
	},
	"LYNX": CollectionData{
		Name:    "LYNX",
		Dirname: "LYNX\\nATARI LYNX\\nATARILYNX",
		Core:    "RetroArch Beetle Lynx",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Handy",
				DisplayName: "Atari - Lynx (Handy)",
				Core:        "handy_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Beetle Lynx",
				DisplayName: "Atari - Lynx (Beetle Lynx)",
				Core:        "mednafen_lynx_libretro",
			},
		},
	},
	"MAME": CollectionData{
		Name:    "MAME",
		Dirname: "MAME\\nMAME ARCADE\\nMAMEARCADE",
		Core:    "RetroArch MAME 2010 (0.139)",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Neo",
				DisplayName: "Arcade (FinalBurn Neo)",
				Core:        "fbneo_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch FinalBurn Alpha",
				DisplayName: "Arcade (FinalBurn Alpha)",
				Core:        "fbalpha_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME (Git)",
				DisplayName: "Arcade (MAME - Current)",
				Core:        "mamearcade_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME 2000 (0.37b5)",
				DisplayName: "Arcade (MAME 2000)",
				Core:        "mame2000_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch MAME 2010 (0.139)",
				DisplayName: "Arcade (MAME 2010)",
				Core:        "mame2010_libretro",
			},
		},
	},
	"MD": CollectionData{
		Name:    "MD",
		Dirname: "MD\\nMEGA DRIVE\\nMEGADRIVE",
		Core:    "RetroArch PicoDrive",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch BlastEm",
				DisplayName: "Sega - Mega Drive - Genesis (BlastEm)",
				Core:        "blastem_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX)",
				Core:        "genesis_plus_gx_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX Wide",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX Wide)",
				Core:        "genesis_plus_gx_wide_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch PicoDrive",
				DisplayName: "Sega - MS/GG/MD/CD/32X (PicoDrive)",
				Core:        "picodrive_libretro",
			},
		},
	},
	"MEGADRIVE": CollectionData{
		Name:    "MD",
		Dirname: "MD\\nMEGA DRIVE\\nMEGADRIVE",
		Core:    "RetroArch PicoDrive",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch BlastEm",
				DisplayName: "Sega - Mega Drive - Genesis (BlastEm)",
				Core:        "blastem_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX)",
				Core:        "genesis_plus_gx_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX Wide",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX Wide)",
				Core:        "genesis_plus_gx_wide_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch PicoDrive",
				DisplayName: "Sega - MS/GG/MD/CD/32X (PicoDrive)",
				Core:        "picodrive_libretro",
			},
		},
	},
	"N64": CollectionData{
		Name:    "N64",
		Dirname: "N64",
		Core:    "Mupen64Plus",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Mupen64Plus-Next",
				DisplayName: "Nintendo - Nintendo 64 (Mupen64Plus-Next)",
				Core:        "mupen64plus_next_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch ParaLLEl N64",
				DisplayName: "Nintendo - Nintendo 64 (ParaLLEl N64)",
				Core:        "parallel_n64_libretro",
			},
			EmulatorData{
				CoreName:    "Mupen64Plus",
				DisplayName: "Mupen64Plus Emulator - Nintendo 64",
				Core:        "",
			},
		},
	},
	"NAOMI": CollectionData{
		Name:    "NAOMI",
		Dirname: "NAOMI",

		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Flycast",
				DisplayName: "Sega - Dreamcast/Naomi (Flycast)",
				Core:        "flycast_libretro",
			},
		},
	},
	"NES": CollectionData{
		Name:    "NES",
		Dirname: "NES",
		Core:    "RetroArch QuickNES",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch FCEUmm",
				DisplayName: "Nintendo - NES / Famicom (FCEUmm)",
				Core:        "fceumm_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Mesen",
				DisplayName: "Nintendo - NES / Famicom (Mesen)",
				Core:        "mesen_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Nestopia",
				DisplayName: "Nintendo - NES / Famicom (Nestopia UE)",
				Core:        "nestopia_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch QuickNES",
				DisplayName: "Nintendo - NES / Famicom (QuickNES)",
				Core:        "quicknes_libretro",
			},
		},
	},
	"NGPC": CollectionData{
		Name:    "NGPC",
		Dirname: "NGPC",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Beetle NeoPop",
				DisplayName: "SNK - Neo Geo Pocket / Color (Beetle NeoPop)",
				Core:        "mednafen_ngp_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch RACE",
				DisplayName: "SNK - Neo Geo Pocket / Color (RACE)",
				Core:        "race_libretro",
			},
		},
	},
	"PCE": CollectionData{
		Name:    "PCE",
		Dirname: "PCE\\nPCE NGINE\\nPCENGINE",
		Core:    "RetroArch Beetle SuperGrafx",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Beetle PCE Fast",
				DisplayName: "NEC - PC Engine / CD (Beetle PCE FAST)",
				Core:        "mednafen_pce_fast_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Beetle PCE",
				DisplayName: "NEC - PC Engine / SuperGrafx / CD (Beetle PCE)",
				Core:        "mednafen_pce_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Beetle SuperGrafx",
				DisplayName: "NEC - PC Engine SuperGrafx (Beetle SuperGrafx)",
				Core:        "mednafen_supergrafx_libretro",
			},
		},
	},
	"PCENGINE": CollectionData{
		Name:    "PCE",
		Dirname: "PCE\\nPCE NGINE\\nPCENGINE",
		Core:    "RetroArch Beetle SuperGrafx",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Beetle PCE Fast",
				DisplayName: "NEC - PC Engine / CD (Beetle PCE FAST)",
				Core:        "mednafen_pce_fast_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Beetle PCE",
				DisplayName: "NEC - PC Engine / SuperGrafx / CD (Beetle PCE)",
				Core:        "mednafen_pce_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Beetle SuperGrafx",
				DisplayName: "NEC - PC Engine SuperGrafx (Beetle SuperGrafx)",
				Core:        "mednafen_supergrafx_libretro",
			},
		},
	},
	"PS1": CollectionData{
		Name:    "PS1",
		Dirname: "PS1\\nPSX\\nPS1 HACK",
		Core:    "PCSX-ReARMed",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch SwanStation",
				DisplayName: "Sony - PlayStation (SwanStation)",
				Core:        "swanstation_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch PCSX-ReARMed",
				DisplayName: "Sony - PlayStation (PCSX ReARMed)",
				Core:        "pcsx_rearmed_libretro",
			},
		},
	},
	"PSP": CollectionData{
		Name:    "PSP",
		Dirname: "PSP",
		Core:    "PPSSPP",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch PPSSPP",
				DisplayName: "Sony - PlayStation Portable (PPSSPP)",
				Core:        "ppsspp_libretro",
			},
			EmulatorData{
				CoreName:    "PPSSPP",
				DisplayName: "PPSSPP Emulator - Sony PSP",
				Core:        "",
			},
		},
	},
	"SFC": CollectionData{
		Name:    "SFC",
		Dirname: "SFC",
		Core:    "RetroArch Stella",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Snes9x",
				DisplayName: "Nintendo - SNES / SFC (Snes9x - Current)",
				Core:        "snes9x_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Stella",
				DisplayName: "Atari - 2600 (Stella)",
				Core:        "stella_libretro",
			},
		},
	},
	"SMS": CollectionData{
		Name:    "SMS",
		Dirname: "SMS\\nMS\\nMASTER SYSTEM\\nMASTERSYSTEM",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Genesis Plus GX",
				DisplayName: "Sega - MS/GG/MD/CD (Genesis Plus GX)",
				Core:        "genesis_plus_gx_libretro",
			},
		},
	},
	"SNES": CollectionData{
		Name:    "SNES",
		Dirname: "SNES",
		Core:    "RetroArch Stella",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Snes9x",
				DisplayName: "Nintendo - SNES / SFC (Snes9x - Current)",
				Core:        "snes9x_libretro",
			},
			EmulatorData{
				CoreName:    "RetroArch Stella",
				DisplayName: "Atari - 2600 (Stella)",
				Core:        "stella_libretro",
			},
		},
	},
	"WS": CollectionData{
		Name:       "WS",
		Dirname:    "WS",
		Extensions: "ws, zip",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Beetle WonderSwan",
				DisplayName: "Bandai - WonderSwan/Color (Beetle Cygne)",
				Core:        "mednafen_wswan_libretro",
			},
		},
	},
	"WSC": CollectionData{
		Name:       "WSC",
		Dirname:    "WSC",
		Extensions: "ws, wsc, zip",
		Emulator: []EmulatorData{
			EmulatorData{
				CoreName:    "RetroArch Beetle WonderSwan",
				DisplayName: "Bandai - WonderSwan/Color (Beetle Cygne)",
				Core:        "mednafen_wswan_libretro",
			},
		},
	},
}

type CollectionData struct {
	Name       string
	Dirname    string
	Extensions string
	Emulator   []EmulatorData
	Roms       []RomData
	Core       string
}

type EmulatorData struct {
	CoreName    string
	DisplayName string
	Core        string
}

type RomData struct {
	FullName string
	Name     string
}

var rootData = `
  {
	"Collection": {
		"Name": "{{.Name}}",
		"Dirname": "{{.Dirname}}"
	},
	"Emulator": [
		{{$lenEmulator := len .Emulator}}{{range $idx, $var := .Emulator}}{
			"CoreName": "{{.CoreName}}",
			"DisplayName": "{{.DisplayName}}",
			"Core": "{{.Core}}"
		}{{if (isLast $idx $lenEmulator)}},{{end}}
		{{end}}
	],
	"Roms": [
		{{$rootCore := .Core}}
		{{$lenMyList := len .Roms}}{{range $index, $var := .Roms}}{
			"Id": "{{.Name}}",
			"Files": [
				"{{.FullName}}"
			],
			"Name": {
				"zh_CN": "{{.Name}}",
				"zh_TW": "{{.Name}}",
				"en_US": "{{.Name}}",
				"ko_KR": "{{.Name}}"
			},
			{{if ne $rootCore  ""}}"Core": "{{$rootCore}}",{{end}}
			"Flag": 1,
			"MD5": "",
			"Search": {}
		} {{if (isLast $index $lenMyList)}}, {{end}}
		{{end}}
  	]
  }
`
