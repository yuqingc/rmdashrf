/*
Package routes defines all routes and APIs
*/
package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yuqingc/rmdashrf/pkg/v1handlers"
)

// Router is gin defalt returned engine
var Router = gin.Default()

func init() {
	log.Println("loading routes...")
	v1 := Router.Group("/api/v1")
	{
		v1.GET("/default/*contentPath", v1handlers.HandleGet)
	}
}
