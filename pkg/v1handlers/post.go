package v1handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func HandlePost(c *gin.Context) {
	paramContentPath := c.Param("contentPath")

	if err := EnsureSecurePaths(paramContentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	fullDirPath := path.Join(MountedVolume, paramContentPath)

	dirInfo, err := os.Stat(fullDirPath)
	if err != nil || !(dirInfo.IsDir()) {
		c.String(http.StatusBadRequest, fmt.Sprintf("no such directory %s", paramContentPath))
		return
	}

	// if there is other post request,
	// it should be processed here
	handleUpload(c)
	return
}

func handleUpload(c *gin.Context) {
	paramContentPath := c.Param("contentPath")

	// upload file
	// TODO: upload directory
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	dst := path.Join(MountedVolume, paramContentPath, file.Filename)

	if _, err := os.Stat(dst); !os.IsNotExist(err) {
		errMsg := fmt.Sprintf("%s already exists", file.Filename)
		c.String(http.StatusBadRequest, errMsg)
		return
	}

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "file uploaded")
}
