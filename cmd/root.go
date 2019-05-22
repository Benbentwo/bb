package cmd

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/testdata/imports/fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "ablevets",
	Short: "a CLI tool to help automate tasks at ablevets",
	Long: "a command line tool built by benjamin smith to help the operations team get GOing faster. :)",
	Run: func(cmd *cobra.Command, args []interface{}) {

	},
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
