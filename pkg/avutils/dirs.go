package avutils

import (
	"github.ablevets.com/Digital-Transformation/av/pkg/log"
	"io/ioutil"
	"os"
	"path/filepath"
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

// Checks fi the AV_HOME variable is set, if it isn't it makes it in the default directory
func ConfigDir() (string, error) {
	path := os.Getenv("AV_HOME")
	if path != "" {
		return path, nil
	}
	h := HomeDir()
	path = filepath.Join(h, ".av")
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

// JXBinLocation finds the AV config directory and creates a bin directory inside it if it does not already exist. Returns the AV bin path
func AVBinLocation() (string, error) {
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
func AVBinaryLocation() (string, error) {
	return AvBinaryLocation(os.Executable)
}

func AvBinaryLocation(osExecutable func() (string, error)) (string, error) {
	avProcessBinary, err := osExecutable()
	if err != nil {
		log.Logger().Debugf("avProcessBinary error %s", err)
		return avProcessBinary, err
	}
	log.Logger().Debugf("avProcessBinary %s", avProcessBinary)
	// make it absolute
	avProcessBinary, err = filepath.Abs(avProcessBinary)
	if err != nil {
		log.Logger().Debugf("avProcessBinary error %s", err)
		return avProcessBinary, err
	}
	log.Logger().Debugf("avProcessBinary %s", avProcessBinary)

	// if the process was started form a symlink go and get the absolute location.
	avProcessBinary, err = filepath.EvalSymlinks(avProcessBinary)
	if err != nil {
		log.Logger().Debugf("avProcessBinary error %s", err)
		return avProcessBinary, err
	}

	log.Logger().Debugf("avProcessBinary %s", avProcessBinary)
	path := filepath.Dir(avProcessBinary)
	log.Logger().Debugf("dir from '%s' is '%s'", avProcessBinary, path)
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

func HomeReplace(input string) string {
	return strings.NewReplacer("~", os.Getenv("HOME")).Replace(input)
}

