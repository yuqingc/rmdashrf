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
	v1 := Router.Group("/api/v1")
	{
		v1.GET("/default/*contentPath", v1handlers.HandleGet)
		v1.PUT("/default/*contentPath", v1handlers.HandlePut)
		v1.DELETE("/default/*contentPath", v1handlers.HandleDelete)
		v1.PATCH("/default/*contentPath", v1handlers.HandlePatch)
	}
}
