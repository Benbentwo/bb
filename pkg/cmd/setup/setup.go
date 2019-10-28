
package setup

import (
	"github.ablevets.com/Digital-Transformation/av/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
	"os"
)


// GetAddonOptions the command line options
type SetupOptions struct {
	*	opts.CommonOptions
	batch       bool
}

var (
	setup_setup_long = templates.LongDesc(`
Base command for which other setup files will be added to
`)

	setup_setup_example = templates.Examples(`
av setup --help`)
)

func NewCmdSetup(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &SetupOptions {
        CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "setup",
		Short:   "utitilies getting started",
		Long:    setup_setup_long,
		Example: setup_setup_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}

	// Section to add commands to:
cmd.AddCommand(NewCmdDev(commonOpts))
	return cmd
}


// Run implements this command
func (o *SetupOptions) Run() error {

    // You must still add the NewCmdSetup.Go to a base command though!
    //   On a base command you need the line
    // 'cmd.AddCommand(NewCmdSetup.Go(commonOpts))''
    //   then rebuild the binary!
	path := os.Getenv("PWD")
	log.Logger().Infof("Nice Job configuring this to run %s", path)
	return nil
}
