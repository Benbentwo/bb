package uninstall

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)


type UninstallOptions struct {
	*opts.CommonOptions
	Output 	string
	All		bool
}

const (
	uninstall_resources = `uninstall options include:

    * config
    * binary
    * all (e.g. 'all above')
    `
)

var (
	get_long = templates.LongDesc(`
		Uninstalls one or more resources.

		` + uninstall_resources + `

`)

	get_example = templates.Examples(`
		# uninstall the binary
			bb uninstall binary

		# Uninstall your current config **Cannot be undone**
			bb uninstall config

		# Uninstall EVERYTHING
			bb uninstall -a
			# or
			bb uninstall --all
	`)
)

// NewCmdGet creates a command object for the generic "get" action, which
// retrieves one or more resources from a server.
func NewCmdUninstall(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &UninstallOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "uninstall TYPE [flags]",
		Short:   "Uninstall one or more resources",
		Long:    get_long,
		Example: get_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	options.AddUninstallFlags(cmd)
	// Section to add commands to:
	cmd.AddCommand(NewCmdUninstallBinary(commonOpts))
	cmd.AddCommand(NewCmdUninstallConfig(commonOpts))
	return cmd
}

// Run implements this command
func (o *UninstallOptions) Run() error {
	if o.All {
		err := UninstallAll()
		if err != nil {
			return err
		}
		return nil
	}
	return o.Cmd.Help()
}


func (o *UninstallOptions) AddUninstallFlags(cmd *cobra.Command) {
	o.Cmd = cmd
	cmd.Flags().BoolVarP(&o.All, "all", "a", false, "Uninstall everything")
}