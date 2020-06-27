package cache

import (
	"github.com/tomiok/fuego-cache/logs"
	"github.com/tomiok/fuego-cache/safe/hash"
	"sync"
)

//FuegoOps
type FuegoOps interface {
	Apply()
}

type g func(key interface{}) string
type a func(e Entry) bool

type WriteOperation struct {
	Operation string
	Key       string
	Value     string
	DoAdd     a
}

func (f *WriteOperation) Apply() {
	e, err := ToEntry(f.Key, f.Value)

	if err != nil {
		logs.LogError(err)
		return
	}

	f.DoAdd(e)
}

type ReadOperation struct {
	Operation string
	Key       string
	DoGet     g
}

func (r *ReadOperation) Apply() {
	r.DoGet(r.Key)
}

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

//AddOne will add an entry into the key-value store.
func (c *Cache) AddOne(e Entry) bool {
	c.Lock.Lock()
	c.Cache.Entries[e.Key] = e.Value
	c.Lock.Unlock()
	return true
}

func (c *Cache) GetOne(key interface{}) string {
	c.Lock.RLock()
	val := c.Cache.Entries[hash.Hash(key)]
	c.Lock.RUnlock()
	return val
}

//ToEntry convert key value interfaces into a system Entry.
func ToEntry(key interface{}, value string) (Entry, error) {
	return Entry{
		Key:   hash.Hash(key),
		Value: value,
	}, nil
}
