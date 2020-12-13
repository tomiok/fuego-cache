package persistence

import (
	"errors"
	"fmt"
	cache "github.com/tomiok/fuego-cache/fuego"
	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/logs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Persist interface {
	Save(operation string, k int, value string)
	Get(key string) (string, error)
}

type Data struct {
	operation string
	key       int
	value     string
}

type FilePersistence struct {
	File     string
	InMemory bool
}

func (f *FilePersistence) Save(operation string, k int, value string) {
	//read a file if already exists, or create a new one
	file, err := os.OpenFile(filepath.Join(f.File), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		logs.LogError(err)
		// no error returned, just a shame
		return
	}

	defer internal.OnCloseError(file.Close)

	record := buildRecord(operation, k, value, f.InMemory)

	_, err = file.WriteString(record)

	if err != nil {
		logs.LogError(err)
		// no error returned, just a shame
	}
}

func (f *FilePersistence) Get(key string) (string, error) {
	bytes, err := ioutil.ReadFile(f.File)

	if err != nil {
		return "", err
	}

	text := string(bytes)

	pairs := strings.Split(text, "\n")

	for _, kv := range pairs {
		values := strings.Split(kv, ",")
		hashedKey := values[0]
		i, err := strconv.Atoi(hashedKey)

		if err != nil {
			logs.Error("cannot parse key into INT type. " + err.Error())
			return "", nil
		}

		if i == cache.ApplyHash(key) {
			return values[1], nil
		}
	}

	return "", errors.New("key not found")

}

func buildRecord(operation string, k int, value string, inMemory bool) string {
	if inMemory {
		return fmt.Sprintf("%d,%s\n", k, value)
	} else {
		return fmt.Sprintf("%s,%d,%s,%d\n", operation, k, value, time.Now().Unix())
	}
}
