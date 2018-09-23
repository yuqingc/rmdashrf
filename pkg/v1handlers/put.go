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
	createdFile, err := manager.CreateFile(fullFilePath)
	if err != nil {
		log.Println(err)
		var errMsg = "file already exits"
		if os.IsNotExist(err) {
			errMsg = "no such file or directory"
		}
		c.String(http.StatusBadRequest, errMsg)
		return
	}
	defer createdFile.Close()
	c.String(http.StatusCreated, fmt.Sprintf("%s is created", filepath.Base(createdFile.Name())))
}
