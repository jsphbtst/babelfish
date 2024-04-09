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

	babelfish configs"
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

func init() {
	rootCmd.AddCommand(configsCmd)
	configsCmd.AddCommand(listConfigsCmd)
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
