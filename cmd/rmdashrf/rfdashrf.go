package main

import (
	"fmt"

	"github.com/yuqingc/rmdashrf/pkg/routes"
)

func main() {
	const port = "8080"
	fmt.Println("RMDASHRF is running at :" + port)
	router := routes.Router
	router.Run(":" + port)
}
