package main

import "github.com/gin-gonic/gin"

func main() {
	rr := gin.Default()
	rr.GET("/ping", func(cc *gin.Context) {
		cc.JSON(200, gin.H{
			"message": "pong",
		})
	})
	rr.Run()
}
