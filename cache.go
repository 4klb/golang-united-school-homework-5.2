package cache

import (
	"time"
)

type Cache struct {
	m map[string]Data
}

type Data struct {
	val     string
	expired bool
	date    time.Time
}

func NewCache() *Cache {
	return &Cache{
		make(map[string]Data),
	}
}

func (c *Cache) Put(key, value string) {
	data := Data{
		val:     value,
		expired: false,
	}
	c.m[key] = data
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	data := Data{
		val:     value,
		expired: true,
		date:    deadline,
	}
	c.m[key] = data
}

func (c *Cache) Get(key string) (string, bool) {
	startTime := time.Now()

	data, ok := c.m[key] //has that key
	if !ok {
		return "", ok
	}
	if !data.expired { // has deadline
		return data.val, ok
	}
	res := startTime.Before(data.date)
	if res {
		return data.val, true
	} else {
		delete(c.m, key)
		return "", false
	}
}

func (c *Cache) Keys() []string {
	var arr []string
	for key := range c.m {
		arr = append(arr, key)

	}
	return arr
}
