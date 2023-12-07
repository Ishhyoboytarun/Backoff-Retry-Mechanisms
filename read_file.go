package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

const filePath = "file"

func ReadFile(c *gin.Context) {

	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read the file"})
		return
	}

	defer file.Close()
	contents, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read the file"})
		return
	}

	c.JSON(http.StatusOK, string(contents))
}
