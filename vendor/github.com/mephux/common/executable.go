package common

import "os/exec"

func HasExecutable(binPath string) (string, error) {
	return exec.LookPath(binPath)
}
