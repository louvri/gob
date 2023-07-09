package cache

import (
	"time"
)

type Memoization interface {
	Memoize(data interface{})
	Get() interface{}
	Exists() bool
}

func New(duration time.Duration) Memoization {
	return &memoization{
		duration: duration,
	}
}

type memoization struct {
	data      interface{}
	timestamp time.Time
	duration  time.Duration
}

func (c *memoization) Memoize(data interface{}) {
	c.data = data
	c.timestamp = time.Now().UTC()
}

func (c *memoization) Get() interface{} {
	return c.data
}
func (c *memoization) Exists() bool {
	return c.data != nil && time.Now().UTC().Sub(c.timestamp) < c.duration
}
