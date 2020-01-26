package main

import (
	"os"

	"github.ablevets.com/Digital-Transformation/av/cmd/av/app"
)

// Entrypoint for jx command
func main() {
	if err := app.Run(nil); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
