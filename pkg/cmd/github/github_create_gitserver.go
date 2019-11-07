package github

import (
    "github.ablevets.com/Digital-Transformation/av/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)


// options for the command
type CreateGitServerOptions struct {
	CreateOptions
	Name		string
	Kind		string
	URL			string
}

var (
	github_github_create_gitserver_long = templates.LongDesc(`
Adds a git server object to your ~/.av/gitAuth.yaml configuration. You can add different users to each git server. A git server is like github.com or github.ablevets.com, there is also bitbucket support
`)

	github_github_create_gitserver_example = templates.Examples(`
av gh create git-server
`)
)

func NewCmdGithubCreate_gitserver(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &CreateGitServerOptions{
		CreateOptions: CreateOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "gitserver",
		Short:   "create a git server in your configuration",
		Long:    github_github_create_gitserver_long,
		Example: github_github_create_gitserver_example,
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
func (o *CreateGitServerOptions) Run() error {
	log.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
    return nil
}