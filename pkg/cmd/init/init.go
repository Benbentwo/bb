package init

import (
	"github.ablevets.com/Digital-Transformation/av/pkg/avutils"
	"github.ablevets.com/Digital-Transformation/av/pkg/log"
	"github.ablevets.com/Digital-Transformation/av/pkg/cmd/github"
	_ "github.com/jenkins-x/jx/pkg/auth"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/jenkins-x/jx/pkg/cmd/create"
	_ "github.com/jenkins-x/jx/pkg/gits"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	_ "k8s.io/apimachinery/pkg/util/yaml"
	"os"
	_ "path/filepath"
	"runtime"
	"strings"
	_ "gopkg.in/yaml.v2"
)

const (
	AV_HOME_VAR = "AV_HOME"
	AV_CONFIG_DIR = "~/.av"
)
type InitOptions struct {
	*opts.CommonOptions
	Flags	InitFlags
}

type InitFlags struct {
	ConfigDir		string
	ProjectsDir		string
}

var logs = log.Logger()

var (
	initLong = templates.LongDesc(`
		This Command will setup the configuration for later use.
		This will create configuration files in ~/.av/ that later av commands will use
`)
	initExample = templates.Examples(`
		av init	
`)
)

func NewCmdInit(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &InitOptions{
		CommonOptions: commonOpts,
	}
	cmd := &cobra.Command{
		Use:			"init",
		Short:			"Initializes the "+AV_CONFIG_DIR+" configuration directory",
		Long:			initLong,
		Example:		initExample,
		Run:			func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	options.AddInitFlags(cmd)
	// Section to add commands to:
	return cmd
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (o *InitOptions) Run() error {


	replacer := strings.NewReplacer("~", os.Getenv("HOME"))
	path := replacer.Replace(AV_CONFIG_DIR)

	log.Blank()

	// if it doesn't already exist create it
	exists, err := util.DirExists(path)
	if err != nil {
		return err
	}
	if !exists {
		logs.Debugf("Directory `~/.av` not found... creating")
		err = os.MkdirAll(path, avutils.DefaultWritePermissions)
		if err != nil {
			return err
		}
	}

	// Add to the bash profile
	if os.Getenv(AV_HOME_VAR) != path {
		logs.Debugf("AV HOME Set to %s", os.Getenv(AV_HOME_VAR))
		logs.Debugf("Path Set to %s", path)
		// set current shell
		err = os.Setenv(AV_HOME_VAR, path)
		if err != nil {
			return err
		}

		stringExists,line, err := avutils.DoesFileContainString("export AV_HOME=~/.av", "~/.bash_profile")
		if err != nil {
			return err
		}
		if line != -1 {
			logs.Debugf("Found String at line %s", line)
		}
		if !stringExists {
			f, err := os.OpenFile(replacer.Replace("~/.bash_profile"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				logs.Errorf("Couldn't open or find ~/.bash_profile")
				panic(err)
			}

			defer f.Close()
			var carriage= GetCarriageReturn()

			if _, err = f.WriteString("export AV_HOME=" + path + carriage); err != nil {
				panic(err)
			}
			log.Logger().Debugf("Updated Bash Profile to include AV_HOME")
		}
	}



	configPath := path+"/"+github.GitAuthConfigFile
	exists, err = util.FileExists(configPath)
	check(err)
	if exists {
		log.Logger().Infof("Git yaml file found.")
	} else {
		_, err = os.Create(configPath)
		check(err)
	}

	// Lets setup a Git Profile
	log.Info("Looks like you do not have any git servers configured")
	response := util.Confirm("Would you like to set one up now?", true, "Would you like to create a connection configuraiton to a git server?", o.In, o.Out, o.Err)
	if response {
		SetupGitConfigFile(configPath, *o.CommonOptions)
	}

	log.Blank()
	log.Logger().Infof("SUCCESS: AV Directory configured to %s", path)
	return nil
}

func(o *InitOptions) AddInitFlags(cmd *cobra.Command) {
	// add flags
	cmd.Flags().StringVarP(&o.Flags.ProjectsDir, "project-dir", "p", "~/dev", "The Directory you would like to store your Projects in")

}


// Writes a string to a file and returns whether or not it did exist
func WriteStringIfDoesntExist(writeString string, filePath string) bool{
	if exists, _, _ :=avutils.DoesFileContainString(writeString, filePath); !exists{
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		check(err)
		_, err =f.WriteString(writeString)
		check(err)
		return false
	}
	return true
}

func GetCarriageReturn() string{
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else {
		return "\n"
	}
}
func SetupGitConfigFile(configPath string, o opts.CommonOptions){
	carriage := GetCarriageReturn()
	WriteStringIfDoesntExist("currentserver: "+carriage, configPath)
	WriteStringIfDoesntExist("defaultusername: "+carriage, configPath)
	hasServers := WriteStringIfDoesntExist("servers: "+carriage, configPath)

	if !hasServers {

		serverName, err := util.PickValue("Git Server Name:", "", true, "What would you like to name this gitServer?", o.In, o.Out, o.Err)
		if err != nil {
			panic(err)
		}
		kind, err := util.PickName(github.ServerTypes, "What Git Server would you like to add", "What is your remote repository kind?", o.In, o.Out, o.Err)
		if err != nil {
			panic(err)
		}
		defaultUrl := github.GetDefaultUrlFromGitServer(kind)
		url, err := util.PickValue("Git Server URL:", defaultUrl, true, "What would you like to name this gitServer?", o.In, o.Out, o.Err)
		if err != nil {
			panic(err)
		}
		gitServerOptions := create.CreateGitServerOptions{
			CreateOptions: create.CreateOptions{
				CommonOptions: &o,
				DisableImport: true,
				OutDir:        configPath,
			},
			Name: serverName,
			Kind: kind,
			URL:  url,
		}
		err = gitServerOptions.Run()
		check(err)
	}
}

// Run implements the command
func  Run() error {
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
			serviceName := gitKindToServiceName[kind]
			if serviceName != "" {
				url, err := o.FindService(serviceName)
				if err != nil {
					return errors.Wrapf(err, "Failed to find %s Git service %s", kind, serviceName)
				}
				gitUrl = url
			}
		}
	}
	if gitUrl == "" {
		return util.MissingOption("url")
	}

	authConfigSvc, err := o.CreateGitAuthConfigService()
	if err != nil {
		return errors.Wrap(err, "failed to create CreateGitAuthConfigService")
	}
	config := authConfigSvc.Config()
	server := config.GetOrCreateServerName(gitUrl, name, kind)
	config.CurrentServer = gitUrl
	err = authConfigSvc.SaveConfig()
	if err != nil {
		return errors.Wrap(err, "failed to save GitAuthConfigService")
	}
	log.Logger().Infof("Added Git server %s for URL %s", util.ColorInfo(name), util.ColorInfo(gitUrl))

	err = o.EnsureGitServiceCRD(server)
	if err != nil {
		return err
	}
	return nil
}
