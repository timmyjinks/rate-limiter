package main

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	funnelLoad     int
	mu             sync.Mutex
	funnelCapacity int
	leakRate       int
	lastRequest    time.Time
}

func NewLeakyBucket(queueCapacity, leakRate int) *LeakyBucket {
	return &LeakyBucket{
		funnelCapacity: queueCapacity,
		leakRate:       leakRate,
		lastRequest:    time.Now(),
	}
}

func (l *LeakyBucket) Allow() bool {
	l.leak()
	if l.funnelLoad >= l.funnelCapacity {
		return false
	}

	l.mu.Lock()
	l.funnelLoad++
	l.mu.Unlock()

	return true
}

func (l *LeakyBucket) leak() {
	now := time.Now()
	elapsedTime := int(now.Sub(l.lastRequest) / time.Second)

	leakAmount := elapsedTime * l.leakRate

	if leakAmount > 0 {
		l.funnelLoad = max(0, l.funnelLoad-leakAmount)
		l.lastRequest = now
	}
}
