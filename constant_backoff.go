package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type ConstantBackoff struct {
	maxRetries int
	retries    int
	backOff    time.Duration
	waitTime   time.Duration
	lock       sync.Mutex
}

func (c *ConstantBackoff) Do() (time.Duration, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	var err error
	for c.retries < c.maxRetries {
		fmt.Printf("Calling %v :\n", url)
		_, err = http.Get(url)
		time.Sleep(c.backOff)
		c.retries++
	}
	return c.backOff * time.Duration(c.maxRetries), err
}
