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
	contents, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, 0, err
	}
	total = len(contents)
	results = make([]FileProperty, 0, len(contents))
	for _, content := range contents {
		if len(results) >= max {
			break
		}
		if !all && strings.HasPrefix(content.Name(), ".") {
			continue
		}
		if strings.HasSuffix(content.Name(), ext) {
			results = append(results, FileProperty{
				Name:    content.Name(),
				Size:    content.Size(),
				Mode:    content.Mode().String(),
				ModTime: content.ModTime().String(),
				IsDir:   content.IsDir(),
			})
		}
	}
	return
}
