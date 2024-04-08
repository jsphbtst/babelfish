package loader

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func PrintProgress(message string) func() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" %s", message)

	_ = s.Color("bold", "green")
	s.Start()
	return func() {
		s.Stop()

		// "NOTE(fatih) the spinner library doesn't clear the line properly,
		// hence remove it ourselves. This line should be removed once it's
		// fixed in upstream.  https://github.com/briandowns/spinner/pull/117"
		// -- Grabbed from PlanetScale CLI Repo:
		// https://github.com/planetscale/cli/blob/main/internal/printer/printer.go#L143
		fmt.Println("\n\r\033[2K")
	}
}
