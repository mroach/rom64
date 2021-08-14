package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mroach/n64-go/rom"
)

func main() {
	args := os.Args[1:]

	filePath := args[0]
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	info, err := rom.ParseFile(f)
	if err != nil {
		panic(err)
	}

	jsonInfo, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonInfo)
}
