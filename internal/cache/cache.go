package cache

import (
	"context"
	"sync"
)

type Cache[T any] struct {
	mu    sync.Mutex
	ready bool
	val   T
}

func (c *Cache[T]) Get(
	ctx context.Context,
	compute func(context.Context) (T, error),
) (T, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ready {
		v := c.val
		return v, nil
	}
	v, err := compute(ctx)
	if err != nil {
		var zero T
		return zero, err
	}
	c.val = v
	c.ready = true
	return c.val, nil
}
