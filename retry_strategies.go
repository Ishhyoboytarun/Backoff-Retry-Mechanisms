package main

import (
	"fmt"
	"time"
)

const (
	explicitFailing          = 2
	linearBackoffFactor      = 1
	exponentialBackoffFactor = 2
	constantBackoffFactor    = 3
)

func Do(maxRetries int) (response, error) {

	var strategy1, strategy2, strategy3 RetryStrategy
	strategy1 = &LinearBackoff{maxRetries: maxRetries, backOff: linearBackoffFactor}
	strategy2 = &ExponentialBackoff{maxRetries: maxRetries, backOff: exponentialBackoffFactor}
	strategy3 = &ConstantBackoff{maxRetries: maxRetries, backOff: constantBackoffFactor}

	var response1, response2, response3 time.Duration
	for i := 0; i < explicitFailing; i++ {
		fmt.Printf("Retrying %v time :\n", i)
		response1, _ = strategy1.Do()
		response2, _ = strategy2.Do()
		response3, _ = strategy3.Do()
	}
	return response{
		LinearBackoff:      response1,
		ExponentialBackoff: response2,
		ConstantBackoff:    response3,
	}, nil
}

type RetryStrategy interface {
	Do() (time.Duration, error)
}
