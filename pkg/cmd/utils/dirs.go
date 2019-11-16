package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	ConfigDirEnvVar         = "BB_HOME"
	ConfigDirFolderName     = ".bb"
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
func ConfigDir(envVar string, configFolder string) (string, error) {
	path := os.Getenv(envVar)
	if path != "" {
		return path, nil
	}
	h := HomeDir()
	path = filepath.Join(h, configFolder)
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

// JXBinLocation finds the bb config directory and creates a bin directory inside it if it does not already exist. Returns the bb bin path
func BinLocation() (string, error) {
	c, err := ConfigDir(ConfigDirEnvVar, ConfigDirFolderName)
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
func ThisBinaryLocation() (string, error) {
	return BinaryLocation(os.Executable)
}

func BinaryLocation(osExecutable func() (string, error)) (string, error) {
	processBinary, err := osExecutable()
	if err != nil {
		return processBinary, err
	}
	// make it absolute
	processBinary, err = filepath.Abs(processBinary)
	if err != nil {
		return processBinary, err
	}

	// if the process was started form a symlink go and get the absolute location.
	processBinary, err = filepath.EvalSymlinks(processBinary)
	if err != nil {
		return processBinary, err
	}

	path := filepath.Dir(processBinary)
	return path, nil
}
func ListSubDirectories(inputDir string) []string {
	inputDir = HomeReplace(inputDir)
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		Logger().Errorf("Couldn't list files in %s", inputDir)
	}
	var splice = make([]string, 0)

	for _, f := range files {
		if f.IsDir() {
			Logger().Debugln(f.Name())
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
		// Debug("Walking Path: %s", path)
		if err == nil && info.IsDir() {
			splice = append(splice, path)
		}
		return nil
	})
	Check(e)
	return splice
}

func ListFilesInDir(inputDir string) []string {
	inputDir = HomeReplace(inputDir)       //replace ~
	files, err := ioutil.ReadDir(inputDir) //get an array of file objects
	if err != nil {
		Logger().Errorf("Couldn't list files in %s", inputDir)
	}
	var splice = make([]string, 0) //create an empty array

	for _, f := range files { //for each file, get the name and append it to the list
		if !f.IsDir() {
			Logger().Debugln(f.Name())
			splice = append(splice, f.Name())
		}
	}
	return splice
}
func ListFilesInDirFilter(inputDir string, filter string) []string {
	inputDir = HomeReplace(inputDir)       //replace ~
	files, err := ioutil.ReadDir(inputDir) //get an array of file objects
	if err != nil {
		Logger().Errorf("Couldn't list files in %s", inputDir)
	}
	var splice = make([]string, 0) //create an empty array

	for _, f := range files { //for each file, get the name and append it to the list
		matched, err := regexp.MatchString(filter, f.Name())
		if err != nil {
			return nil
		}
		if !f.IsDir() && matched {
			Logger().Debugln(f.Name())
			splice = append(splice, f.Name())
		}
	}
	return splice
}

func HomeReplace(input string) string {
	return strings.NewReplacer("~", os.Getenv("HOME")).Replace(input)
}
