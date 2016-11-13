package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	r := gin.Default()

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	api := r.Group("/api/v1")
	api.GET("/talks", passFunc)
	api.POST("/talks", newTalk)

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}
