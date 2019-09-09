package common

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetGOPATHs returns all paths in GOPATH variable.
func GetGOPATHs() []string {
	gopath := os.Getenv("GOPATH")
	var paths []string
	if runtime.GOOS == "windows" {
		gopath = strings.Replace(gopath, "\\", "/", -1)
		paths = strings.Split(gopath, ";")
	} else {
		paths = strings.Split(gopath, ":")
	}
	return paths
}

// GetSrcPath returns app. source code path.
// It only works when you have src. folder in GOPATH,
// it returns error not able to locate source folder path.
func GetSrcPath(importPath string) (appPath string, err error) {
	paths := GetGOPATHs()
	for _, p := range paths {
		if IsExist(p + "/src/" + importPath + "/") {
			appPath = p + "/src/" + importPath + "/"
			break
		}
	}

	if len(appPath) == 0 {
		return "", errors.New("Unable to locate source folder path")
	}

	appPath = filepath.Dir(appPath) + "/"
	if runtime.GOOS == "windows" {
		// Replace all '\' to '/'.
		appPath = strings.Replace(appPath, "\\", "/", -1)
	}

	return appPath, nil
}

// HomeDir returns path of '~'(in Linux) on Windows,
// it returns error when the variable does not exist.
func HomeDir() (home string, err error) {
	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
		if len(home) == 0 {
			home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		}
	} else {
		home = os.Getenv("HOME")
	}

	if len(home) == 0 {
		return "", errors.New("Cannot specify home directory because it's empty")
	}

	return home, nil
}
