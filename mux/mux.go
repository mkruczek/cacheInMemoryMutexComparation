package mux

import "sync"

type Cache struct {
	data map[string]int
	cu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]int)}
}

func (c *Cache) Set(key string, val int) {
	c.cu.Lock()
	defer c.cu.Unlock()
	c.data[key] = val
}

func (c *Cache) Get(key string) (int, bool) {
	c.cu.RLock()
	defer c.cu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Delete(key string) {
	c.cu.Lock()
	defer c.cu.Unlock()
	delete(c.data, key)
}
