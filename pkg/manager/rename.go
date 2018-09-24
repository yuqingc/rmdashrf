package manager

import (
	"errors"
	"fmt"
	"os"
)

// Rename file of directory
func Rename(oldpath, newpath string) (err error) {
	// oldpath should exist
	if _, err = os.Stat(oldpath); err != nil {
		return err
	}
	// newpath should not exist
	if _, err = os.Stat(newpath); !os.IsNotExist(err) {
		errMsg := fmt.Sprintf("%s already exists", newpath)
		return errors.New(errMsg)
	}
	err = os.Rename(oldpath, newpath)
	return err
}
