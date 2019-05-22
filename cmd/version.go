package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use: "ablevets version",
	Short: "a CLI tool to help automate tasks at ablevets",
	Long: "a command line tool built by benjamin smith to help the operations team get GOing faster. :)",
	Run: func(cmd *cobra.Command, args []interface{}) {

	},
}
