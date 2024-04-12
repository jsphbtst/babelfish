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

var convCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate text to your target language",
	Long: `Command to translate any word, sentence, or phrase into your target language. The
	default is Castellano.

	For example:

	babelfish translate "this is a sample text"
	babelfish translate "this is a sample text" -t "Spanish"
`,
	Run: generateTranslation,
}

func init() {
	rootCmd.AddCommand(convCmd)
	convCmd.Flags().StringP(
		"target",
		"t",
		"", // TODO: fix this later
		"The target language for translation",
	)
}

func generateTranslation(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Incorrect usage")
		os.Exit(1)
	}

	phrase := args[0]
	if !checkers.IsWithinCwsLimit(phrase) {
		fmt.Println("You've exceeded the currently supported max 180 CWS limit.")
		os.Exit(1)
	}

	targetLang, err := cmd.Flags().GetString("target")
	if err != nil {
		panic(err)
	}

	if targetLang == "" {
		targetLang = globals.Configs.Defaults.TargetLanguage
	}

	if !checkers.IsSupportedLanguage(strings.ToLower(targetLang)) {
		fmt.Println("Currently an unsupported language")
		os.Exit(1)
	}

	lowercasedPhrase := strings.ToLower(phrase)
	for _, p := range globals.History.Data {
		isSamePhrase := strings.ToLower(p.Phrase) == lowercasedPhrase
		isSameTargetLang := p.To == strings.ToLower(targetLang)
		if isSamePhrase && isSameTargetLang {
			fmt.Printf("Translation: %s\n", p.Translation)
			return
		}
	}

	checkers.CheckInternet()
	checkers.CheckOpenAiKey(globals.OpenAiKey)

	prompt := fmt.Sprintf(
		"How do you say %s in %s? Offer the translation directly. Thanks!",
		phrase,
		targetLang,
	)

	end := loader.PrintProgress("Translating...")
	defer end()

	result, err := requests.RequestGpt4Translation(prompt, globals.OpenAiKey)
	if err != nil {
		panic(err)
	}

	end()

	for _, choice := range result.Choices {
		translation := choice.Message.Content

		globals.History.Data = append(
			globals.History.Data,
			types.HistoryRecord{
				Phrase:      phrase,
				Translation: translation,
				To:          strings.ToLower(targetLang),
				CreatedAt:   time.Now().UTC(),
			},
		)
		fmt.Printf("Translation: %s\n", translation)
	}

	file, err := json.MarshalIndent(globals.History, "", " ")
	if err != nil {
		log.Fatalf("Failed to marshal json: %s\n", err.Error())
		return
	}

	filePath := filepath.Join(globals.RootDir, "history.json")
	err = os.WriteFile(filePath, file, 0644)
	if err != nil {
		log.Fatalf("Failed to update history file: %s\n", err.Error())
		return
	}
}
