package main

import (
	"time"
)

type RateLimiter struct {
	tokens         int
	capacity       int
	refillInterval int
	lastRefillTime time.Time
}

func NewRateLimiter(capacity int, refillInterval int) *RateLimiter {
	rl := &RateLimiter{
		capacity:       capacity,
		refillInterval: refillInterval,
	}
	return rl
}

func (r *RateLimiter) Allow() bool {
	r.refillTokens()
	if r.tokens > 0 {
		r.tokens--
		return true
	}
	return false
}

func (r *RateLimiter) refillTokens() {
	now := time.Now()
	elapsedTime := int(now.Sub(r.lastRefillTime) / time.Second)
	refillTokens := elapsedTime / r.refillInterval

	if refillTokens > 0 {
		r.tokens = min(r.capacity, r.tokens+refillTokens)
		r.lastRefillTime = now
	}
}
