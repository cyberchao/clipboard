package main

import (
	"io/ioutil"
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//new template engine
	router.HTMLRender = gintemplate.Default()

	router.GET("/", func(ctx *gin.Context) {
		data, _ := ioutil.ReadFile("db.txt")
		ctx.HTML(http.StatusOK, "new.html", gin.H{"data": string(data)})
	})
	router.POST("/modify", func(c *gin.Context) {
		data := c.PostForm("data")
		ioutil.WriteFile("db.txt", []byte(data), 0644)
		c.String(http.StatusOK, "%s", data)
	})

	router.Run(":8080")
}
