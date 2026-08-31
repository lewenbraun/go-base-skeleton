// Command project is the entry point stub of the skeleton template.
package main

import (
	"os"

	"github.com/lewenbraun/go-base-skeleton/project/internal/app"
)

func main() {
	os.Exit(app.Main(os.Args[1:]))
}
