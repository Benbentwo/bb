package uninstall

import (
	"github.ablevets.com/Digital-Transformation/av/pkg/avutils"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func UninstallAll() error {
	err := UninstallConfig()
	if err != nil {
		return errors.Wrapf(err, "%s")
	}
	err = UninstallBinary()
	if err != nil {
		return errors.Wrapf(err, "%s")
	}
	return nil
}
func UninstallConfig() error {
	err := os.RemoveAll(avutils.HomeReplace("~/.av"))
	if err != nil {
		return errors.Errorf("Couldn't delete ~/.av directory")
	}
	return nil
}

func UninstallBinary() error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.Errorf("Couldn't delete Binary at %s", dir)
	}
	err = os.Remove(dir+"/av")
	if err != nil {
		return errors.Errorf("Couldn't delete Binary at %s", dir)
	}
	return nil
}

//TODO finish implementing
func UnsetAVHome() error {
	line := 0
	for line != -1 {
		_, checkLine, err :=avutils.DoesFileContainString("export AV_HOME=~/.av", "~/.bash_profile")
		if err != nil {
			return errors.Wrapf(err,"Something went wrong reading the bash profile")
		}
		line = checkLine
		// remove the line
	}
	return nil
}