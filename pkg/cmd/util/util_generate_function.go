package util

import (
	"bufio"
	"github.ablevets.com/Digital-Transformation/av/pkg/avutils"
	"github.ablevets.com/Digital-Transformation/av/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// This should be found in the template_base.txt
const BASE_COMMAND_INSERT_LINE = `Section to add commands to:`
const BASE_COMMAND_TEMPLATE = `template_base.txt`

var TemplateFUNctionMap = template.FuncMap{
	"title":   strings.Title,
	"toLower": strings.ToLower,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
type SupportedOptions struct {
	UtilOptions
	opts.CommonOptions
}
// GetAddonOptions the command line options
type UtilGenerateFunctionOptions struct {
	UtilOptions
	isBaseCommand bool

	Folder               string
	Filename             string
	TitledFolderFilename string

	LongDescription  string
	ExampleString    string
	CommandUse       string
	ShortDescription string
	SupportedOptions SupportedOptions
	ChosenOption     string
	NoExtensionFilename	string
	TemplateFile		string

}

var (
	util_generate_function_long = templates.LongDesc(`
		Attempts to generate a go file to help the development of this application

`)

	util_generate_function_example = templates.Examples(`
		# Utility to search a file for a string
		av util generate function util util_generate_function

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
	cmd.Flags().StringVarP(&options.Folder, "Folder-name", "d", "", "Folder to create the file in")
	cmd.Flags().StringVarP(&options.Filename, "file-name", "f", "", "File to create in a Folder")

	return cmd
}


// Run implements this command
func (o *UtilGenerateFunctionOptions) Run() error {
	var err error
	o.TemplateFile, err = avutils.Pick(o.CommonOptions, "What template would you like to use?", avutils.ListFilesInDirFilter("./templates",`(.*\.txt)`), "template_command.txt")
	check(err)

	var isBase = o.TemplateFile == BASE_COMMAND_TEMPLATE
	if isBase {
		log.Info("Base Commands should be one word then the .go extension")
		log.Info("There should not be more than one base in a folder")
		log.Info("	this allows us to run `av <some base> <some other base>` and see a list of commands added to that other base.")
		log.Info("	which will create command structure similar to the commands.", nil)
	}
	if o.Folder == "" {
		// o.Folder, err = util.PickValue("What Folder would you like to put this in", "",true, "Folder inside of cmd of this project root", o.In, o.Out, o.Err)
		log.Info("What Folder would you like this in? starts in ./pkg/cmd/<your-answer>", nil)
		log.Info("You can create new ones, and subdirectories ./pkg/cmd/a/b", nil)
		// o.Folder, err = avutils.Pick(o.CommonOptions, "What Folder would you like this in (starting from pkg/cmd/...? you can create new directories", avutils.ListSubDirectories("./pkg/cmd/"), "dev")

		o.Folder, err = avutils.Pick(o.CommonOptions, "What Folder would you like this in?", avutils.ListSubDirectoriesRecusively("./pkg/cmd/"), "dev")
		if err != nil {
			return err
		}
		log.Info("Folders: %s", o.Folder)
	}
	var path = o.Folder
	splitPath := strings.Split(o.Folder,"/")
	o.Folder = splitPath[len(splitPath) -1]
	var originalFilename = ""
	if o.Filename == "" {
		if isBase {
			o.Filename, err = avutils.PickValueFromPath("What would you like to call the file", o.Folder,true, "File name should follow the structure of foldername_filename", o.In, o.Out, o.Err)
		} else {
			o.Filename, err = util.PickValue("What would you like to call the file", "",true, "File name should follow the structure of foldername_filename", o.In, o.Out, o.Err)
			originalFilename = o.Filename
			o.TitledFolderFilename = strings.Title(o.Folder)+strings.Title(RemoveGoExtension(o.Filename))
			o.Filename = strings.ToLower(o.Folder) + "_" + o.Filename
		}
		check(err)
		matched, _ := regexp.MatchString(`(.*\.go)`, o.Filename)
		if !matched {
			log.Logger().Debugln("Adding .go extension")
			o.NoExtensionFilename = o.Filename
			o.Filename = o.Filename+".go"
			originalFilename += ".go"
		} else {
			log.Logger().Debugln("Not adding .go extension")

			var extension = filepath.Ext(o.Filename)
			o.NoExtensionFilename = o.NoExtensionFilename[0:len(o.Filename)-len(extension)]
		}
	}
	var fullFilePath = util.StripTrailingSlash(path) + "/" + o.Filename
	b, err := util.FileExists(fullFilePath)
	if b {
		response := util.Confirm("Are you Sure you want to override the file that already exists? This is NOT recommended", false, "that file name already exists, confirming this will override it", o.In, o.Out, o.Err)
		if !response { //answered no I don't
			return nil
		}
	}

	{ 	// section for command stuff - braces help you collapse it in your IDE
		fileNameStripped := RemoveGoExtension(originalFilename)
		if isBase {
			o.CommandUse, err = util.PickValue("What would you like for the command use, this should be a single word, or hyphenated", fileNameStripped,true, "Command Use", o.In, o.Out, o.Err)
		} else {
			o.CommandUse, err = util.PickValue("What would you like for the command use, this should be a single word, or hyphenated", fileNameStripped,true, "Command Use", o.In, o.Out, o.Err)
		}
		check(err)

		if !isBase {
			o.ShortDescription, err = util.PickValue("What would you like for the short description, this should be a single word, or hyphenated", "",true, "Long Description", o.In, o.Out, o.Err)
			check(err)
			o.LongDescription, err = util.PickValue("What would you like for the long description", "",true, "Long Description", o.In, o.Out, o.Err)
			check(err)
			o.ExampleString, err = util.PickValue("What would you like for the example command", "",true, "Example command", o.In, o.Out, o.Err)
			check(err)
		}
	}

	var bases = make([]string, 0) //create an empty array
	if isBase {
		bases, err = FindBaseCommands("./pkg/cmd")
	}
	// File Gen - put it down here so we don't create the file till they answer all the questions
	// if they ctrl-c we don't want empty files cluttering our project.
	if exists, _ := util.DirExists(path); !exists {
		err := os.MkdirAll(path, 0760)
		if err != nil {
			return errors.Wrap(err, "couldn't make dir for folder")
		}
	}
	//
	var FunctionGenerationTemplate, errRead = ioutil.ReadFile("templates/"+o.TemplateFile)
	check(errRead)

	t := template.Must(template.New("template").Funcs(TemplateFUNctionMap).Parse(string(FunctionGenerationTemplate)))
	log.Var("t", t)

	if t == nil {
		return log.Fatal("Unable to parse template %s", nil)
	}

	// create the file they want
	f, err := os.Create(fullFilePath)
	check(err)


	err = t.Execute(f, o)
	if err != nil {
		return errors.Wrapf(err, "Error executing template %s", t.Name())
	}
	if !isBase {
		return nil // We're done here
	}
	// BASE command stuff continues on



	log.Debug("BASES: %s", bases)
	var pickedBase = ""
	pickedBase, err = avutils.Pick(o.CommonOptions, "What base file would you like to use?", bases, "")
	check(err)
	log.Var("PICKED BASE", pickedBase)

	// Time to determine what type of line we're adding, as its different in cmd.go (the main cmd)
	//   which is found by running just `av` - it shows groups
	//   vs anything else in which case we just add the New Cmd string.
	// Must match Template of generated function
	if pickedBase == "pkg/cmd/cmd.go" {

	} else {
		// NewCmdDev(commonOpts *opts.CommonOptions)
		err = addNewCmdToBaseFile(pickedBase, "cmd.AddCommand(NewCmd"+strings.Title(o.CommandUse)+"(commonOpts))\n")
	}
	check(err)

	return nil
}

func FindBaseCommands(path string) ([]string, error) {
	libRegEx, e := regexp.Compile("^[^_]*go$")
	if e != nil {
		return nil, log.Fatal("Error: %s", e, e)
	}
	var splice = make([]string, 0) //create an empty array

	e = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) && !info.IsDir(){
			splice = append(splice, path)
		}
		return nil
	})
	if e != nil {
		return nil, log.Fatal("Error: %s", e, e)
	}
	log.Info("%", splice)
	return splice, nil
}

func FindLineToInsertCommandTo(path string, search string) (int, error) {
	arr, err := avutils.FindMatchesInFile(search, path)
	check(err)
	if len(arr) == 0 {
		return -1, log.Fatal("No String `%s` found in file `%s`, Error: %s", err, search, path, err)
	}
	if len(arr) > 1 {
		log.Warn("Found multiple lines to insert to on lines %s", arr)
		// TODO add support for multiple finds incase there are multiple declarations in one file that we want to support
		//   e.g. Theres NewCmdUtil and NewCmdUtility or something.
	}
	return arr[0], nil
}

// Common practice should be make an exportable (Titled func) generic, then make a local named the same call with your default value.
func findLineToInsertCommandTo(path string) (int, error) {
	val, err := FindLineToInsertCommandTo(path, BASE_COMMAND_INSERT_LINE)
	check(err)
	return val, nil
}

func AddNewCmdToBaseFile(path string, insertString string, lineNumber int) error {
	err := InsertStringToFile(path, insertString, lineNumber)
	check(err)
	return nil
}
func addNewCmdToBaseFile(path string, functionName string) error {
	line, err := findLineToInsertCommandTo(path)
	check(err)
	err = AddNewCmdToBaseFile(path, functionName, line)
	check(err)
	return nil
}


func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
/**
 * Insert sting to n-th line of file.
 * If you want to insert a line, append newline '\n' to the end of the string.
 */
func InsertStringToFile(path, str string, index int) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for i, line := range lines {
		if i == index {
			fileContent += str
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func RemoveGoExtension(fileName string) string {
	matched, _ := regexp.MatchString(`(.*\.go)`, fileName)
	if matched {
		return fileName[0:len(fileName)-len(".go")]
	}
	return fileName
}