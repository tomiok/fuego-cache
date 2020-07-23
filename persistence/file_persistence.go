package persistence

import (
	"fmt"
	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/logs"
	"os"
	"time"
)

type Persist interface {
	Save(operation string, k int, value string)
}

func Apply(p Persist, operation string, key int, value string) {
	data := Data{
		operation: operation,
		key:       key,
		value:     value,
	}
	p.Save(data.operation, data.key, data.value)
}

type Data struct {
	operation string
	key       int
	value     string
}

type FilePersistence struct {
	File string
}

func (f *FilePersistence) Save(operation string, k int, value string) {
	file, err := os.Create(f.File)

	if err != nil {
		logs.LogError(err)
		// no error returned, just a shame
	}
	defer internal.OnCloseError(file.Close)

	record := buildRecord(operation, k, value)

	_, err = file.Write(record)

	if err != nil {
		logs.LogError(err)
		// no error returned, just a shame
	}
}

func buildRecord(operation string, k int, value string) []byte {
	s := fmt.Sprintf("%s %d %s %s", operation, k, value, time.Now().String())
	return []byte(s)
}
