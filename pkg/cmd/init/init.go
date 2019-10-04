package init

import (
	"github.ablevets.com/Digital-Transformation/av/pkg/avutils"
	"github.ablevets.com/Digital-Transformation/av/pkg/log"
	_ "github.com/jenkins-x/jx/pkg/auth"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	_ "github.com/jenkins-x/jx/pkg/gits"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/spf13/cobra"
	_ "k8s.io/apimachinery/pkg/util/yaml"
	"os"
	_ "path/filepath"
	"runtime"
	"strings"
)

type InitOptions struct {
	*opts.CommonOptions
	Flags	InitFlags
}

type InitFlags struct {
	ConfigDir		string
	ProjectsDir		string
}
//type InitOptions struct {
//	Dir 			string
//
//}

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
		Short:			"Initializes the ~/.av configuration directory",
		Long:			initLong,
		Example:		initExample,
		Run:			func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	//cmd.Flags().StringVarP
	options.AddInitFlags(cmd)
	return cmd
}

func (o *InitOptions) Run() error {

	replacer := strings.NewReplacer("~", os.Getenv("HOME"))
	path := replacer.Replace(o.Flags.ConfigDir)
	path = util.StripTrailingSlash(path)
	path = path+"/.av"
	log.Blank()
	log.Logger().Debugln("ConfigDir Flag set to: "+o.Flags.ConfigDir)
	log.Logger().Debugln("path var set to: "+path)
	log.Blank()
	// if it doesn't already exist create it
	exists, err := util.DirExists(path)
	if err != nil {
		return err
	}
	if !exists {
		err = os.MkdirAll(path, avutils.DefaultWritePermissions)
		if err != nil {
			return err
		}
	}

	// Add to the bash profile
	if os.Getenv("AV_ENABLED") != path {
		err = os.Setenv("AV_HOME", path)
		if err != nil {
			return err
		}
		//exists, err = util.FileExists("~/.bash_profile")
		f, err := os.OpenFile(replacer.Replace("~/.bash_profile"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		var carriage = ""
		if runtime.GOOS == "windows" {
			carriage = "\r\n"
		} else {
			carriage = "\n"
		}
		if _, err = f.WriteString("export AV_HOME="+path+carriage); err != nil {
			panic(err)
		}
		log.Logger().Debugf("Updated Bash Profile to include AV_HOME")
	}
	exists, err = util.FileExists(path+"/config.yaml")
	if err != nil {
		return err
	}
	if exists {
		log.Logger().Infof("AV Directory configured to %s", path)
		log.Blank()
		log.Logger().Debugf("AV_HOME is set to %s", os.Getenv("AV_HOME"))
		// TODO add git yaml marshal and questions
		return nil
	}

	log.Logger().Infof("AV Directory configured to %s", path)
	log.Blank()
	//saver := auth.FileAuthConfigSaver{
	//	FileName: path + "/config.yaml",
	//}
	//authConfig := &auth.AuthConfig{
	//
	//}
	//saver.SaveConfig(authConfig)
	//gits.PickNewOrExistingGitRepository(false,)
	return nil
}

func(o *InitOptions) AddInitFlags(cmd *cobra.Command) {
	// add flags
	cmd.Flags().StringVarP(&o.Flags.ProjectsDir, "project-dir", "p", "~/dev", "The Directory you would like to store your Projects in")

}