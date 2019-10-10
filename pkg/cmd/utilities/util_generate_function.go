package utilities

import (
	"github.ablevets.com/Digital-Transformation/av/pkg/avutils"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"regexp"
	"text/template"
)

const FunctionGenerationTemplate = `
package {{ .Folder | strings.ToLower }}

import (
	"github.ablevets.com/Digital-Transformation/av/pkg/avutils"
    "github.ablevets.com/Digital-Transformation/av/pkg/log"
	"github.com/fatih/structs"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)


// GetAddonOptions the command line options
type {{ .Folder | strings.Title }}{{.Filename | strings.Title}}Options struct {
	*opts.CommonOptions
	batch       bool
}

var (
	{{.Folder | strings.ToLower }}_{{.Filename | strings.ToLower }}_long = templates.LongDesc("
{{.LongDescription}}
")

	{{.Folder | strings.ToLower}}_{{.Filename | strings.ToLower}}_example = templates.Examples("
{{.ExampleString}}
")
)

func NewCmd{{.Folder | strings.Title }}{{.Filename | strings.Title}}(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &{{.Folder | strings.Title}}{{.Filename | strings.Title }}Options {
        CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "{{.CommandUse | strings.toLower}}",
		Short:   "{{.ShortDescription}}",
		Long:    {{.Folder | strings.toLower}}_{{.Filename | strings.toLower}}_long,
		Example: {{.Folder | strings.toLower}}_{{.Filename | strings.toLower}}_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	// this command is not intended to be run in batch mode...
	cmd.Flags().BoolVarP(&options.batch, "batch", "b", false, "Batch commands don't prompt for user input")

	return cmd
}


// Run implements this command
func (o *{{ .Folder | strings.Title }}{{.Filename | strings.Title}}Options) Run() error {

    // You must still add the NewCmd{{.Folder | strings.Title }}{{.Filename | strings.Title}} to a base command though!
    //   On a base command you need the line
    // 'cmd.AddCommand(NewCmd{{.Folder | strings.Title }}{{.Filename | strings.Title}}(commonOpts))''
    //   then rebuild the binary!
    log.Logger().Infof("Nice Job configuring this to run %s", path)

}
`

type SupportedOptions struct {
	UtilOptions
	opts.CommonOptions
}
// GetAddonOptions the command line options
type UtilGenerateFunctionOptions struct {
	UtilOptions
	isBaseCommand bool

	Folder   string
	Filename string

	LongDescription  string
	ExampleString    string
	CommandUse       string
	ShortDescription string
	SupportedOptions SupportedOptions
	ChosenOption     string

}

var (
	util_generate_function_long = templates.LongDesc(`
		Attempts to generate a go file to help the development of this application

`)

	util_generate_function_example = templates.Examples(`
		# Utility to search a file for a string
		av util generate function utilities util_generate_function

		# Don't ask questions - run in batch mode
	`)
)

func NewCmdUtilityGenerateFunction(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &UtilGenerateFunctionOptions{
		UtilOptions: UtilOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "generate-function",
		Short:   "Generates a go file for adding a new command",
		Long:    util_generate_function_long,
		Example: util_generate_function_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	// this command is not intended to be run in batch mode...
	//cmd.Flags().BoolVarP(&options.isBaseCommand, "base-command", "b", false, "Base Commands contain other commands, like av util")
	cmd.Flags().StringVarP(&options.Folder, "Folder-name", "d", "", "Folder to create the file in")
	cmd.Flags().StringVarP(&options.Filename, "file-name", "f", "", "File to create in a Folder")

	return cmd
}


// Run implements this command
func (o *UtilGenerateFunctionOptions) Run() error {
	var err error
	if o.isBaseCommand {
		return errors.Errorf("Base Command is not currently supported")
	}
	if o.Folder == "" {
		//o.Folder, err = util.PickValue("What Folder would you like to put this in", "",true, "Folder inside of cmd of this project root", o.In, o.Out, o.Err)
		o.Folder, err = avutils.Pick(o.CommonOptions, "What Folder would you like this in?", avutils.ListSubDirectories("./pkg/cmd/"), "dev")
		if err != nil {
			return err
		}
	}
	if o.Filename == "" {
		o.Filename, err = util.PickValue("What would you like to call the file", "",true, "File name should follow the structure of foldername_filename", o.In, o.Out, o.Err)
		if err != nil {
			return err
		}
		matched, _ := regexp.MatchString(`(.*\.go)`, o.Filename)
		if !matched {
			log.Logger().Debugln("adding .go extension")
			o.Filename = o.Filename+".go"
		} else {
			log.Logger().Debugln("Not adding .go extension")
		}
	}
	//o.ChosenOption, err = avutils.Pick(o.CommonOptions, "What Set of Options would like to use", structs.Names(&SupportedOptions{}), "CommonOptions")
	//if err != nil {
	//	return err
	//}
	o.CommandUse, err = util.PickValue("What would you like for the command use, this should be a single word, or hyphenated", "",true, "Command Use", o.In, o.Out, o.Err)
	if err != nil {
		return err
	}
	o.ShortDescription, err = util.PickValue("What would you like for the short description, this should be a single word, or hyphenated", "",true, "Long Description", o.In, o.Out, o.Err)
	if err != nil {
		return err
	}
	o.LongDescription, err = util.PickValue("What would you like for the long description", "",true, "Long Description", o.In, o.Out, o.Err)
	if err != nil {
		return err
	}
	o.ExampleString, err = util.PickValue("What would you like for the example command", "",true, "Example command", o.In, o.Out, o.Err)
	if err != nil {
		return err
	}


	t := template.Must(template.New("template").Parse(FunctionGenerationTemplate))
	if t == nil {
		return errors.Wrap(nil, "something when wrong parsing the template")
	}
	err = t.Execute(o.Out, o)
	if err != nil {
		return errors.Wrapf(err, "Error executing template %s", t.Name())
	}

	return nil
}
