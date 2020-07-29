package cache

import (
	"errors"
	"github.com/tomiok/fuego-cache/persistence"
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
	//the cache instance itself.
	cache *fuego
	//read and write lock.
	lock sync.RWMutex
	//persist interface
	persist persistence.Persist
	//shortcut for persistence in disk boolean property
	diskPersistence bool
}

func NewCache(config FuegoConfig) *cache {
	filePersist := persistence.FilePersistence{File: config.FileLocation}

	return &cache{
		cache: &fuego{
			entries: make(map[int]fuegoValue),
		},
		lock:            sync.RWMutex{},
		diskPersistence: config.DiskPersistence,
		persist:         &filePersist,
	}
}

//fuego is the actual node of the cache.
type fuego struct {
	entries map[int]fuegoValue
}

//FuegoEntry is the cache entry, composed by the key and the fuego value, which contains the value to store and the ttl.
type entry struct {
	key    int
	object fuegoValue
}

//fuegoValue is the actual value to store and the ttl.
type fuegoValue struct {
	value string
	ttl   int64
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
	if c.diskPersistence {
		c.persist.Save("set", e.key, e.object.value)
	}

	c.lock.Unlock()
	return responseOK, nil
}

//GetOne will return a value in the cache if the key lookup is OK and the
//value is not expired. Otherwise, the an error will be returned.
func (c *cache) GetOne(key interface{}) (string, error) {
	c.lock.RLock()
	hashedKey := Apply(key)
	val, ok := c.cache.entries[hashedKey]

	if ok {
		if ttl := val.ttl; ttl > 0 { // when TTL is negative, the entry will not expire
			if time.Now().Unix() < ttl { // TTL bigger means that is not expired
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
	c.lock.RUnlock()
	return responseNil, errors.New("key not found")
}

//DeleteOne will delete the entry given a key, returns "ok" if is deleted, otherwise "nil".
func (c *cache) DeleteOne(key interface{}) string {
	c.lock.RLock()
	hashKey := Apply(key)
	_, ok := c.cache.entries[hashKey]

	if ok {
		delete(c.cache.entries, hashKey)
		if c.diskPersistence {
			c.persist.Save("del", hashKey, "") //value does not matter in delete operation
		}
		c.lock.RUnlock()
		return responseOK
	}
	c.lock.RUnlock()
	return responseNil
}

//Count will show how many elements are in the cache (all the nodes)
func (c *cache) Count() int {
	return len(c.cache.entries)
}

//toEntry convert key value interfaces into a system Entry.
func toEntry(key interface{}, value string, ttl int) entry {
	// client add a TTL into the entry
	hashedKey := Apply(key)
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
