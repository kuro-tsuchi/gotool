package main

import "github.com/gin-gonic/gin"

logrus

func main() {
	//gin http.Server
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")

}
