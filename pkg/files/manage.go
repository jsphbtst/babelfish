package files

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

func CreateOrParse(rootDir string, filename string, defaultData []byte, structure interface{}) error {
	filePath := filepath.Join(rootDir, filename)
	jsonFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	info, err := jsonFile.Stat()
	if err != nil {
		return err
	}

	if info.Size() == 0 {
		if _, err := jsonFile.Write(defaultData); err != nil {
			return err
		}

		if _, err := jsonFile.Seek(0, 0); err != nil {
			return err
		}
	}

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
