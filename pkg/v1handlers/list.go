/*
Package v1handlers functions handling api/v1 Gin routes
*/
package v1handlers

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuqingc/rmdashrf/pkg/manager"
)

var _ = fmt.Print // ONLY for debug

// Metadata of result for listing content
type Metadata struct {
	Total int `json:"total"`
}

// ListResponse is type of result for listing content
type ListResponse struct {
	Metadata Metadata               `json:"metadata"`
	Items    []manager.FileProperty `json:"items"`
}

// GetList returns all files and directories of specified path
// api: /default/:path
func GetList(c *gin.Context) {
	var err error
	contentPath := c.Param("contentPath")
	if err = EnsureSecurePaths(contentPath); err != nil {
		log.Println("checkpath failed:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	paramAll := c.Query("all")
	paramMaxresults := c.Query("maxresults")
	paramExtension := c.Query("extension")

	dir := path.Join(MountedVolume, contentPath)
	all := paramAll == "true"
	var maxresults = MaxListResults
	if paramMaxresults != "" {
		maxresults, err = strconv.Atoi(paramMaxresults)
		if err != nil {
			log.Println(err)
			c.String(http.StatusBadRequest, "invalid maxresult")
			return
		}
		if maxresults > MaxListResults {
			maxresults = MaxListResults
		}
	}

	listedFiles, total, err := manager.ListDir(dir, all, maxresults, paramExtension)
	if err != nil {
		log.Println(err)
		c.String(http.StatusNotFound, fmt.Sprintf("no such directory: %s", contentPath))
		return
	}
	result := ListResponse{
		Metadata: Metadata{Total: total},
		Items:    listedFiles,
	}
	c.JSON(http.StatusOK, result)
}
