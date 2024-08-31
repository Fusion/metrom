package net

import "time"

type ConnBehavior struct {
	pause     time.Duration
	timeout   time.Duration
	quiescent time.Duration
	retries   int
	ttl       int
}
