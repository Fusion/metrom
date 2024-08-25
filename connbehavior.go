package main

import "time"

type ConnBehavior struct {
	pause   time.Duration
	timeout time.Duration
	retries int
	ttl     int
}
