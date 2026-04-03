package main

import (
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	mu         sync.Mutex
	requests   int
	threshold  int
	window     time.Time
	windowSize time.Duration
}

func NewFixedWindowRateLimiter(threshold int, windowSize time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		threshold:  threshold,
		windowSize: windowSize,
		window:     time.Now(),
	}
}

func (r *FixedWindowRateLimiter) Allow() bool {
	r.refreshWindow()
	if r.requests >= r.threshold {
		return false
	}
	r.mu.Lock()
	r.requests++
	r.mu.Unlock()

	return true
}

func (r *FixedWindowRateLimiter) refreshWindow() {
	now := time.Now()

	if now.Sub(r.window) > r.windowSize {
		r.window = now
		r.requests = 0
		return
	}
}
