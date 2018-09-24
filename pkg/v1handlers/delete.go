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

// HandleDelete handles DELETE method
// remove a file or directory
func HandleDelete(c *gin.Context) {
	contentPath := c.Param("contentPath")
	if err := CheckContentPath(contentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid path \"%s\"\n", contentPath))
		return
	}

	fullPath := path.Join(MountDir, contentPath)
	paramRecuresive := c.Query("recursive")
	recursive := false
	if paramRecuresive == "true" {
		recursive = true
	}
	if err := manager.Remove(fullPath, recursive); err != nil {
		log.Println(err)
		var ErrMsg = "Deleting file failed"
		if headache, ok := err.(*os.PathError); ok {
			ErrMsg = headache.Err.Error()
		}
		c.String(http.StatusBadRequest, ErrMsg)
		return
	}
	c.String(http.StatusAccepted, "Deleted")
}
