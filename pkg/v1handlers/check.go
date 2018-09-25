package v1handlers

import (
	"fmt"
	"strings"
)

// EnsureSecurePaths stops checking and returns an error when a path is insecure
func EnsureSecurePaths(contentPaths ...string) error {
	for _, contentPath := range contentPaths {
		if strings.Contains(contentPath, "..") {
			err := fmt.Errorf("invalid path `%s`: relative parent path is not allowed", contentPath)
			return err
		}
	}
	return nil
}
