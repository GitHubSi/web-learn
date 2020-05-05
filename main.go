package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/resource", "./resource")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"templateTitle": "card-list.tmpl",
		})
	})

	router.GET("/get-request", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"templateTitle": "get-request.tmpl",
		})
	})

	router.POST("/set-progress", func(context *gin.Context) {
		type Request struct {
			Progress string `json:"progress" form:"progress"`
		}

		request := &Request{}
		if err := context.Bind(request); err != nil {
		}

		context.JSON(http.StatusOK, gin.H{
			"progress": request.Progress,
		})
		return
	})

	router.Run(":8080")
}
