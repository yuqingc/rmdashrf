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

// HandlePut handles all PUT request
// including creating files and directories,
// copying files and directories
func HandlePut(c *gin.Context) {
	action := c.Query("action")
	restype := c.Query("restype")
	if action == "create" && restype == "" {
		CreateFile(c)
		return
	}
	if action == "create" && restype == "directory" {
		CreateDir(c)
		return
	}
	if action == "copy" && restype == "" {
		CopyFile(c)
		return
	}
	c.String(http.StatusBadRequest, BadRequestErrMsg)
}

// TODO: split these functions into different files

// CreateFile handles the request for creating a new file
func CreateFile(c *gin.Context) {
	paramContentPath := c.Param("contentPath")
	if err := EnsureSecurePaths(paramContentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid path \"%s\"\n", paramContentPath))
		return
	}

	fullFilePath := path.Join(MountDir, paramContentPath)
	createdFile, err := manager.CreateFile(fullFilePath)
	if err != nil {
		log.Println(err)
		var errMsg = "file or directory already exits"
		if os.IsNotExist(err) {
			errMsg = "no such file or directory"
		}
		c.String(http.StatusBadRequest, errMsg)
		return
	}
	defer createdFile.Close()
	c.String(http.StatusCreated, fmt.Sprintf("%s is created", filepath.Base(createdFile.Name())))
}

// CreateDir handles the request for creating a new directory
func CreateDir(c *gin.Context) {
	paramContentPath := c.Param("contentPath")
	if err := EnsureSecurePaths(paramContentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid path \"%s\"\n", paramContentPath))
		return
	}

	parents := false
	if paramParents := c.Query("parents"); paramParents == "true" {
		parents = true
	}

	fullDirPath := path.Join(MountDir, paramContentPath)
	if err := manager.CreateDir(fullDirPath, parents); err != nil {
		log.Println(err)
		var errMsg = "file or directory already exits"
		if os.IsNotExist(err) {
			errMsg = "no such file or directory"
		}
		c.String(http.StatusBadRequest, errMsg)
		return
	}
	c.String(http.StatusCreated, "directory is created")
}

// CopyFile handles the request for copying a file
func CopyFile(c *gin.Context) {
	paramContentPath := c.Param("contentPath")
	paramFrom := c.Query("from")

	if paramFrom == "" {
		c.String(http.StatusBadRequest, "param `from` is required")
		return
	}

	if err := EnsureSecurePaths(paramContentPath, paramFrom); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	fullFilePath := path.Join(MountDir, paramContentPath)
	fullFromPath := path.Join(MountDir, paramFrom)

	if err := manager.CopyFile(fullFromPath, fullFilePath); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusCreated, "Copied")
}
