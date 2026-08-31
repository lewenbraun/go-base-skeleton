// Package app is the command surface of the project: argument dispatch,
// usage text and exit codes. Replace the stub with real commands.
package app

import (
	"fmt"
	"os"
)

const usageText = `project - one line description of the tool

Usage:
  project <command> [flags]
`

// Main dispatches args and returns the process exit code.
func Main(args []string) int {
	if len(args) == 0 {
		fmt.Print(usageText)
		return 2
	}
	fmt.Fprintln(os.Stderr, "unknown command:", args[0])
	fmt.Print(usageText)
	return 2
}
