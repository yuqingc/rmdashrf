package v1handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var _ = fmt.Print // this statement is used only when DEBUGGING

const mountDir = "/home/matt/Projects/github.com/yuqingc/data"

func Test(c *gin.Context) {
	queryPath := c.DefaultQuery("path", "/") // this will be passed as the query argument
	if strings.Contains(queryPath, "..") {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid path \"%s\"\n", queryPath))
		log.Printf("test: querying path \"%s\" is denied\n", queryPath)
		return
	}
	dir := path.Join(mountDir, queryPath)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		c.String(http.StatusNotFound, "no such directory: "+queryPath)
		log.Println(err)
		return
	}
	var fileArr []map[string]interface{}
	for _, file := range files {
		// fmt.Println("file is", file.Name())
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		fileMap := make(map[string]interface{})
		fileMap["name"] = file.Name()
		fileMap["size"] = file.Size()
		fileMap["mode"] = fmt.Sprintf("%v", file.Mode())
		fileMap["modTime"] = file.ModTime()
		fileMap["isDir"] = file.IsDir()
		fileArr = append(fileArr, fileMap)
	}
	c.JSON(200, gin.H{
		"files": fileArr,
	})
}
