package cache

import (
	"sync"

	"github.com/Adit0507/autocomplete-search/internal/models"
)

type Cache struct {
	store map[string][]models.Suggestion
	mu    sync.RWMutex
	size  int
}

func NewCache(size int) *Cache {
	return &Cache{
		store: make(map[string][]models.Suggestion),
		size:  size,
	}
}

func (c *Cache) Get(key string) ([]models.Suggestion, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	suggestions, exists := c.store[key]

	return suggestions, exists
}

func (c *Cache) Set(key string, suggestions []models.Suggestion) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.store) >= c.size {
		for k := range c.store {
			// clear half the cache
			delete(c.store, k)
			if len(c.store) <= c.size/2 {
				break
			}
		}
	}

	c.store[key] = suggestions
}
