package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yuqingc/rmdashrf/pkg/routes"
	"github.com/yuqingc/rmdashrf/pkg/v1handlers"
)

func main() {
	if v1handlers.RequestVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	if v1handlers.MountedVolume == "" {
		fmt.Println("volume is not specified")
		os.Exit(1)
	}
	if !strings.HasPrefix(v1handlers.MountedVolume, "/") {
		fmt.Println("volume must be an absolute path")
		os.Exit(1)
	}
	fmt.Println("RMDASHRF is running at :" + v1handlers.Port)
	router := routes.Router
	router.Run(":" + v1handlers.Port)
}
