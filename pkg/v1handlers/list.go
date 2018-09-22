/*
package v1handlers functions handling api/v1 Gin routes
*/
package v1handlers

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuqingc/rmdashrf/pkg/manager"
)

var _ = fmt.Print // ONLY used for debugging

// TODO: checks if mount dir exists. Create it if not
// TODO: mountpath must be an absolute path
const MountDir = "/home/matt/Projects/github.com/yuqingc/data"
const MaxListResults = 5000

type Metadata struct {
	Total int `json:"total"`
}

type ListResponse struct {
	Metadata Metadata               `json:"metadata"`
	Items    []manager.FileProperty `json:"items"`
}

// GetList returns all files and directories of specified path
// /default/:path
func GetList(c *gin.Context) {
	contentPath := c.Param("contentPath")
	if strings.Contains(contentPath, "..") {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid path \"%s\"\n", contentPath))
		log.Printf("querying path \"%s\" is denied\n", contentPath)
		return
	}
	paramAll := c.Query("all")
	paramMaxresults := c.Query("maxresults")
	paramExtension := c.Query("extension")

	dir := path.Join(MountDir, contentPath)
	all := paramAll == "true"
	var maxresults = MaxListResults
	if paramMaxresults != "" {
		maxresults, err := strconv.Atoi(paramMaxresults)
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
