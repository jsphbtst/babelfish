package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/jsphbtst/babelfish/cmd"
	"github.com/jsphbtst/babelfish/pkg/data"
	"github.com/jsphbtst/babelfish/pkg/files"
	"github.com/jsphbtst/babelfish/pkg/types"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	rootDir := filepath.Join(homeDir, ".babelfish")
	err = os.MkdirAll(rootDir, 0755)
	if err != nil {
		panic(err)
	}

	cmd.InitRootDir(rootDir)

	data.InitializeDirFiles(rootDir)

	// TODO: refactor this part later
	keyFile, err := os.OpenFile(
		filepath.Join(rootDir, "openai-access-key"),
		os.O_RDWR|os.O_CREATE,
		0666,
	)
	if err != nil {
		panic(err)
	}
	defer keyFile.Close()

	openAiBytes, err := io.ReadAll(keyFile)
	if err != nil {
		panic(err)
	}

	openAiApiKey := string(openAiBytes)
	cmd.InitOpenAiKey(openAiApiKey)

	// Configs file
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
