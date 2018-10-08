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
// and count of all f/d rather than the count of the returned files.
// Caveats: the size of directory is solid, which is the size of the directory
// itself, not including what is inside the directory.
// Calculating the total size of a certain directory needs the program
// to calculate the sizes of all its sub contents recursively, which costs too much time and memory.
// It is not necessary to make this calculation.
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
