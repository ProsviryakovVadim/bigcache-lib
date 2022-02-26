package cache

import (
	"encoding/json"

	"github.com/allegro/bigcache"
)

// Interface - interface
type Interface interface {
	Get(key string, object interface{}) interface{}
	Set(key string, object interface{}) error
	Delete(key ...string) error
	Clear() error
}

type BigCache struct {
	cache *bigcache.BigCache
}

//NewBigCache - new instance
func NewBigCache(c *bigcache.BigCache) *BigCache {
	return &BigCache{cache: c}
}

func (c *BigCache) Get(key string, object interface{}) interface{} {
	result, err := c.cache.Get(key)
	if err != nil {
		return nil
	}
	_ = json.Unmarshal([]byte(result), object)

	return object
}

func (c *BigCache) Set(key string, object interface{}) error {
	value, err := json.Marshal(object)
	if err != nil {
		return err
	}
	err = c.cache.Set(key, []byte(value))
	return err
}

func (c *BigCache) Delete(key ...string) error {
	for _, v := range key {
		err := c.cache.Delete(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *BigCache) Clear() error {
	err := c.cache.Reset()
	return err
}
