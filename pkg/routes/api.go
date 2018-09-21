/*
	routes package defines all routes and APIs
*/
package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuqingc/rmdashrf/pkg/v1handlers"
)

var Router = gin.Default()

func init() {
	fmt.Println("Init function is executed in package `routes`")
	v1 := Router.Group("/api/v1")
	{
		v1.GET("/test", v1handlers.Test)
	}
}
