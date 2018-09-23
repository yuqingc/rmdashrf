package v1handlers

import (
	"errors"
	"strings"
)

// CheckContentPath returns an error is the path is not allowed
func CheckContentPath(contentPath string) (err error) {
	if strings.Contains(contentPath, "..") {
		err = errors.New("relative parent path is not allowed")
	}
	return
}
