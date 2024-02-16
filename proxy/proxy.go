package proxy

import (
	"fmt"
	"log"
	"time"

	"github.com/sleep46/gs-Proxy/cache"
	"github.com/sleep46/gs-Proxy/redis"
)


type Proxy struct {
	cache  *cache.LRUCache
	redis  *redis.Client
}

func NewProxy(cacheCapacity int, globalExpiry int, redisAddress string) (*Proxy, error) {
	cache := cache.NewLRUCache(cacheCapacity, globalExpiry)
	redis, err := redis.NewClient(redisAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to create Redis client: %v", err)
	}

	return &Proxy{
		cache:  cache,
		redis:  redis,
	}, nil
}


func (p *Proxy) Get(key string) (string, error) {
	value, found := p.cache.Get(key)
	if found {
		return value, nil
	}

	value, err := p.redis.Get(key)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve value from Redis: %v", err)
	}

	p.cache.Set(key, value)

	return value, nil
}
