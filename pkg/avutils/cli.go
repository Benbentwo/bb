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
//
//func PickBoolean(o *opts.CommonOptions, message string, defaultValue bool) (string, error) {
//	if message == "" {
//		return "", errors.New("message Required")
//	}
//	surveyOpts := survey.WithStdio(o.In, o.Out, o.Err)
//	prompt := &survey.Confirm{
//		Message: message,
//		Default: defaultValue,
//	}
//	response := ""
//	err := survey.AskOne(prompt, &response, nil, surveyOpts)
//	if err != nil {
//		return "", errors.Wrap(err, "Couldn't Understand Response")
//	}
//
//
//}