package v1handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleGet handles all GET request of
// api: /default/*contentPath
// params: restype=<file|directory>&comp={list|metadata}
func HandleGet(c *gin.Context) {
	restype := c.Query("restype")
	comp := c.Query("comp")
	if restype == "directory" && comp == "list" {
		GetList(c)
		return
	}
	c.String(http.StatusBadRequest, BadRequestErrMsg)
}
