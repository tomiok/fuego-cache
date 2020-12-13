package cache

import (
	"errors"
	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/persistence"
	"sync"
	"time"
)

const (
	get         = "GET"
	set         = "SET"
	del         = "DELETE"
	responseOK  = "ok"
	responseNil = "nil"
)

//cache is the base structure for Fuego cache.
type cache struct {
	//the cache instance itself.
	cache *fuego
	//read and write lock.
	lock sync.RWMutex
	//shortcut for persistence in disk boolean property
	diskPersistence bool
	//persist interface - in memory or persistent cache
	persist persistence.Persist
	//inMemory
	inMemory bool
}

func NewCache(config FuegoConfig) *cache {
	var inMemory bool

	if config.Mode == "inMemory" {
		inMemory = true
	}

	filePersist := persistence.FilePersistence{File: config.FileLocation, InMemory: inMemory}

	return &cache{
		cache: &fuego{
			entries: make(map[int]fuegoValue),
		},
		lock:            sync.RWMutex{},
		diskPersistence: config.DiskPersistence,
		persist:         &filePersist,
		inMemory:        filePersist.InMemory,
	}
}

//fuego is the actual node of the cache.
type fuego struct {
	entries map[int]fuegoValue
}

type InMemoryDB struct {
	Fuego *cache
}

//fuegoValue is the actual value to store and the ttl.
type fuegoValue struct {
	value string
	ttl   int64 // time to live in seconds
}

//FuegoEntry is the cache entry, composed by the key and the fuego value, which contains the value to store and the ttl.
type entry struct {
	key    int
	object fuegoValue
}

func (c *cache) Clear() {
	c.cache.entries = make(map[int]fuegoValue)
}

//SetOne will add an entry into the key-value store.
func (c *cache) SetOne(k string, v string, ttl ...int) (string, error) {
	expiration := -1
	if len(ttl) > 0 {
		expiration = ttl[0]
	}
	e := toEntry(k, v, expiration)

	c.lock.Lock()
	c.cache.entries[e.key] = fuegoValue{value: e.object.value, ttl: e.object.ttl}
	if c.diskPersistence {
		c.persist.Save(get, e.key, e.object.value)
	}

	c.lock.Unlock()
	return responseOK, nil
}

//GetOne will return a value in the cache if the key lookup is OK and the
//value is not expired. Otherwise, the an error will be returned.
func (c *cache) GetOne(key string) (string, error) {
	c.lock.RLock()
	hashedKey := internal.ApplyHash(key)
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
func (c *cache) DeleteOne(key string) string {
	c.lock.RLock()
	hashKey := internal.ApplyHash(key)
	_, ok := c.cache.entries[hashKey]

	if ok {
		delete(c.cache.entries, hashKey)
		if c.diskPersistence {
			c.persist.Save(del, hashKey, "") //value does not matter in delete operation
		}
		c.lock.RUnlock()
		return responseOK
	}
	c.lock.RUnlock()
	return responseNil
}

func (c *cache) List() []string {
	entries := c.cache.entries
	var values []string
	for _, v := range entries {
		values = append(values, v.value)
	}

	return values
}

//Count will show how many elements are in the cache (all the nodes)
func (c *cache) Count() int {
	return len(c.cache.entries)
}

//toEntry convert key value interfaces into a system Entry.
func toEntry(key string, value string, ttl int) entry {
	// client add a TTL into the entry
	hashcode := internal.ApplyHash(key)
	if ttl > 0 {
		return entry{
			key: hashcode,
			object: fuegoValue{
				value: value,
				ttl:   int64(ttl) + time.Now().Unix(),
			},
		}
	}

	return entry{
		key: hashcode,
		object: fuegoValue{
			value: value,
			ttl:   -1,
		},
	}
}
