package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jsphbtst/babelfish/pkg/checkers"
	"github.com/jsphbtst/babelfish/pkg/loader"
	"github.com/jsphbtst/babelfish/pkg/requests"
	"github.com/jsphbtst/babelfish/pkg/types"
	"github.com/spf13/cobra"
)

var breakdownCmd = &cobra.Command{
	Use:   "breakdown",
	Short: "Breakdown",
	Long: `Breakdown.

	For example:

	translate-cli breakdown -p "donde esta la biblioteca?"
`,
	Run: runBreakdownCmd,
}

func init() {
	rootCmd.AddCommand(breakdownCmd)
	breakdownCmd.Flags().StringP("phrase", "p", "", "The language phrase to be broken down")
}

func runBreakdownCmd(cmd *cobra.Command, args []string) {
	phrase, err := cmd.Flags().GetString("phrase")
	if err != nil {
		panic(err)
	}

	if !checkers.IsWithinCwsLimit(phrase) {
		fmt.Println("You've exceeded the currently supported max 180 CWS limit.")
		os.Exit(1)
	}

	lowercasedPhrase := strings.ToLower(phrase)

	for _, p := range globals.Explanations.Data {
		if strings.ToLower(p.Phrase) == lowercasedPhrase {
			fmt.Printf("Breakdown: %s\n", p.Breakdown)
			return
		}
	}

	checkers.CheckInternet()

	prompt := fmt.Sprintf(
		"Can you break down the phrase \"%s\"? I'm trying to learn this language and I need a breakdown. In this scenario, act as if you're a robot who isn't familiar with manners, therefore, you only provide the explanation directly.",
		phrase,
	)

	end := loader.PrintProgress("Breakdown in progress...")
	defer end()

	result, err := requests.RequestGpt4Translation(prompt)
	if err != nil {
		panic(err)
	}

	end()

	for _, choice := range result.Choices {
		answer := choice.Message.Content

		globals.Explanations.Data = append(
			globals.Explanations.Data,
			types.BreakdownRecord{
				Phrase:    phrase,
				Breakdown: answer,
				CreatedAt: time.Now().UTC(),
			},
		)
		fmt.Printf("Breakdown: %s\n", answer)
	}

	file, err := json.MarshalIndent(globals.Explanations, "", " ")
	if err != nil {
		log.Fatalf("Failed to marshal json: %s\n", err.Error())
		return
	}

	filePath := filepath.Join(globals.RootDir, "breakdowns.json")
	err = os.WriteFile(filePath, file, 0644)
	if err != nil {
		log.Fatalf("Failed to update history file: %s\n", err.Error())
		return
	}
}
