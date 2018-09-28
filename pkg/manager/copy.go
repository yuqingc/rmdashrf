package manager

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// CopyDir copies a directory and its children recursively
// Terminates when it meets an error (no fall back)
func CopyDir(src, dst string) (err error) {
	if strings.HasPrefix(dst, src) {
		return fmt.Errorf("Operation not allowed. Copying directory aborted. Trying to copy a path into its child directory will cause infinite loop.")
	}
	// src should exist and must be a directory
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sfi.IsDir() {
		return fmt.Errorf("src path %s is not a directory", src)
	}

	// dst should not exist
	if _, err = os.Stat(dst); !os.IsNotExist(err) {
		return fmt.Errorf("%s, file or directory already exists", dst)
	}

	if err = walkAndCopy(src, dst); err != nil {
		return err
	}

	return
}

// walkDir copies child contents of a directory recursively
func walkAndCopy(dir, counterpartDst string) (err error) {
	if err = os.Mkdir(counterpartDst, 0755); err != nil {
		return err
	}
	contents, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, content := range contents {
		contentPath := path.Join(dir, content.Name())
		counterpartPath := path.Join(counterpartDst, content.Name())
		if content.IsDir() {
			if err = walkAndCopy(contentPath, counterpartPath); err != nil {
				return err
			}
		} else {
			if err = CopyFile(contentPath, counterpartPath); err != nil {
				return err
			}
		}
	}
	return
}

// ref: stackoverflow: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherwise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	err = ensureBeforeCopy(src, dst)
	if err != nil {
		return err
	}

	if err = os.Link(src, dst); err != nil {
		err = copyFileContents(src, dst)
		if err != nil {
			return
		}
	}
	return
}

// ensureBeforeCopy ensures copying process for copying FILE
func ensureBeforeCopy(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !(sfi.Mode().IsRegular()) {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	if _, err := os.Stat(dst); !os.IsNotExist(err) {
		return fmt.Errorf("CopyFile: file already exists")
	}
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
