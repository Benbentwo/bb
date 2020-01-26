package utilities

import (
	"io/ioutil"
	"github.com/benbentwo/bb/pkg/log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	DefaultWritePermissions = 0760
)

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	h := os.Getenv("USERPROFILE") // windows
	if h == "" {
		h = "."
	}
	return h
}

// Checks fi the BB_HOME variable is set, if it isn't it makes it in the default directory
func ConfigDir() (string, error) {
	path := os.Getenv("BB_HOME")
	if path != "" {
		return path, nil
	}
	h := HomeDir()
	path = filepath.Join(h, ".bb")
	err := os.MkdirAll(path, DefaultWritePermissions)
	if err != nil {
		return "", err
	}
	return path, nil
}

// KubeConfigFile gets the .kube/config file
func KubeConfigFile() string {
	path := os.Getenv("KUBECONFIG")
	if path != "" {
		return path
	}
	h := HomeDir()
	return filepath.Join(h, ".kube", "config")
}

// JXBinLocation finds the BB config directory and creates a bin directory inside it if it does not already exist. Returns the BB bin path
func BBBinLocation() (string, error) {
	c, err := ConfigDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(c, "bin")
	err = os.MkdirAll(path, DefaultWritePermissions)
	if err != nil {
		return "", err
	}
	return path, nil
}

// JXBinaryLocation Returns the path to the currently installed JX binary.
func BBBinaryLocation() (string, error) {
	return BBBinaryLocation(os.Executable)
}

func BbBinaryLocation(osExecutable func() (string, error)) (string, error) {
	bbProcessBinary, err := osExecutable()
	if err != nil {
		log.Logger().Debugf("bbProcessBinary error %s", err)
		return bbProcessBinary, err
	}
	log.Logger().Debugf("bbProcessBinary %s", bbProcessBinary)
	// make it absolute
	bbProcessBinary, err = filepath.Abs(bbProcessBinary)
	if err != nil {
		log.Logger().Debugf("bbProcessBinary error %s", err)
		return bbProcessBinary, err
	}
	log.Logger().Debugf("bbProcessBinary %s", bbProcessBinary)

	// if the process was started form a symlink go and get the absolute location.
	bbProcessBinary, err = filepath.EvalSymlinks(bbProcessBinary)
	if err != nil {
		log.Logger().Debugf("bbProcessBinary error %s", err)
		return bbProcessBinary, err
	}

	log.Logger().Debugf("bbProcessBinary %s", bbProcessBinary)
	path := filepath.Dir(bbProcessBinary)
	log.Logger().Debugf("dir from '%s' is '%s'", bbProcessBinary, path)
	return path, nil
}
func ListSubDirectories(inputDir string) []string {
	inputDir = HomeReplace(inputDir)
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Logger().Errorf("Couldn't list files in %s", inputDir)
	}
	var splice = make([]string, 0)

	for _, f := range files {
		if f.IsDir() {
			log.Logger().Debugln(f.Name())
			splice = append(splice, f.Name())
		}
	}
	return splice
}

// I realize the above function and this could be joined with a boolean parameter but with the different implementation
// I didn't feel like doing it immediately.
func ListSubDirectoriesRecusively(inputDir string) []string {
	var splice = make([]string, 0)
	e := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		// log.Debug("Walking Path: %s", path)
		if err == nil && info.IsDir(){
			splice = append(splice, path)
		}
		return nil
	})
	errors.Check(e)
	return splice
}

func ListFilesInDir(inputDir string) []string {
	inputDir = HomeReplace(inputDir) //replace ~
	files, err := ioutil.ReadDir(inputDir) //get an array of file objects
	if err != nil {
		log.Logger().Errorf("Couldn't list files in %s", inputDir)
	}
	var splice = make([]string, 0) //create an empty array

	for _, f := range files { //for each file, get the name and append it to the list
		if !f.IsDir() {
			log.Logger().Debugln(f.Name())
			splice = append(splice, f.Name())
		}
	}
	return splice
}
func ListFilesInDirFilter(inputDir string, filter string) []string {
	inputDir = HomeReplace(inputDir) //replace ~
	files, err := ioutil.ReadDir(inputDir) //get an array of file objects
	if err != nil {
		log.Logger().Errorf("Couldn't list files in %s", inputDir)
	}
	var splice = make([]string, 0) //create an empty array

	for _, f := range files { //for each file, get the name and append it to the list
		matched, err := regexp.MatchString(filter, f.Name())
		if err != nil {
			return nil
		}
		if !f.IsDir() && matched{
			log.Logger().Debugln(f.Name())
			splice = append(splice, f.Name())
		}
	}
	return splice
}

func HomeReplace(input string) string {
	return strings.NewReplacer("~", os.Getenv("HOME")).Replace(input)
}

