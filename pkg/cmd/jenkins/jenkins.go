package jenkins

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/spf13/cobra"
)

// options for the command
type JenkinsOptions struct {
	*	opts.CommonOptions
}


func NewCmdJenkins(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &JenkinsOptions {
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "jenkins",
		Short:   "Jenkins Utilities and base command",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	options.AddJenkinsFlags(cmd)
	// the line below (Section to...) is for the generate-function command to add a template_command to.
	// the line above this and below can be deleted.
	// DO NOT DELETE THE FOLLOWING LINE:
	// Section to add commands to:
	cmd.AddCommand(NewCmdJenkinsConnect(commonOpts))
	return cmd
}

// Run implements this command
func (o *JenkinsOptions) Run() error {
	return o.Cmd.Help()
}


func (o *JenkinsOptions) AddJenkinsFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}
