package v1handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/yuqingc/rmdashrf/pkg/manager"
)

// DownloadFile serves static file
func DownloadFile(c *gin.Context) {
	contentPath := c.Param("contentPath")
	if err := EnsureSecurePaths(contentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	fullFilePath := path.Join(MountedVolume, contentPath)
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
	fullDirPath := path.Join(MountedVolume, contentPath)
	dirInfo, err := os.Stat(fullDirPath)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	if !(dirInfo.IsDir()) {
		c.String(http.StatusBadRequest, "not a directory")
		return
	}

	var headerContentDisposition = fmt.Sprintf("attachment; filename=%s.zip", filepath.Base(contentPath))
	c.Header("Content-Disposition", headerContentDisposition)
	c.Header("Content-Type", "application/x-zip-compressed")
	err = manager.ZipDir(fullDirPath, c.Writer)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "fail to zip directory"+contentPath)
		return
	}
}
