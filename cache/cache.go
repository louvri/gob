package cache

import (
	"sync"
	"time"
)

type Memoization interface {
	Memoize(data any)
	Get() any
	Exists() bool
}

func New(duration time.Duration) Memoization {
	return &memoization{
		duration: duration,
	}
}

type memoization struct {
	mu        sync.RWMutex
	data      any
	timestamp time.Time
	duration  time.Duration
}

func (c *memoization) Memoize(data any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = data
	c.timestamp = time.Now().UTC()
}

func (c *memoization) Get() any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data
}

func (c *memoization) Exists() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data != nil && time.Now().UTC().Sub(c.timestamp) < c.duration
}
