package embedded

import (
	cache "github.com/tomiok/fuego-cache/fuego"
	"github.com/tomiok/fuego-cache/logs"
	"testing"
)

func Test_Insert_and_get_InMemory(t *testing.T) {
	c := cache.NewCache(cache.FuegoConfig{
		DiskPersistence: true,
		FileLocation:    "C:\\Users\\Tomás\\Downloads\\fuego.csv",
		Mode:            "",
	})
	fuego := FuegoInMemory{
		DB: &cache.InMemoryDB{Fuego: c},
	}

	err := fuego.Insert("1", "hola amigos")

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	res, err := fuego.Get("1")

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	logs.Info(res)
}

