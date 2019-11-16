package jenkins

import (
	"github.com/Benbentwo/bb/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)

// options for the command
type JenkinsCreateOptions struct {
	*opts.CommonOptions
	batch bool
}

var (
	jenkins_jenkins_create_long = templates.LongDesc(`
initialize the config dir, and create a github profile if it doesn't exist
`)

	jenkins_jenkins_create_example = templates.Examples(`
 
`)
)

func NewCmdJenkinsCreate(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &JenkinsCreateOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   "creates a github profile",
		Long:    jenkins_jenkins_create_long,
		Example: jenkins_jenkins_create_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}

	return cmd
}

// Run implements this command
func (o *JenkinsCreateOptions) Run() error {
	log.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
	return nil
}
