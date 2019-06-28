package av

import (
	"os"

	"github.ablevets.com/benjamin-smith/av/av/cmd/av/app"
)

// Entrypoint for av command
func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
