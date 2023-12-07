package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

const (
	url                = "http://localhost:8080/read-file"
	linearBackoff      = "LinearBackoff"
	exponentialBackoff = "ExponentialBackoff"
	constantBackoff    = "ConstantBackoff"
)

type response struct {
	LinearBackoff      time.Duration
	ExponentialBackoff time.Duration
	ConstantBackoff    time.Duration
}

func Operation(c *gin.Context) {
	maxRetries, err := strconv.Atoi(c.Request.Header.Get("max-retries"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "maxRetries header value is expected to be integer")
		return
	}

	response, err := Do(maxRetries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error occurred while calculating")
		return
	}

	c.JSON(http.StatusOK, gin.H{linearBackoff: response.LinearBackoff, exponentialBackoff: response.ExponentialBackoff, constantBackoff: response.ConstantBackoff})
}
