package main

import (
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"
)

type ExponentialBackoff struct {
	maxRetries int
	retries    int
	backOff    int
	waitTime   time.Duration
	lock       sync.Mutex
}

func (e *ExponentialBackoff) Do() (time.Duration, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	var err error
	for e.retries < e.maxRetries {
		fmt.Printf("Calling %v :\n", url)
		_, err = http.Get(url)
		e.waitTime += time.Duration(math.Pow(float64(e.backOff), float64(e.retries)))
		time.Sleep(time.Duration(math.Pow(float64(e.backOff), float64(e.retries))))
		e.retries++
	}
	return e.waitTime, err
}
