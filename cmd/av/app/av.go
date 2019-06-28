// +build !windows

package app

import (
	"os"

	"github.ablevets.com/benjamin-smith/av/pkg/cmd"
	"github.ablevets.com/benjamin-smith/av/pkg/cmd/clients"
)

// Run runs the command, if args are not nil they will be set on the command
func Run(args []string) error {
	cmd := cmd.NewAVCommand(clients.NewFactory(), os.Stdin, os.Stdout, os.Stderr, nil)
	if args != nil {
		args = args[1:]
		cmd.SetArgs(args)
	}
	return cmd.Execute()
}
