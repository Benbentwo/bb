
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
type Setup_github_profileOptions struct {
	*	opts.CommonOptions
	batch       bool
}

var (
	setup_setup_github_profile_long = templates.LongDesc(`
A series of questions designed for seting up a profile, specifically git
`)

	setup_setup_github_profile_example = templates.Examples(`
av setup git
`)
)

func NewCmdSetup_github_profile(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &Setup_github_profileOptions {
        CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "git",
		Short:   "asks the questions for setting up a github profile",
		Long:    setup_setup_github_profile_long,
		Example: setup_setup_github_profile_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	cmd.Flags().BoolVarP(&options.batch, "batch", "b", false, "Batch commands don't prompt for user input")

	return cmd
}


// Run implements this command
func (o *Setup_github_profileOptions) Run() error {

    // You must still add the NewCmdSetupSetup_github_profile.Go to a base command though!
    //   On a base command you need the line
    // 'cmd.AddCommand(NewCmdSetupSetup_github_profile.Go(commonOpts))''
    //   then rebuild the binary!
    path := os.Getenv("PWD")
    log.Logger().Infof("Nice Job configuring this to run %s", path)
    return nil

}
