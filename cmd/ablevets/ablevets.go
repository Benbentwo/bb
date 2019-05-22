package ablevets

import (
	"os"

	"github.ablevets.com/benjamin-smith/av/jx/cmd/jx/app"
)

// Entrypoint for jx command
func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
