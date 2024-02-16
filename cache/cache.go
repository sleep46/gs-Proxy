package cache

import (
	"container/list"
	"time"
	"sync"
)

type LRUCache struct {
	capacity     int
	globalExpiry int
	cache        map[string]*list.Element
	evictList    *list.List
	mutex        sync.Mutex
}

type entry struct {
	key       string
	value     string
	timestamp time.Time
}

func NewLRUCache(capacity, globalExpiry int) *LRUCache {
	return &LRUCache{
		capacity:     capacity,
		globalExpiry: globalExpiry,
		cache:        make(map[string]*list.Element),
		evictList:    list.New(),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		cacheEntry := elem.Value.(*entry)
		if time.Since(cacheEntry.timestamp) <= time.Duration(c.globalExpiry)*time.Millisecond {
			c.evictList.MoveToFront(elem)
			return cacheEntry.value, true
		}
		c.evictList.Remove(elem)
		delete(c.cache, key)
	}

	return "", false
}

func (c *LRUCache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		cacheEntry := elem.Value.(*entry)
		cacheEntry.value = value
		cacheEntry.timestamp = time.Now()
		c.evictList.MoveToFront(elem)
		return
	}

	if len(c.cache) >= c.capacity {
		lastElem := c.evictList.Back()
		if lastElem != nil {
			delete(c.cache, lastElem.Value.(*entry).key)
			c.evictList.Remove(lastElem)
		}
	}

	newEntry := &entry{
		key:       key,
		value:     value,
		timestamp: time.Now(),
	}
	elem := c.evictList.PushFront(newEntry)
	c.cache[key] = elem
}
