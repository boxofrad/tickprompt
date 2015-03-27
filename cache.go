package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

const CACHE_FILE = ".tickprompt"

type Cache struct {
	Hours    float32   `json:"hours"`
	CachedAt time.Time `json:"cached_at"`
}

func NewCache(hours float32) *Cache {
	return &Cache{
		Hours:    hours,
		CachedAt: time.Now(),
	}
}

func ReadCacheFromFile() (*Cache, error) {
	path, err := FilePath(CACHE_FILE)

	if err != nil {
		return nil, err
	}

	jsonStr, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	cache := Cache{}
	err = json.Unmarshal(jsonStr, &cache)

	return &cache, err
}

func (c *Cache) WriteToFile() error {
	path, err := FilePath(CACHE_FILE)

	if err != nil {
		return err
	}

	json, err := json.Marshal(c)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, json, 0644)
}

func (c *Cache) HasExpired(ttl time.Duration) bool {
	return c.CachedAt.Add(ttl).Before(time.Now())
}
