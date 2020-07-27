package persistence

import (
	"fmt"
	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/logs"
	"os"
	"path/filepath"
	"time"
)

type Persist interface {
	Save(operation string, k int, value string)
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
	//read a file if already exists, or create a new one
	file, err := os.OpenFile(filepath.Join(f.File), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		logs.LogError(err)
		// no error returned, just a shame
		return
	}

	defer internal.OnCloseError(file.Close)

	record := buildRecord(operation, k, value)

	_, err = file.WriteString(record)

	if err != nil {
		logs.LogError(err)
		// no error returned, just a shame
	}
}

func buildRecord(operation string, k int, value string) string {
	return fmt.Sprintf("%s,%d,%s,%d\n", operation, k, value, time.Now().Unix())
}
