package store

import (
	"os"
	"os/exec"
)

// GetExecPath gets the path of the executable on the filesystem. This can be unreliable as
// it expects argv[0] to be the name of the executable may not not always the case.
func GetExecPath() (string, error) {
	return exec.LookPath(os.Args[0])
}
