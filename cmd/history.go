package cmd

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "History",
	Long: `History.

	For example:

	babelfish history
	babelfish history -l 1
	babelfish history --limit 1
`,
	Run: runHistoryCmd,
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().IntP("limit", "l", math.MinInt, "Limit the results of history to be returned")
}

func runHistoryCmd(cmd *cobra.Command, args []string) {
	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		panic(err)
	}

	/*
	 * Limit flag rules:
	 *  -math.MinInt means display everything.
	 *  a non-negative limit (0 or positive) restricts the output to that many items.
	 *  any other negative limit is treated as if the limit is 0, meaning no entries should be displayed.
	 */
	if limit == math.MinInt {
		limit = len(globals.History.Data)
	} else if limit < 0 {
		limit = 0
	}

	displayCount := min(limit, len(globals.History.Data))
	for i := 0; i < displayCount; i++ {
		idx := len(globals.History.Data) - i - 1

		p := globals.History.Data[idx]

		fmt.Println("Phrase: ", p.Phrase)
		fmt.Println("Translation: ", p.Translation)
		fmt.Println("Translated To: ", p.To)
		fmt.Printf("Created At: %+v\n\n", p.CreatedAt)
	}
}
