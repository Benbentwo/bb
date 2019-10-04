package avutils

import (
	"bufio"
	"github.com/jenkins-x/jx/pkg/log"
	"os"
	"strings"
)

var logs = log.Logger()

func DoesFileContainString( s string, pathToFile string) (bool, int, error) {
	replacer := strings.NewReplacer("~", os.Getenv("HOME"))
	pathToFile = replacer.Replace(pathToFile)
	logs.Debugf("Looking for text : %s", s)
	logs.Debugf("In File          : %s", pathToFile)

	f, err := os.Open(pathToFile)
	if err != nil {
		logs.Errorf("Error Opening file: %s, Error: %s", pathToFile, err)
		return false, -1, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), s) {
			return true,line, nil
		}
		line++
	}
	if someError := scanner.Err(); someError != nil {
		return false, -1, err
	}
	return false,-1, nil
}