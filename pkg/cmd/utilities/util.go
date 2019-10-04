package utilities

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)

type UtilOptions struct {
	*opts.CommonOptions
	Output 	string
}

var (
	get_long = templates.LongDesc(`
		# Run a utility function
			av util [FUNCTION]
`)

	get_example = templates.Examples(`
		Runs Supporting Functions, options include:
		* searchfile
	`)
)

// NewCmdGet creates a command object for the generic "get" action, which
// retrieves one or more resources from a server.
func NewCmdUtil(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &UtilOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "util FUNC [flags]",
		Short:   "Run a Utility",
		Long:    get_long,
		Example: get_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	options.AddUtilFlags(cmd)
	cmd.AddCommand(NewCmdUtilitySearchFile(commonOpts))
	return cmd
}

// Run implements this command
func (o *UtilOptions) Run() error {
	return o.Cmd.Help()
}


func (o *UtilOptions) AddUtilFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}