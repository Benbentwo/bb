package cmd

import (
	"github.ablevets.com/Digital-Transformation/av/pkg/cmd/github"
	initialize "github.ablevets.com/Digital-Transformation/av/pkg/cmd/init"
	jenkins "github.ablevets.com/Digital-Transformation/av/pkg/cmd/jenkins"
	"github.ablevets.com/Digital-Transformation/av/pkg/cmd/setup"
	"github.ablevets.com/Digital-Transformation/av/pkg/cmd/uninstall"
	"github.ablevets.com/Digital-Transformation/av/pkg/cmd/util"
	"github.ablevets.com/Digital-Transformation/av/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/jenkins-x/jx/pkg/extensions"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
	"io"
	"os"
	"strconv"
	"strings"
)

func NewAVCommand(f clients.Factory, in terminal.FileReader, out terminal.FileWriter, err io.Writer, args []string) *cobra.Command {
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	baseCommand := &cobra.Command{
		Use:              "av",
		Short:            "AV CLI tool and utility",
		PersistentPreRun: setLoggingLevel,
		Run:              runHelp,
	}
	//log.Logger().Debugf("running %s", baseCommand.CalledAs())
	commonOpts := opts.NewCommonOptionsWithTerm(f, in, out, err)
	commonOpts.AddBaseFlags(baseCommand)
	if len(args) == 0 {
		args = os.Args
	}
	if len(args) > 1 {
		cmdPathPieces := args[1:]

		if _, _, err := baseCommand.Find(cmdPathPieces); err != nil {
			log.Logger().Errorf("%v", err)
			os.Exit(1)
		}
	}
	groups := templates.CommandGroups{
		// Section to add commands to:
		{
			Message: "Installing and initializing AV:",
			Commands: []*cobra.Command{
				initialize.NewCmdInit(commonOpts),
			},
		},
		{
			Message: "Uninstalling AV:",
			Commands: []*cobra.Command{
				uninstall.NewCmdUninstall(commonOpts),
			},
		},
		{
			Message: "AV Utility Functions:",
			Commands: []*cobra.Command{
				util.NewCmdUtil(commonOpts),
			},
		},
		{
			Message: "setup",
			Commands: []*cobra.Command{
				setup.NewCmdSetup(commonOpts),
			},
		},
		{
			Message: "Jenkins Tools",
			Commands: []*cobra.Command{
				jenkins.NewCmdJenkins(commonOpts),
			},
		},
		{
			Message: "Github Tools",
			Commands: []*cobra.Command{
				github.NewCmdGh(commonOpts),
			},
		},
	}

	groups.Add(baseCommand)
	getPluginCommandGroups := func() (templates.PluginCommandGroups, bool) {
		verifier := &extensions.CommandOverrideVerifier{
			Root:        baseCommand,
			SeenPlugins: make(map[string]string, 0),
		}
		pluginCommandGroups, managedPluginsEnabled, err := commonOpts.GetPluginCommandGroups(verifier)
		if err != nil {
			log.Logger().Errorf("%v", err)
		}
		return pluginCommandGroups, managedPluginsEnabled
	}

	templates.ActsAsRootCommand(baseCommand, []string{"options"}, getPluginCommandGroups, groups...)
	return baseCommand
}

func setLoggingLevel(cmd *cobra.Command, args []string) {
	verbose, err := strconv.ParseBool(cmd.Flag(opts.OptionVerbose).Value.String())
	if err != nil {
		log.Logger().Errorf("Unable to determine log level")
	}

	if verbose {
		err := log.SetLevel("debug")
		if err != nil {
			log.Logger().Errorf("Unable to set log level to debug")
		}
	} else {
		err := log.SetLevel("info")
		if err != nil {
			log.Logger().Errorf("Unable to set log level to info")
		}
	}
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
