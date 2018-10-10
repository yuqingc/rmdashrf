/*
Package routes defines all routes and APIs
*/
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yuqingc/rmdashrf/pkg/v1handlers"
)

// Router is gin defalt returned engine
var Router = gin.Default()

func init() {
	// gin.SetMode(gin.ReleaseMode)
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	Router.MaxMultipartMemory = 8 << 20 // 8 MiB
	v1 := Router.Group("/api/v1")
	{
		v1.GET("/default/*contentPath", v1handlers.HandleGet)
		v1.PUT("/default/*contentPath", v1handlers.HandlePut)
		v1.DELETE("/default/*contentPath", v1handlers.HandleDelete)
		v1.PATCH("/default/*contentPath", v1handlers.HandlePatch)
		v1.POST("/default/*contentPath", v1handlers.HandlePost)
	}
}
