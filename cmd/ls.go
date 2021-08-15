package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mroach/n64-go/rom"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var Output string

func init() {
	lsCmd.Flags().StringVarP(&Output, "output", "o", "", "Output format (csv, json, text, table)")
	rootCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List Nintendo 64 ROM or ROMs",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		fi, err := os.Stat(path)
		if err != nil {
			return err
		}

		if fi.IsDir() {
			files, err := findRomsInPath(path)
			if err != nil {
				return err
			}

			fileInfos := make([]rom.RomFile, 0)
			for _, romPath := range files {
				info, err := rom.FromPath(romPath)
				if err != nil {
					return err
				}

				fileInfos = append(fileInfos, info)
			}

			if Output == "" {
				Output = "table"
			}

			printAll(fileInfos)
		} else {
			info, err := rom.FromPath(path)
			if err != nil {
				return err
			}

			if Output == "" {
				Output = "text"
			}

			printOne(info)
		}

		return nil
	},
}

func printOne(item rom.RomFile) {
	switch Output {
	case "json":
		printJson(item)
	case "csv":
		items := []rom.RomFile{item}
		records := romsToCsvRecords(items)
		printCsv(records, defaultCsvHeaders)
	case "table":
		items := []rom.RomFile{item}
		records := romsToCsvRecords(items)
		printTable(records, defaultTableHeaders)
	case "text":
		printText(item)
	default:
		fmt.Println("Invalid output format", Output)
	}
}

func printAll(items []rom.RomFile) {
	switch Output {
	case "json":
		printJson(items)
	case "csv":
		records := romsToCsvRecords(items)
		printCsv(records, defaultCsvHeaders)
	case "table":
		records := romsToCsvRecords(items)
		printTable(records, defaultTableHeaders)
	case "text":
		for _, item := range items {
			printText(item)
			fmt.Println("")
		}
	default:
		fmt.Println("Invalid output format", Output)
	}
}

var romExtensions = []string{"n64", "v64", "z64", "N64", "V64", "Z64"}

func findRomsInPath(path string) ([]string, error) {
	searchPattern := filepath.Join(path, "*")
	files, err := filepath.Glob(searchPattern)
	if err != nil {
		return []string{}, err
	}

	romFiles := make([]string, 0)
	for _, file := range files {
		if isProbablyRom(file) {
			romFiles = append(romFiles, file)
		}
	}

	return romFiles, nil
}

func isProbablyRom(path string) bool {
	ext := filepath.Ext(path)
	for _, v := range romExtensions {
		if "."+v == ext {
			return true
		}
	}
	return false
}

var defaultCsvHeaders = []string{
	"file_name", "file_format", "file_size",
	"image_name", "media_format", "cartridge_id",
	"region_name", "cic", "crc1", "crc2",
}

var defaultTableHeaders = []string{
	"File Name", "Format", "Size", "Title",
	"Type", "ID", "Region", "CIC", "CRC1", "CRC2",
}

func romsToCsvRecords(infos []rom.RomFile) [][]string {
	out := make([][]string, 0)
	for _, info := range infos {
		out = append(out, infoToRecord(info))
	}
	return out
}

// TODO: Take a list of named fields to customise this
func infoToRecord(info rom.RomFile) []string {
	return []string{
		info.File.Name,
		info.File.Format,
		fmt.Sprintf("%d", info.File.Size),
		info.ImageName,
		info.MediaFormat.Code,
		info.CartridgeId,
		info.Region.Description,
		info.CIC,
		info.CRC1,
		info.CRC2,
	}
}

var textFormat = `File:
  Name:    %s
  Size:    %d MB
  Format:  %s

Title:     %s
ROM ID:    %s
Media:     %s (%s)
Version:   1.%d
Region:    %s (%s)
CIC:       %s
CRC 1:     %s
CRC 2:     %s`

func printText(info rom.RomFile) {
	fmt.Printf(textFormat,
		info.File.Name,
		info.File.Size,
		info.File.Format,
		info.ImageName,
		info.CartridgeId,
		info.MediaFormat.Code, info.MediaFormat.Description,
		info.Version,
		info.Region.Code, info.Region.Description,
		info.CIC,
		info.CRC1,
		info.CRC2,
	)
	fmt.Println("")
}

func printTable(records [][]string, headers []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(headers)
	table.AppendBulk(records)
	table.Render()
}

func printCsv(records [][]string, headers []string) {
	w := csv.NewWriter(os.Stdout)
	if err := w.Write(headers); err != nil {
		log.Fatalln("error writing csv headers:", err)
	}
	if err := w.WriteAll(records); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func printJson(data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalln("error generating json:", err)
	}
	fmt.Println(string(bytes))
}
