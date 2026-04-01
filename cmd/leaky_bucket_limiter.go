package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	queueSize     int
	mu            sync.Mutex
	queueCapacity int
	leakRate      int
	lastRequest   time.Time
}

func NewLeakyBucket(queueCapacity, leakRate int) *LeakyBucket {
	return &LeakyBucket{
		queueCapacity: queueCapacity,
		leakRate:      leakRate,
		lastRequest:   time.Now(),
	}
}

func (l *LeakyBucket) Allow() bool {
	l.leak()
	if l.queueSize >= l.queueCapacity {
		return false
	}

	l.mu.Lock()
	l.queueSize++
	l.mu.Unlock()

	return true
}

func (l *LeakyBucket) leak() {
	now := time.Now()
	elapsedTime := int(now.Sub(l.lastRequest) / time.Second)

	leakAmount := elapsedTime * l.leakRate

	if leakAmount > 0 {
		l.queueSize = max(0, l.queueSize-leakAmount)
		l.lastRequest = now
	}
}
