package v1handlers

import (
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

	// TODO create and write zip to stream

}
