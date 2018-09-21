package v1handlers

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

var _ = fmt.Print // this statement is used only when DEBUGGING

func Test(c *gin.Context) {
	files, err := ioutil.ReadDir("/home/matt")
	if err != nil {
		log.Fatal(err)
	}
	var fileNames []string
	for _, file := range files {
		// fmt.Println("file is", file.Name())
		fileNames = append(fileNames, file.Name())
	}
	c.JSON(200, gin.H{
		"files": fileNames,
	})
}
