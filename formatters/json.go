package formatters

import (
	"encoding/json"
	"fmt"
)

func PrintJson(data interface{}) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = fmt.Println(string(bytes))
	return err
}
