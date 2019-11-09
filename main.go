package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Fatalln(router.Run(":8080"))
}
