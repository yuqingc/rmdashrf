package v1handlers

import (
	"archive/zip"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

// DownloadFile serves static file
func DownloadFile(c *gin.Context) {
	contentPath := c.Param("contentPath")
	if err := EnsureSecurePaths(contentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	fullFilePath := path.Join(MountDir, contentPath)
	fileInfo, err := os.Stat(fullFilePath)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	if !(fileInfo.Mode().IsRegular()) {
		// cannot download non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		c.String(http.StatusBadRequest, "not a regular file")
		return
	}
	c.File(fullFilePath)
}

// DownloadDirAsZip zip folder and response
// Do not create a file in local disk
func DownloadDirAsZip(c *gin.Context) {
	contentPath := c.Param("contentPath")
	if err := EnsureSecurePaths(contentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	fullDirPath := path.Join(MountDir, contentPath)
	dirInfo, err := os.Stat(fullDirPath)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	if !(dirInfo.IsDir()) {
		c.String(http.StatusBadRequest, "not a directory")
		return
	}

	// buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(c.Writer)
	// const testFileName = "/home/matt/Projects/github.com/yuqingc/data/package.json"
	// testFileData, err := ioutil.ReadFile(testFileName)
	const testFileData = "hello world\n"
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	f, err := zipWriter.Create("haha/hahaha.txt")
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	_, err = f.Write([]byte(testFileData))
	_, err = f.Write([]byte(testFileData))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	// TODO create and write zip to stream
	// TODO: open download window, maybe not file download
	c.Header("Content-Disposition", "attachment; filename=aaa.zip")
	c.Header("Content-Type", "application/x-zip-compressed")
	// This will be a for loop to write zip file
	// generate and write
	// c.Writer.Write([]byte("hello"))
	// io.Copy(c.Writer, buf)
	err = zipWriter.Close()
	if err != nil {
		log.Println(err)
	}

}
