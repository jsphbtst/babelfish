package data

import (
	"os"
	"path/filepath"
)

func InitializeDirFiles(rootDir string) error {
	err := os.MkdirAll(rootDir, 0755)
	if err != nil {
		return err
	}

	// ~/.babelfish/breakdowns.json
	breakdownsFilepath := filepath.Join(rootDir, "breakdowns.json")
	if !FileExists(breakdownsFilepath) {
		file, err := os.Create(breakdownsFilepath)
		if err != nil {
			return err
		}

		defer file.Close()
		content := "{ \"data\": [] }"
		file.Write([]byte(content))
	}

	// ~/.babelfish/configs.json
	configsFilepath := filepath.Join(rootDir, "configs.json")
	if !FileExists(configsFilepath) {
		file, err := os.Create(configsFilepath)
		if err != nil {
			return err
		}

		defer file.Close()

		content := "{ \"defaults\": { \"targetLanguage\": \"spanish\", \"stream\": false } }"
		file.Write([]byte(content))
	}

	// ~/.babelfish/history.json
	historyFilepath := filepath.Join(rootDir, "history.json")
	if !FileExists(historyFilepath) {
		file, err := os.Create(historyFilepath)
		if err != nil {
			return err
		}

		defer file.Close()
		content := "{ \"data\": [] }"
		file.Write([]byte(content))
	}

	// ~/.babelfish/openai-access-key
	accessKeyFilepath := filepath.Join(rootDir, "openai-access-key")
	if !FileExists(accessKeyFilepath) {
		file, err := os.Create(accessKeyFilepath)
		if err != nil {
			return err
		}

		defer file.Close()
		file.WriteString("")
	}

	return nil
}
