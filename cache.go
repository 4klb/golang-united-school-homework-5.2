package cache

import "time"

type Cache struct {
	m map[string]string
}

func NewCache() Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	k, ok := c.m[key]
	if !ok {
		return "", ok
	}
	return k, ok
}

func (c *Cache) Put(key, value string) {
	c.m[key] = value
}

func (c *Cache) Keys() []string {
	var keys []string

	for key, _ := range c.m {
		keys = append(keys, key)
	}

	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	startTime := time.Now()
	res := startTime.Before(deadline)
	if res {
		time.Sleep(2 * time.Second)
		c.m[key] = value
		time.Sleep(2 * time.Second)
	}
}
