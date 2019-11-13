package jenkins

import (
    "github.com/Benbentwo/bb/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)


// options for the command
type JenkinsConnectOptions struct {
	*	opts.CommonOptions
	batch       bool
}

var (
	jenkins_jenkins_connect_long = templates.LongDesc(`
something slightly longer
`)

	jenkins_jenkins_connect_example = templates.Examples(`
my example
`)
)

func NewCmdJenkinsConnect(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &JenkinsConnectOptions {
       CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "connect",
		Short:   "something short",
		Long:    jenkins_jenkins_connect_long,
		Example: jenkins_jenkins_connect_example,
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
func (o *JenkinsConnectOptions) Run() error {
	log.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
    return nil
}