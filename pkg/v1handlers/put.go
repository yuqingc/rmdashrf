package v1handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandlePut(c *gin.Context) {
	restype := c.Query("restype")
	if restype == "" {
		CreateFile(c)
		return
	}
	c.String(http.StatusBadRequest, BadRequestErrMsg)
}

func CreateFile(c *gin.Context) {
	contentPath := c.Param("contentPath")
	if err := CheckContentPath(contentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid path \"%s\"\n", contentPath))
		return
	}
	fullFilePath := path.Join(MountDir, contentPath)
	if _, err := os.Stat(fullFilePath); !os.IsNotExist(err) {
		log.Printf("file %s already exists", fullFilePath)
		c.String(http.StatusBadRequest, "file already exists")
		return
	}

	// Do not use os.CreateFile in case there is a race
	// where a new file is created at the same time,
	// and the file will be overwritten.
	// There is a chance that at a file with same name is created meanwhile.
	// New file is not created in this case but no error is thrown.
	// This is a bug but it's almost impossible to happen.
	createdFile, err := os.OpenFile(fullFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		var errMsg = "creating file failed"
		log.Println(errMsg, err)
		if os.IsNotExist(err) {
			errMsg = "no such file or directory"
		}
		c.String(http.StatusBadRequest, errMsg)
		return
	}
	defer createdFile.Close()
	c.String(http.StatusCreated, fmt.Sprintf("%s is created", filepath.Base(createdFile.Name())))
}
