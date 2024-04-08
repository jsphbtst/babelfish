package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jsphbtst/babelfish/cmd"
	"github.com/jsphbtst/babelfish/pkg/checkers"
	"github.com/jsphbtst/babelfish/pkg/types"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %s\n", err.Error())
	}

	envkeys := []string{"OPENAI_API_KEY"}
	checkers.CheckEnv(envkeys, true)
	checkers.CheckInternet()

	// History File
	if _, err := os.Stat("history.json"); err != nil {
		file, err := os.Create("history.json")
		if err != nil {
			log.Fatalf("Failed to create history file: %s\n", err.Error())
			os.Exit(1)
		}

		content := "{ \"data\": [] }"
		file.Write([]byte(content))
		defer file.Close()
	}

	historyJsonFile, err := os.Open("history.json")
	if err != nil {
		log.Fatalf("Failed to open history file: %s\n", err.Error())
		os.Exit(1)
	}

	historyJsonData, err := io.ReadAll(historyJsonFile)
	if err != nil {
		log.Fatalf("Failed to parse history file: %s\n", err.Error())
		os.Exit(1)
	}

	var history types.HistoryJson
	err = json.Unmarshal(historyJsonData, &history)
	if err != nil {
		log.Fatalf("Failed to unmarshal history file: %s\n", err.Error())
		os.Exit(1)
	}

	cmd.InitHistory(&history)

	// Breakdowns File
	if _, err := os.Stat("breakdowns.json"); err != nil {
		file, err := os.Create("breakdowns.json")
		if err != nil {
			log.Fatalf("Failed to create breakdowns file: %s\n", err.Error())
			os.Exit(1)
		}

		content := "{ \"data\": [] }"
		file.Write([]byte(content))
		defer file.Close()
	}

	breakdownsJsonFile, err := os.Open("breakdowns.json")
	if err != nil {
		log.Fatalf("Failed to open breakdowns file: %s\n", err.Error())
		os.Exit(1)
	}

	breakdownsJsonData, err := io.ReadAll(breakdownsJsonFile)
	if err != nil {
		log.Fatalf("Failed to parse breakdowns file: %s\n", err.Error())
		os.Exit(1)
	}

	var breakdowns types.BreakdownJson
	err = json.Unmarshal(breakdownsJsonData, &breakdowns)
	if err != nil {
		log.Fatalf("Failed to unmarshal breakdowns file: %s\n", err.Error())
		os.Exit(1)
	}

	cmd.InitBreakdowns(&breakdowns)

	// Begin CLI
	cmd.Execute()
}
