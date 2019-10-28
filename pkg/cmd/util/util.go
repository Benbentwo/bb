package util

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/spf13/cobra"
)

type UtilOptions struct {
	*opts.CommonOptions
	Output 	string
}

func NewCmdUtil(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &UtilOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "util",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	options.AddUtilFlags(cmd)
	// Section to add commands to:
	cmd.AddCommand(NewCmdUtilitySearchFile(commonOpts))
	cmd.AddCommand(NewCmdUtilityGenerateFunction(commonOpts))
	return cmd
}

// Run implements this command
func (o *UtilOptions) Run() error {
	return o.Cmd.Help()
}


func (o *UtilOptions) AddUtilFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}