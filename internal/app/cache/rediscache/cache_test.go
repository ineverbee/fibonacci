package rediscache_test

import (
	"testing"
	"time"

	"github.com/ineverbee/fibonacci/internal/app/cache/rediscache"
	"github.com/stretchr/testify/assert"
)

func TestCache_Set(t *testing.T) {
	cache, teardown := rediscache.TestCache(t)
	defer teardown()
	err := cache.Set(0, 0, time.Minute)
	assert.NoError(t, err)
}

func TestCache_Get(t *testing.T) {
	cache, teardown := rediscache.TestCache(t)
	defer teardown()

	err := cache.Set(1, 1, time.Minute)
	assert.NoError(t, err)
	data, err := cache.Get(1)
	val := data
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
}
