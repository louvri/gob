package cache

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew_ReturnsNonNil(t *testing.T) {
	m := New(time.Second)
	assert.NotNil(t, m)
}

func TestMemoize_And_Get(t *testing.T) {
	m := New(time.Minute)
	m.Memoize("hello")
	assert.Equal(t, "hello", m.Get())
}

func TestMemoize_OverwritesPrevious(t *testing.T) {
	m := New(time.Minute)
	m.Memoize("first")
	m.Memoize("second")
	assert.Equal(t, "second", m.Get())
}

func TestGet_ReturnsNilWhenEmpty(t *testing.T) {
	m := New(time.Minute)
	assert.Nil(t, m.Get())
}

func TestExists_TrueWhenFresh(t *testing.T) {
	m := New(time.Minute)
	m.Memoize("data")
	assert.True(t, m.Exists())
}

func TestExists_FalseWhenEmpty(t *testing.T) {
	m := New(time.Minute)
	assert.False(t, m.Exists())
}

func TestExists_FalseWhenExpired(t *testing.T) {
	m := New(10 * time.Millisecond)
	m.Memoize("data")
	time.Sleep(20 * time.Millisecond)
	assert.False(t, m.Exists())
}

func TestExists_TrueJustBeforeExpiry(t *testing.T) {
	m := New(time.Second)
	m.Memoize("data")
	assert.True(t, m.Exists())
}

func TestMemoize_StructValue(t *testing.T) {
	type item struct {
		Name string
		ID   int
	}
	m := New(time.Minute)
	val := item{Name: "test", ID: 42}
	m.Memoize(val)
	got := m.Get().(item)
	assert.Equal(t, val, got)
}

func TestConcurrentAccess(t *testing.T) {
	m := New(time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func(v int) {
			defer wg.Done()
			m.Memoize(v)
		}(i)
		go func() {
			defer wg.Done()
			m.Get()
			m.Exists()
		}()
	}
	wg.Wait()
}
