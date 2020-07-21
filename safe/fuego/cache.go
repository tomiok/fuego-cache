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

//Cache is the base structure for Fuego cache.
type Cache struct {
	Cache *Fuego
	Lock  sync.RWMutex // read and write lock
	TTL   int64        // TTL in seconds
}

func NewCache() *Cache {
	return &Cache{
		Cache: &Fuego{
			Entries: make(map[int]FuegoValue),
		},
		Lock: sync.RWMutex{},
	}
}

//Fuego
type Fuego struct {
	Entries map[int]FuegoValue
}

type FuegoValue struct {
	Value string
	TTL   int64
}

//FuegoEntry
type Entry struct {
	Key    int
	Object FuegoValue
}

//SetOne will add an entry into the key-value store.
func (c *Cache) SetOne(k interface{}, v string, ttl ...int) (string, error) {
	expiration := -1
	if len(ttl) > 0 {
		expiration = ttl[0]
	}
	e := ToEntry(k, v, expiration)

	c.Lock.Lock()
	c.Cache.Entries[e.Key] = FuegoValue{Value: e.Object.Value, TTL: e.Object.TTL}
	c.Lock.Unlock()
	return responseOK, nil
}

func (c *Cache) GetOne(key interface{}) (string, error) {
	c.Lock.RLock()
	hashedKey := hash.Apply(key)
	val, ok := c.Cache.Entries[hashedKey]

	if ok {
		if ttl := val.TTL; ttl > 0 { // when TTL is negative, the entry will not expire
			if time.Now().Unix() > ttl {
				c.Lock.RUnlock()
				return val.Value, nil
			}
			delete(c.Cache.Entries, hashedKey)
			c.Lock.RUnlock()
			return responseNil, errors.New("key expired")
		}

		c.Lock.RUnlock()
		return val.Value, nil
	}

	return responseNil, errors.New("key not found")
}

func (c *Cache) DeleteOne(key interface{}) string {
	c.Lock.RLock()
	hashKey := hash.Apply(key)
	_, ok := c.Cache.Entries[hashKey]

	if ok {
		delete(c.Cache.Entries, hashKey)
		c.Lock.RUnlock()
		return responseOK
	}
	c.Lock.RUnlock()
	return responseNil
}

func (c *Cache) Count() int {
	return len(c.Cache.Entries)
}

//ToEntry convert key value interfaces into a system Entry.
func ToEntry(key interface{}, value string, ttl int) Entry {
	// client add a TTL into the entry
	hashedKey := hash.Apply(key)
	if ttl > 0 {
		return Entry{
			Key: hashedKey,
			Object: FuegoValue{
				Value: value,
				TTL:   int64(ttl) + time.Now().Unix(),
			},
		}
	}

	return Entry{
		Key: hashedKey,
		Object: FuegoValue{
			Value: value,
			TTL:   -1,
		},
	}
}
