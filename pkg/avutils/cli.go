package avutils

import (
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"gopkg.in/AlecAivazis/survey.v1"
)

func Pick(o *opts.CommonOptions, message string, names []string, defaultChoice string) (string, error) {
	if len(names) == 0 {
		return "", nil
	}
	if len(names) == 1 {
		return names[0], nil
	}
	name := ""
	prompt := &survey.Select{
		Message: message,
		Options: names,
		Default: defaultChoice,
	}

	surveyOpts := survey.WithStdio(o.In, o.Out, o.Err)
	err := survey.AskOne(prompt, &name, nil, surveyOpts)
	return name, err
}
