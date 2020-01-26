package utilities

import (
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"gopkg.in/AlecAivazis/survey.v1"
	"io"
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

// PickValue gets an answer to a prompt from a user's free-form input
func PickValueFromPath(message string, defaultValue string, required bool, help string, in terminal.FileReader, out terminal.FileWriter, outErr io.Writer) (string, error) {
	answer := ""
	prompt := &survey.Input{
		Message: message,
		Default: defaultValue,
		Help:    help,
	}
	validator := survey.Required
	if !required {
		validator = nil
	}
	surveyOpts := survey.WithStdio(in, out, outErr)

	err := survey.AskOne(prompt, &answer, validator, surveyOpts)
	if err != nil {
		return "", err
	}
	return answer, nil
}