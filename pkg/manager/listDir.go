/*
Package manager is the core package that contacts directly with the file system
*/
package manager

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var _ = fmt.Print // ONLY for debug

// ListDir returns all files and/or directories infomation of a specified path
// and count of all f/d rather than the count of the returned files
func ListDir(dir string, all bool, max int, ext string) (results []FileProperty, total int, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, 0, err
	}
	total = len(files)
	results = make([]FileProperty, 0, len(files))
	for _, file := range files {
		if len(results) >= max {
			break
		}
		if !all && strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if strings.HasSuffix(file.Name(), ext) {
			results = append(results, FileProperty{
				Name:    file.Name(),
				Size:    file.Size(),
				Mode:    file.Mode().String(),
				ModTime: file.ModTime().String(),
				IsDir:   file.IsDir(),
			})
		}
	}
	return
}
