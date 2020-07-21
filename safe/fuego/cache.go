package cache

import (
	"errors"
	"github.com/tomiok/fuego-cache/safe/hash"
	"sync"
	"time"
)

const (
	get         = "GET"
	set         = "SET"
	responseOK  = "ok"
	responseNil = "nil"
)

//cache is the base structure for Fuego cache.
type cache struct {
	cache *fuego
	lock  sync.RWMutex // read and write lock
}

func NewCache() *cache {
	return &cache{
		cache: &fuego{
			entries: make(map[int]fuegoValue),
		},
		lock: sync.RWMutex{},
	}
}

//fuego
type fuego struct {
	entries map[int]fuegoValue
}

type fuegoValue struct {
	value string
	ttl   int64
}

//FuegoEntry
type entry struct {
	key    int
	object fuegoValue
}

//SetOne will add an entry into the key-value store.
func (c *cache) SetOne(k interface{}, v string, ttl ...int) (string, error) {
	expiration := -1
	if len(ttl) > 0 {
		expiration = ttl[0]
	}
	e := toEntry(k, v, expiration)

	c.lock.Lock()
	c.cache.entries[e.key] = fuegoValue{value: e.object.value, ttl: e.object.ttl}
	c.lock.Unlock()
	return responseOK, nil
}

func (c *cache) GetOne(key interface{}) (string, error) {
	c.lock.RLock()
	hashedKey := hash.Apply(key)
	val, ok := c.cache.entries[hashedKey]

	if ok {
		if ttl := val.ttl; ttl > 0 { // when TTL is negative, the entry will not expire
			if time.Now().Unix() > ttl {
				c.lock.RUnlock()
				return val.value, nil
			}
			delete(c.cache.entries, hashedKey)
			c.lock.RUnlock()
			return responseNil, errors.New("key expired")
		}

		c.lock.RUnlock()
		return val.value, nil
	}

	return responseNil, errors.New("key not found")
}

func (c *cache) DeleteOne(key interface{}) string {
	c.lock.RLock()
	hashKey := hash.Apply(key)
	_, ok := c.cache.entries[hashKey]

	if ok {
		delete(c.cache.entries, hashKey)
		c.lock.RUnlock()
		return responseOK
	}
	c.lock.RUnlock()
	return responseNil
}

func (c *cache) Count() int {
	return len(c.cache.entries)
}

//toEntry convert key value interfaces into a system Entry.
func toEntry(key interface{}, value string, ttl int) entry {
	// client add a TTL into the entry
	hashedKey := hash.Apply(key)
	if ttl > 0 {
		return entry{
			key: hashedKey,
			object: fuegoValue{
				value: value,
				ttl:   int64(ttl) + time.Now().Unix(),
			},
		}
	}

	return entry{
		key: hashedKey,
		object: fuegoValue{
			value: value,
			ttl:   -1,
		},
	}
}
