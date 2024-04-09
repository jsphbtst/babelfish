package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jsphbtst/babelfish/pkg/checkers"
	"github.com/spf13/cobra"
)

var configsCmd = &cobra.Command{
	Use:   "configs",
	Short: "configs",
	Long: `Configs.

	For example:

	babelfish configs
`,
}

var listConfigsCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long: `This command allows you to list your current configs. By default, you have:
	{
		"defaults": {
			"targetLanguage": "spanish",
			"stream": false
		}
	}

	For example:

	babelfish configs list
	`,
	Run: runListConfigs,
}

var updateConfigCmd = &cobra.Command{
	Use:   "update",
	Short: "update",
	Long: `This command allows you to update a specific config in question.

	For example:

	babelfish configs update dot.separated.path TO_CHANGE
	babelfish configs update defaults.targetLanguage "spanish"
	`,
	Run: runUpdateConfig,
}

func init() {
	rootCmd.AddCommand(configsCmd)
	configsCmd.AddCommand(listConfigsCmd)
	configsCmd.AddCommand(updateConfigCmd)
}

func runListConfigs(cmd *cobra.Command, args []string) {
	configs := globals.Configs
	jsonFmt, err := json.MarshalIndent(configs, "", " ")
	if err != nil {
		fmt.Println("Failed to list configs because: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(jsonFmt))
}

func runUpdateConfig(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Incorrect usage.")
		os.Exit(1)
	}

	path := strings.Split(args[0], ".")
	updateValue := args[1]

	switch path[0] {
	case "defaults":
		switch path[1] {
		case "targetLanguage":
			if len(path) > 2 {
				fmt.Println("Config path does not exist")
				os.Exit(1)
			}

			targetLanguage := strings.ToLower(updateValue)
			if !checkers.IsSupportedLanguage(targetLanguage) {
				fmt.Println("Currently an unsupported language")
				os.Exit(1)
			}

			globals.Configs.Defaults.TargetLanguage = targetLanguage

		case "stream":
			if len(path) > 2 {
				fmt.Println("Config path does not exist")
				os.Exit(1)
			}

			if updateValue == "true" {
				globals.Configs.Defaults.Stream = true
			} else if updateValue == "false" {
				globals.Configs.Defaults.Stream = false
			} else {
				fmt.Println("Invalid value for stream, must be true or false")
				os.Exit(1)
			}

		default:
			fmt.Printf("Unknown config field \"%s\" found\n", path[1])
			os.Exit(1)
		}

	default:
		fmt.Println("Incorrect usage")
		os.Exit(1)
	}

	configsBytes, err := json.MarshalIndent(globals.Configs, "", " ")
	if err != nil {
		fmt.Printf("Failed to marshal indent configs: %s\n", err.Error())
		os.Exit(1)
	}

	err = os.WriteFile("configs.json", configsBytes, 0644)
	if err != nil {
		fmt.Printf("Failed to update configs file: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Configs saved successfully: \n%s\n", string(configsBytes))
}
