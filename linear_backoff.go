package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LinearBackoff struct {
	maxRetries int
	retries    int
	backOff    int
	waitTime   time.Duration
	lock       sync.Mutex
}

func (l *LinearBackoff) Do() (time.Duration, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	var err error
	for l.retries < l.maxRetries {
		fmt.Printf("Calling %v :\n", url)
		_, err = http.Get(url)
		l.waitTime += time.Duration(l.backOff + l.retries)
		time.Sleep(time.Duration(l.backOff + l.retries))
		l.retries++
	}
	return l.waitTime, err
}
