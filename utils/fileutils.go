package utils

import (
	"errors"
	"fmt"
	"os"
)

// PathExists returns true if the given path exists, and false if it does not.
// If it cannot determine whether or not a path exists (e.g., because of a
// permissions problem), it will log an error and exit.
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, fmt.Errorf("unable to check for %s: %w", path, err)
	}

	return true, nil
}
