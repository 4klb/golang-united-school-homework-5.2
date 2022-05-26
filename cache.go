package cache

import (
	"time"
)

type Cache struct {
	m map[string]Data
}

type Data struct {
	val         string
	hasDeadline bool
	deadline    time.Time
}

func NewCache() *Cache {
	return &Cache{
		make(map[string]Data),
	}
}

func (c *Cache) Put(key, value string) {
	data := Data{
		val:         value,
		hasDeadline: false,
	}
	c.m[key] = data
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	data := Data{
		val:         value,
		hasDeadline: true,
		deadline:    deadline,
	}
	c.m[key] = data
}

func (c *Cache) Get(key string) (string, bool) {
	startTime := time.Now()

	data, ok := c.m[key]
	if !ok {
		return "", ok
	}

	if data.hasDeadline {
		notExp := startTime.Before(data.deadline)
		if notExp {
			return data.val, true
		} else {
			delete(c.m, key)
			return "", false
		}
	}
	return data.val, ok
}

func (c *Cache) Keys() []string {
	var arr []string
	for key := range c.m {
		arr = append(arr, key)

	}
	return arr
}
