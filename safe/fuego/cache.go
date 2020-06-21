package cache

import (
	"github.com/tomiok/fuego-cache/safe/encoding"
	"github.com/tomiok/fuego-cache/safe/hash"
	"sync"
)

//Cache is the base structure for Fuego cache.
type Cache struct {
	Cache *Fuego
	Lock  sync.RWMutex
}

//Fuego
type Fuego struct {
	Entries map[int][]byte
}

//FuegoEntry
type Entry struct {
	Key   int
	Value []byte
}

//AddOne will add an entry into the key-value store
func (c *Cache) AddOne(e Entry) bool {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Cache.Entries[e.Key] = e.Value
	return true
}


func ToEntry(key, value interface{}) Entry {
	encode := encoding.Encode(value)
	return Entry{
		Key:   hash.Hash(key),
		Value: encode.Bytes(),
	}
}