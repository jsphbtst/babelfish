package files

import (
	"encoding/json"
	"io"
	"os"
)

func Parse(fullpath string, structure interface{}) error {
	jsonFile, err := os.Open(fullpath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	configsJsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(configsJsonData, &structure)
	if err != nil {
		return err
	}

	return nil
}
