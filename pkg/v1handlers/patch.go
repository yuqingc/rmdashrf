package v1handlers

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/yuqingc/rmdashrf/pkg/manager"
)

// HandlePatch handles all PATCH requests
// rename,
func HandlePatch(c *gin.Context) {
	action := c.Query("action")

	if action == "rename" {
		Rename(c)
		return
	}
	c.String(http.StatusBadRequest, "invalid request param action")
}

// Rename handles all rename requests from `HandlePatch`
func Rename(c *gin.Context) {
	paramContentPath := c.Param("contentPath")
	paramTo := c.Query("to")

	// This function might be unnecessary
	if paramTo == "" {
		c.String(http.StatusBadRequest, "param `to` is required")
		return
	}

	if err := EnsureSecurePaths(paramContentPath, paramTo); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	oldPath := path.Join(MountDir, paramContentPath)
	newPath := path.Join(MountDir, paramTo)
	if err := manager.Rename(oldPath, newPath); err != nil {
		log.Println("rename failed:", err)
		var ErrMsg = "rename failed: old path should exists and new path should not exist"
		if os.IsNotExist(err) {
			ErrMsg = "file or directory does not exist"
		}
		c.String(http.StatusBadRequest, ErrMsg)
		return
	}
	c.String(http.StatusNoContent, "")
}
