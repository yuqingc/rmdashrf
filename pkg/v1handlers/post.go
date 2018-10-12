package v1handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// HandlePost handles all POST requests
func HandlePost(c *gin.Context) {
	paramPath := c.Param("defaultSlashContentPath")

	if err := EnsureSecurePaths(paramPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// hack
	// all params will not be in the path in the next version (probably)
	/***** UGLY CODE *****/
	if !(paramPath == "/default" || strings.HasPrefix(paramPath, "/default/")) {
		log.Printf("invalid path: %s does not start with default", paramPath)
		c.Status(http.StatusNotFound)
		return
	}
	/***** UGLY CODE *****/

	contentPath := getContentPath(paramPath)

	fullDirPath := path.Join(MountedVolume, contentPath)

	dirInfo, err := os.Stat(fullDirPath)
	if err != nil || !(dirInfo.IsDir()) {
		c.String(http.StatusBadRequest, fmt.Sprintf("no such directory %s", contentPath))
		return
	}

	// if there is other post request,
	// it should be processed here
	handleUpload(c)
	return
}

func handleUpload(c *gin.Context) {
	paramPath := c.Param("defaultSlashContentPath")

	paramContentPath := getContentPath(paramPath)

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

// getContentPath trims the `/default` prefix and returns the content path starting with a slash
func getContentPath(pathWithDefault string) string {
	var fullPath = pathWithDefault
	if fullPath == "/default" {
		fullPath = "/default/"
	}
	return strings.TrimPrefix(fullPath, "/default")
}
