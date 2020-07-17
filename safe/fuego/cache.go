package cache

import (
	"github.com/tomiok/fuego-cache/safe/hash"
	"sync"
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
	Lock  sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		Cache: &Fuego{
			Entries: make(map[int]string),
		},
		Lock: sync.RWMutex{},
	}
}

//Fuego
type Fuego struct {
	Entries map[int]string
}

//FuegoEntry
type Entry struct {
	Key        int
	Value      string
	Expiration int64
}

//SetOne will add an entry into the key-value store.
func (c *Cache) SetOne(e Entry) string {
	c.Lock.Lock()
	c.Cache.Entries[e.Key] = e.Value
	c.Lock.Unlock()
	return responseOK
}

func (c *Cache) GetOne(key interface{}) string {
	c.Lock.RLock()
	val := c.Cache.Entries[hash.Hash(key)]
	c.Lock.RUnlock()
	return val
}

func (c *Cache) DeleteOne(key interface{}) string {
	c.Lock.RLock()
	hashKey := hash.Hash(key)
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
func ToEntry(key interface{}, value string) (Entry, error) {
	return Entry{
		Key:   hash.Hash(key),
		Value: value,
	}, nil
}
