package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/jsphbtst/babelfish/cmd"
	"github.com/jsphbtst/babelfish/pkg/checkers"
	"github.com/jsphbtst/babelfish/pkg/files"
	"github.com/jsphbtst/babelfish/pkg/types"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	rootDir := filepath.Join(home, ".babelfish")
	err = os.MkdirAll(rootDir, 0755)
	if err != nil {
		panic(err)
	}

	cmd.InitRootDir(rootDir)

	envPath := filepath.Join(rootDir, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Printf("Error loading .env file: %s\n", err.Error())
	}

	envkeys := []string{"OPENAI_API_KEY"}
	checkers.CheckEnv(envkeys, true)

	var configs types.Configs
	err = files.CreateOrParse(
		rootDir,
		"configs.json",
		[]byte("{ \"defaults\": { \"targetLanguage\": \"spanish\", \"stream\": false } }"),
		&configs,
	)

	if err != nil {
		panic(err)
	}

	cmd.InitConfigs(&configs)

	// History File
	var history types.HistoryJson
	err = files.CreateOrParse(
		rootDir,
		"history.json",
		[]byte("{ \"data\": [] }"),
		&history,
	)

	if err != nil {
		panic(err)
	}

	cmd.InitHistory(&history)

	// Breakdowns File
	var breakdowns types.BreakdownJson
	err = files.CreateOrParse(
		rootDir,
		"breakdowns.json",
		[]byte("{ \"data\": [] }"),
		&breakdowns,
	)

	if err != nil {
		panic(err)
	}

	cmd.InitBreakdowns(&breakdowns)

	// Begin CLI
	cmd.Execute()
}
