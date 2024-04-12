package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "History",
	Long: `History.

	For example:

	translate-cli history
`,
	Run: runHistoryCmd,
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().IntP("limit", "l", -1, "Limit the results of history to be returned")
}

func runHistoryCmd(cmd *cobra.Command, args []string) {
	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		panic(err)
	}

	if limit == -1 {
		limit = len(globals.History.Data)
	}

	for i, p := range globals.History.Data {
		if i+1 > limit {
			return
		}
		fmt.Println("Phrase: ", p.Phrase)
		fmt.Println("Translation: ", p.Translation)
		fmt.Println("Translated To: ", p.To)
		fmt.Printf("Created At: %+v\n\n", p.CreatedAt)
	}
}
