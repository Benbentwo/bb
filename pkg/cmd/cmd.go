package cmd

import (
	initialize "github.com/Benbentwo/bb/pkg/cmd/init"
	"github.com/Benbentwo/bb/pkg/cmd/jenkins"
	"github.com/Benbentwo/bb/pkg/cmd/uninstall"
	"github.com/Benbentwo/bb/pkg/cmd/util"
	"github.com/Benbentwo/bb/pkg/log"
	jenkinsv1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/jenkins-x/jx/pkg/extensions"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
	"io"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/sets"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"strings"
)

func NewBBCommand(f clients.Factory, in terminal.FileReader, out terminal.FileWriter, err io.Writer, args []string) *cobra.Command {
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	baseCommand := &cobra.Command{
		Use:              "bb",
		Short:            "CLI tool and utility",
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
			Message: "Installing and initializing BB:",
			Commands: []*cobra.Command{
				initialize.NewCmdInit(commonOpts),
			},
		},
		{
			Message: "Jenkins Tools",
			Commands: []*cobra.Command{
				jenkins.NewCmdJenkins(commonOpts),
			},
		},
		{
			Message: "Uninstalling BB:",
			Commands: []*cobra.Command{
				uninstall.NewCmdUninstall(commonOpts),
			},
		},
		{
			Message: "BB Utility Functions:",
			Commands: []*cobra.Command{
				util.NewCmdUtil(commonOpts),
			},
		},
	}

	groups.Add(baseCommand)
	getPluginCommandGroups := func() (templates.PluginCommandGroups, bool) {
		verifier := &extensions.CommandOverrideVerifier{
			Root:        baseCommand,
			SeenPlugins: make(map[string]string, 0),
		}
		pluginCommandGroups, managedPluginsEnabled, err := GetPluginCommandGroups(verifier)
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

// GetPluginCommandGroups returns the plugin groups
func GetPluginCommandGroups(verifier extensions.PathVerifier) (templates.PluginCommandGroups, bool,
	error) {

	otherCommands := templates.PluginCommandGroup{
		Message: "Other Commands",
	}
	groups := make(map[string]templates.PluginCommandGroup, 0)

	pathCommands := templates.PluginCommandGroup{
		Message: "Locally Available Commands:",
	}

	path := "PATH"
	if runtime.GOOS == "windows" {
		path = "path"
	}

	paths := sets.NewString(filepath.SplitList(os.Getenv(path))...)
	for _, dir := range paths.List() {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if !strings.HasPrefix(f.Name(), "jx-") {
				continue
			}

			pluginPath := filepath.Join(dir, f.Name())
			subCommand := strings.TrimPrefix(strings.Replace(filepath.Base(pluginPath), "-", " ", -1), "jx ")
			pc := &templates.PluginCommand{
				PluginSpec: jenkinsv1.PluginSpec{
					SubCommand:  subCommand,
					Description: pluginPath,
				},
				Errors: make([]error, 0),
			}
			pathCommands.Commands = append(pathCommands.Commands, pc)
			if errs := verifier.Verify(filepath.Join(dir, f.Name())); len(errs) != 0 {
				for _, err := range errs {
					pc.Errors = append(pc.Errors, err)
				}
			}
		}
	}

	pcgs := templates.PluginCommandGroups{}
	for _, g := range groups {
		pcgs = append(pcgs, g)
	}
	if len(otherCommands.Commands) > 0 {
		pcgs = append(pcgs, otherCommands)
	}
	if len(pathCommands.Commands) > 0 {
		pcgs = append(pcgs, pathCommands)
	}
	return pcgs, false, nil
}
