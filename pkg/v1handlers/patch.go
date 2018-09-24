package v1handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/yuqingc/rmdashrf/pkg/manager"
)

func HandlePatch(c *gin.Context) {
	action := c.Query("action")

	if action == "rename" {
		Rename(c)
		return
	}
	c.String(http.StatusBadRequest, "invalid request param action")
}

func Rename(c *gin.Context) {
	contentPath := c.Param("contentPath")
	to := c.Query("to")

	// This function might be unnecessary
	if to == "" {
		c.String(http.StatusBadRequest, "param `to` is required")
		return
	}

	if err := CheckContentPath(contentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid path \"%s\"\n", contentPath))
		return
	}
	fullOldPath := path.Join(MountDir, contentPath)
	fullNewPath := path.Join(MountDir, to)
	if err := manager.Rename(fullOldPath, fullNewPath); err != nil {
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
