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
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var TemplateFUNctionMap = template.FuncMap{
	"title":   strings.Title,
	"toLower": strings.ToLower,
}

const FunctionGenerationTemplate = `
package {{ .Folder | toLower  }}

import (
    "github.ablevets.com/Digital-Transformation/av/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/spf13/cobra"
)


// GetAddonOptions the command line options
type {{ .NoExtensionFilename | title }}Options struct {
	*	opts.CommonOptions
	batch       bool
}

var (
	{{ .Folder | toLower  }}_{{ .Filename | toLower  }}_long = templates.LongDesc("
{{ .LongDescription }}
")

	{{ .Folder | toLower }}_{{ .Filename | toLower }}_example = templates.Examples("
{{ .ExampleString }}
")
)

func NewCmd{{ .Folder | title  }}{{ .Filename | title }}(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &{{ .Folder | title }}{{ .Filename | title  }}Options {
        CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "{{ .CommandUse | toLower }}",
		Short:   "{{ .ShortDescription }}",
		Long:    {{ .Folder | toLower }}_{{ .Filename | toLower }}_long,
		Example: {{ .Folder | toLower }}_{{ .Filename | toLower }}_example,
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
func (o *{{ .Folder | title  }}{{ .Filename | title }}Options) Run() error {

    // You must still add the NewCmd{{ .Folder | title  }}{{ .Filename | title }} to a base command though!
    //   On a base command you need the line
    // 'cmd.AddCommand(NewCmd{{ .Folder | title  }}{{ .Filename | title }}(commonOpts))''
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
	NoExtensionFilename	string

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
			o.NoExtensionFilename = o.Filename
			log.Logger().Debugf("No Extension var set to: %s", o.NoExtensionFilename)
			o.Filename = o.Filename+".go"
		} else {
			log.Logger().Debugln("Not adding .go extension")

			var extension = filepath.Ext(o.Filename)
			o.NoExtensionFilename = o.NoExtensionFilename[0:len(o.Filename)-len(extension)]
		}
	}
	var fullFilePath = "./pkg/cmd/" + util.StripTrailingSlash(o.Folder) + "/" + o.Filename
	b, err := util.FileExists(fullFilePath)
	if b {
		response := util.Confirm("Are you Sure you want to override the file that already exists? This is NOT recommended", false, "that file name already exists, confirming this will override it", o.In, o.Out, o.Err)
		if !response { //answered no I don't
			return nil
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


	if exists, _ := util.DirExists("./pkg/cmd/"+o.Folder); !exists {
		err := os.MkdirAll("./pkg/cmd/"+o.Folder, 0760)
		if err != nil {
			return errors.Wrap(err, "couldn't make dir for folder")
		}

	}

	t := template.Must(template.New("template").Funcs(TemplateFUNctionMap).Parse(FunctionGenerationTemplate))
	if t == nil {
		return errors.Wrap(nil, "something when wrong parsing the template")
	}
	f, err := os.Create(fullFilePath)
	if err != nil {
		return err
	}
	err = t.Execute(f, o)
	if err != nil {
		return errors.Wrapf(err, "Error executing template %s", t.Name())
	}

	return nil
}


// ParseFiles creates a new Template and parses the template definitions from
// the named files. The returned template's name will have the base name and
// parsed contents of the first file. There must be at least one file.
// If an error occurs, parsing stops and the returned *Template is nil.
//
// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
// For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template
// named "foo", while "a/foo" is unavailable.
//func ParseFiles(filenames ...string) (*Template, error) {
//	return parseFiles(nil, filenames...)