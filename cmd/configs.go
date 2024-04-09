package cmd

import (
	"encoding/json"
	"fmt"
	"os"

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

	babelfish configs update defaults targetLanguage "spanish"
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
	switch args[0] {
	case "defaults":

		switch args[1] {
		case "targetLanguage":
			if len(args) >= 3 {
				globals.Configs.Defaults.TargetLanguage = args[2]
			} else {
				fmt.Println("Incorrect usage.")
				os.Exit(1)
			}

		case "stream":
			if args[2] == "true" {
				globals.Configs.Defaults.Stream = true
			} else if args[2] == "false" {
				globals.Configs.Defaults.Stream = false
			} else {
				fmt.Println("Invalid value for stream, must be true or false")
				os.Exit(1)
			}

		default:
			fmt.Printf("Unknown config field \"%s\" found\n", args[1])
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
