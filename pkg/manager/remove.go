package manager

import (
	"os"
)

// Remove file or directory
func Remove(fullPath string, recursive bool) (err error) {
	// check if file or directory exists
	if _, err = os.Stat(fullPath); err != nil {
		return err
	}
	if recursive {
		if err = os.RemoveAll(fullPath); err != nil {
			return err
		}
	} else {
		if err = os.Remove(fullPath); err != nil {
			return err
		}
	}
	return
}
