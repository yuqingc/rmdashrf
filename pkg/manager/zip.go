package manager

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// ZipDir archives a directory to a zip file.
// `dirPath` is the source path of the directory to be zipped.
// `w` is the writer you want to write the zip bytes to
func ZipDir(dirPath string, w io.Writer) (err error) {
	zipWriter := zip.NewWriter(w)
	err = walkAndZip(dirPath, "", zipWriter)
	if err != nil {
		return
	}
	err = zipWriter.Close()
	return
}

// walkAndZip walk through directory recursively.
// `dst` is the path where you put your file,
// which is relative to the root of src.
func walkAndZip(src string, dst string, zipWriter *zip.Writer) (err error) {
	contents, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, content := range contents {
		contentName := content.Name()
		contentPath := path.Join(src, contentName)
		if content.IsDir() {
			if err = walkAndZip(contentPath, dst+contentName+"/", zipWriter); err != nil {
				return err
			}
		} else {
			buf := make([]byte, 1024)
			originFile, err := os.Open(contentPath)
			if err != nil {
				return err
			}
			defer originFile.Close()
			fileWriter, err := zipWriter.Create(dst + contentName) // for a file
			for {                                                  // read and write
				nr, err := originFile.Read(buf)
				if nr < 0 {
					return err
				} else if nr == 0 { // EOF
					break // Do not use switch-case or break won't go out of the loop
				} else {
					nw, err := fileWriter.Write(buf[0:nr])
					if nw != nr {
						return fmt.Errorf("write buff length is not equal to read buff")
					}
					if err != nil && err != io.EOF {
						return err
					}
				}
			}
		}
	}
	return
}
