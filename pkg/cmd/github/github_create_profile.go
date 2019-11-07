package github

import (
    "github.ablevets.com/Digital-Transformation/av/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)


// options for the command
type GithubCreate_profileOptions struct {
	*	opts.CommonOptions
	batch       bool
}

var (
	github_github_create_profile_long = templates.LongDesc(`
Create a github profile for GH or GHE and add to your ~/.av folder
`)

	github_github_create_profile_example = templates.Examples(`
av gh create profile
`)
)

func NewCmdGithubCreate_profile(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &GithubCreate_profileOptions {
       CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "profile",
		Short:   "create a github profile",
		Long:    github_github_create_profile_long,
		Example: github_github_create_profile_example,
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
func (o *GithubCreate_profileOptions) Run() error {
	log.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
    return nil
}