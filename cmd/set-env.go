package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var setEnvCmd = &cobra.Command{
	Use:   "set-env",
	Short: "Environment Settings",
	Long: `Environment Settings.

	For example:

	babelfish set-env openai MY-ENV-KEY-HERE
`,
	Run: runSetEnvCmd,
}

func init() {
	rootCmd.AddCommand(setEnvCmd)
}

// TODO: decide later if we want users to be able
// to see their own access keys
func runSetEnvCmd(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Printf("Expected 2 arguments, got %d\n", len(args))
		os.Exit(1)
	}

	if args[0] != "openai" {
		fmt.Println("Service currently unsupported.")
		os.Exit(1)
	}

	filePath := filepath.Join(globals.RootDir, "openai-access-key")
	keyFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer keyFile.Close()

	_, err = keyFile.Write([]byte(args[1]))
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully saved OpenAI API Key.")
}
