package github

import (
    "github.ablevets.com/Digital-Transformation/av/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)


// options for the command
type GithubCreate_issueOptions struct {
	*	opts.CommonOptions
	batch       bool
}

var (
	github_github_create_issue_long = templates.LongDesc(`
Creates a github issue
`)

	github_github_create_issue_example = templates.Examples(`
av github create issue
`)
)

func NewCmdGithubCreate_issue(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &GithubCreate_issueOptions {
       CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "issue",
		Short:   "create github issue",
		Long:    github_github_create_issue_long,
		Example: github_github_create_issue_example,
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
func (o *GithubCreate_issueOptions) Run() error {
	log.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
    return nil
}