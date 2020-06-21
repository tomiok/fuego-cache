package cache

import (
	"errors"
	"github.com/tomiok/fuego-cache/safe/encoding"
	"github.com/tomiok/fuego-cache/safe/hash"
	"sync"
)

//Cache is the base structure for Fuego cache.
type Cache struct {
	Cache *Fuego
	Lock  sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		Cache: &Fuego{},
		Lock:  sync.RWMutex{},
	}
}

//Fuego
type Fuego struct {
	Entries map[int][]byte
}

//FuegoEntry
type Entry struct {
	Key        int
	Value      []byte
	Expiration int64
}

//AddOne will add an entry into the key-value store.
func (c *Cache) AddOne(e Entry) bool {
	c.Lock.Lock()
	c.Cache.Entries[e.Key] = e.Value
	c.Lock.Unlock()
	return true
}

func (c *Cache) GetOne(key interface{}) interface{} {
	c.Lock.RLock()
	val := c.Cache.Entries[hash.Hash(key)]
	c.Lock.RUnlock()
	return val
}

//ToEntry convert key value interfaces into a system Entry.
func ToEntry(key, value interface{}) (Entry, error) {
	encode := encoding.Encode(value)
	if encode.Len() == 0 {
		return Entry{}, errors.New("cannot encode the value")
	}

	return Entry{
		Key:   hash.Hash(key),
		Value: encode.Bytes(),
	}, nil
}
