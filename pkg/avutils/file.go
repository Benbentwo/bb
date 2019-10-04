package avutils

import (
	"bufio"
	"github.com/jenkins-x/jx/pkg/log"
	"os"
	"strings"
)

var logs = log.Logger()

func DoesFileContainString( s string, pathToFile string) (bool, error) {
	replacer := strings.NewReplacer("~", os.Getenv("HOME"))
	pathToFile = replacer.Replace(pathToFile)
	logs.Debugf("Looking for text : %s", s)
	logs.Debugf("In File          : %s", pathToFile)

	f, err := os.Open(pathToFile)
	if err != nil {
		logs.Errorf("Error Opening file: %s, Error: %s", pathToFile, err)
		return false, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), s) {
			return true, nil
		}
	}
	if someError := scanner.Err(); someError != nil {
		return false, err
	}
	return false, nil
}