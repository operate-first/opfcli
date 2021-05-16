package utils

import (
	"os/exec"
	"strings"
)

// FindRepoDir returns the path to the top directory of a git
// repository. If it is run outside of a git repository, or otherwise
// cannot determine the top directory, it returns an error.
func FindRepoDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
