// +build !windows

package app

import (
	"github.com/Benbentwo/bb/pkg/cmd"
	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"os"
)

// Run runs the command, if args are not nil they will be set on the command
func Run(args []string) error {
	cmd := cmd.NewBBCommand(clients.NewFactory(), os.Stdin, os.Stdout, os.Stderr, nil)
	if args != nil {
		args = args[1:]
		cmd.SetArgs(args)
	}
	return cmd.Execute()
}
