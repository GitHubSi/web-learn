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
		c.HTML(http.StatusOK, "default.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")
}
