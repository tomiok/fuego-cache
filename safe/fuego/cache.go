package cache

import (
	"errors"
	"github.com/tomiok/fuego-cache/logs"
	"github.com/tomiok/fuego-cache/safe/hash"
	"sync"
)

//FuegoOps
type FuegoOps interface {
	AddOne(func(entry Entry) bool)
	GetOne(func(key interface{}) string) string
}

type FuegoOperation struct {
	Operation string
	Key       string
	Value     string
}

func (f *FuegoOperation) AddOne(addFn func(e Entry) bool) error {
	e, err := ToEntry(f.Key, f.Value)

	if err != nil {
		logs.LogError(err)
		return err
	}

	res := addFn(e)

	if !res {
		return errors.New("cannot add into cache")
	}
	return nil
}

func (f *FuegoOperation) GetOne(getFn func(key interface{}) string) string {
	return getFn(f.Key)
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
