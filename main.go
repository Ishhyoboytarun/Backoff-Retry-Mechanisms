package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	//read-file api is assumed to be an external api
	router.GET("/read-file", ReadFile)

	router.GET("/operation", Operation)
	router.Run("localhost:8080")
}
