package embedded

import (
	cache "github.com/tomiok/fuego-cache/fuego"
	"github.com/tomiok/fuego-cache/logs"
)

type FuegoEmbedded interface {
	Insert(key, value string) error
	Delete(key string) error
	Get(key string) (string, error)
	List() []string
}

// FuegoInMemory is a mode for embedded database
type FuegoInMemory struct {
	DB *cache.InMemoryDB
}

func (f *FuegoInMemory) Insert(key, value string) error {
	res, err := f.DB.Fuego.SetOne(key, value)

	if err != nil {
		logs.Error(err.Error())
		return err
	}

	logs.Info("result: " + res)

	return nil
}

func (f *FuegoInMemory) Get(key string) (string, error) {
	res, err := f.DB.Fuego.GetOne(key)

	if err != nil {
		logs.Error(err.Error())
		return "", err
	}

	return res, nil
}
