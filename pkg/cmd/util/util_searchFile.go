package util

import (
	"github.com/Benbentwo/bb/pkg/utilities"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
)

// GetAddonOptions the command line options
type UtilSearchFileOptions struct {
	UtilOptions

	SearchString		string
	SearchFile			string
}

var (
	util_searchfile_long = templates.LongDesc(`
		Searches a file for a string and prints all line numbers

`)

	util_searchfile_example = templates.Examples(`
		# Utility to search a file for a string
		bb util search "hello world" HelloWorld.java
	`)
)

func NewCmdUtilitySearchFile(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &UtilSearchFileOptions{
		UtilOptions: UtilOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "searchfile String File [flags]",
		Short:   "Searches a file for a string",
		Long:    util_searchfile_long,
		Example: util_searchfile_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	cmd.Flags().StringVarP(&options.SearchString, "search-string", "s", "", "Search for this string")
	cmd.Flags().StringVarP(&options.SearchFile, "search-file", "f", "", "Search this file for the string")

	return cmd
}

// Run implements this command
func (o *UtilSearchFileOptions) Run() error {
	if o.SearchFile == "" {
		return util.MissingOption("search-file")
	}
	if o.SearchString == "" {
		return util.MissingOption("search-string")
	}
	o.SearchFile = avutils.HomeReplace(o.SearchFile)
	count, err := avutils.FindMatchesInFile(o.SearchString, o.SearchFile)
	if err != nil {
		return errors.Wrapf(err,"Could not search the file %s for the string %s.", o.SearchFile, o.SearchString)
	}
	log.Logger().Infof("Found %d instances of `%s` in the file `%s`", len(count), o.SearchString, o.SearchFile)
	if len(count) > 0 {
		log.Logger().Infof("Found on lines %d", count)
	}
	return nil
}
