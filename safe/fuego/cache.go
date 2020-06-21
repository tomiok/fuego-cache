package cache

import "sync"

//Cache is the base structure for Fuego cache.
type Cache struct {
	Cache *Fuego
	Lock  sync.RWMutex
}

//Fuego
type Fuego struct {
	Entries map[int64][]byte
}

//FuegoEntry
type Entry struct {
	Key   int64
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

	return Entry{
		Key:   0,
		Value: nil,
	}
}