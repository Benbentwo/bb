package github

import (
	Err "github.com/Benbentwo/bb/pkg/cmd/errors"
	"github.com/Benbentwo/bb/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// options for the command
type CreateGitServerOptions struct {
	CreateOptions
	Name string
	Kind string
	URL  string
}

var (
	github_github_create_gitserver_long = templates.LongDesc(`
Adds a git server object to your ~/.bb/gitAuth.yaml configuration. You can add different users to each git server. A git server is like github.com or github.com, there is also bitbucket support
`)

	github_github_create_gitserver_example = templates.Examples(`
bb gh create git-server
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

	cmd.Flags().StringVarP(&options.Name, "name", "n", "", "The name for the Git server being created")
	cmd.Flags().StringVarP(&options.Kind, "kind", "k", "", "The kind of Git server being created")
	cmd.Flags().StringVarP(&options.URL, "url", "u", "", "The git server URL")

	return cmd
}

// Run implements this command
func (o *CreateGitServerOptions) Run() error {
	args := o.Args
	kind := o.Kind
	if kind == "" {
		if len(args) < 1 {
			return util.MissingOption("kind")
		}
		kind = args[0]
	}
	name := o.Name
	if name == "" {
		name = kind
	}
	gitUrl := o.URL
	if gitUrl == "" {
		if len(args) > 1 {
			gitUrl = args[1]
		} else {
			// lets try find the git URL based on the provider
			gitUrl = GetDefaultUrlFromGitServer(kind)
			// if serviceName != "" {
			// 	url, err := o.FindService(serviceName)
			// 	if err != nil {
			// 		return errors.Wrapf(err, "Failed to find %s Git service %s", kind, serviceName)
			// 	}
			// 	gitUrl = url
			// }
		}
	}
	if gitUrl == "" {
		return util.MissingOption("url")
	}
	factory := clients.NewFactory()
	log.Info("factory: %s", factory)
	// configService, err := auth.NewFileAuthConfigService(GitAuthConfigFile)
	authConfigSvc, err := clients.Factory.CreateAuthConfigService(factory, GitAuthConfigFile, "")
	if err != nil {
		return errors.Wrap(err, "failed to create CreateGitAuthConfigService")
	}

	log.Info("authconfigsvc: %s", authConfigSvc)
	config := authConfigSvc.Config()
	log.Info("1")
	server := config.GetOrCreateServerName(gitUrl, name, kind)
	log.Info("server: %s", server)
	log.Info("2")
	config.CurrentServer = gitUrl
	log.Info("3")
	err = authConfigSvc.SaveConfig()
	log.Info("4")
	if err != nil {
		log.Info("4err")
		return errors.Wrap(err, "failed to save GitAuthConfigService")
	}
	log.Logger().Infof("Added Git server %s for URL %s", util.ColorInfo(name), util.ColorInfo(gitUrl))

	log.Info("5")
	log.Var("server", server)
	err = authConfigSvc.SaveConfig()
	Err.Check(err)
	return nil
}
