package {{ .Folder | toLower  }}

import (
    "github.com/Benbentwo/bb/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)


// options for the command
type {{ .TitledFolderFilename  }}Options struct {
	*	opts.CommonOptions
	batch       bool
}

var (
	{{ .Folder | toLower  }}_{{ .NoExtensionFilename | toLower  }}_long = templates.LongDesc(`
{{ .LongDescription }}
`)

	{{ .Folder | toLower }}_{{ .NoExtensionFilename | toLower }}_example = templates.Examples(`
{{ .ExampleString }}
`)
)

func NewCmd{{ .TitledFolderFilename }}(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &{{ .TitledFolderFilename  }}Options {
       CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "{{ .CommandUse | toLower }}",
		Short:   "{{ .ShortDescription }}",
		Long:    {{ .Folder | toLower }}_{{ .NoExtensionFilename | toLower }}_long,
		Example: {{ .Folder | toLower }}_{{ .NoExtensionFilename | toLower }}_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}

	return cmd
}


// Run implements this command
func (o *{{ .TitledFolderFilename  }}Options) Run() error {
	log.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
    return nil
}