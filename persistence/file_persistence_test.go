package persistence

import (
	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/logs"
	"testing"
)

var filePersistence = FilePersistence{
	File:     "./test.csv",
	InMemory: true,
}

func TestFilePersistence_Get(t *testing.T) {
	key := internal.ApplyHash("123")
	filePersistence.Save(key, "test")

	res, err := filePersistence.Get("123")

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	logs.Info(res)
}
