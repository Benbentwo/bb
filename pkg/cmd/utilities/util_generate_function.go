package utilities

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"text/template"

	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
)

type SupportedOptions struct {
	UtilOptions
	opts.CommonOptions
}
// GetAddonOptions the command line options
type UtilGenerateFunctionOptions struct {
	UtilOptions
	isBaseCommand bool

	folder				string
	filename			string

	longDescription		string
	exampleString		string
	commandUse			string
	shortDescription	string
	supportedOptions	SupportedOptions
	chosenOption		string

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
	cmd.Flags().StringVarP(&options.folder, "folder-name", "d", "", "Folder to create the file in")
	cmd.Flags().StringVarP(&options.filename, "file-name", "f", "", "File to create in a folder")

	return cmd
}


// Run implements this command
func (o *UtilGenerateFunctionOptions) Run() error {
	var err error
	if o.isBaseCommand {
		return errors.Errorf("Base Command is not currently supported")
	}
	if o.folder == "" {
		o.folder, err = util.PickValue("What would you like to call the file", "",true, "File name should follow the structure of foldername_filename", o.In, o.Out, o.Err)
		if err != nil {
			return err
		}
	}
	if o.filename == "" {
		o.filename, err = util.PickValue("What would you like to call the file", "",true, "File name should follow the structure of foldername_filename", o.In, o.Out, o.Err)
		if err != nil {
			return err
		}
	}
	//o.chosenOption, err = avutils.Pick(o.CommonOptions, "What Set of Options would like to use", structs.Names(&SupportedOptions{}), "CommonOptions")
	//if err != nil {
	//	return err
	//}
	o.commandUse, err = util.PickValue("What would you like for the command use, this should be a single word, or hyphenated", "",true, "Command Use", o.In, o.Out, o.Err)
	if err != nil {
		return err
	}
	o.shortDescription, err = util.PickValue("What would you like for the command use, this should be a single word, or hyphenated", "",true, "Long Description", o.In, o.Out, o.Err)
	if err != nil {
		return err
	}
	o.longDescription, err = util.PickValue("What would you like for the long description", "",true, "Long Description", o.In, o.Out, o.Err)
	if err != nil {
		return err
	}
	o.exampleString, err = util.PickValue("What would you like for the example command", "",true, "Example command", o.In, o.Out, o.Err)
	if err != nil {
		return err
	}


	tmpl, err := template.New("newFunction").ParseFiles("./templates/template_command")
	if err != nil {
		return err
	}
	err = tmpl.Execute(o.Out, o)
	if err != nil {
		return err
	}
	return nil
}
