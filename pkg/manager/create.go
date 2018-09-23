package manager

import (
	"errors"
	"fmt"
	"os"
)

// Createfile creates a new file
// If the file is already exists of the directory does not exists,
// it will return an error
func CreateFile(fullFilePath string) (*os.File, error) {
	if _, err := os.Stat(fullFilePath); !os.IsNotExist(err) {
		errMsg := fmt.Sprintf("%s already exists", fullFilePath)
		return nil, errors.New(errMsg)
	}

	// Do not use os.CreateFile in case there is a race
	// where a new file is created at the same time,
	// and the file will be overwritten.
	// There is a chance that at a file with same name is created meanwhile.
	// New file is not created in this case but no error is thrown.
	// This is a bug but it's almost impossible to happen.
	createdFile, err := os.OpenFile(fullFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return createdFile, nil
}

// CreateDir creates new a new directory
// It creates parents directory if needed when `parents` is `true`
// It is similar to `mkdir -p`
func CreateDir(fullDirPath string, parents bool) error {
	if _, err := os.Stat(fullDirPath); !os.IsNotExist(err) {
		errMsg := fmt.Sprintf("%s already exists", fullDirPath)
		return errors.New(errMsg)
	}

	if parents {
		if err := os.MkdirAll(fullDirPath, 0755); err != nil {
			return err
		}
	} else {
		if err := os.Mkdir(fullDirPath, 0755); err != nil {
			return err
		}
	}
	return nil
}
