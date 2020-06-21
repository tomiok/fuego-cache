package hash

import (
	"bytes"
	"encoding/gob"
	"github.com/tomiok/fuego-cache/logs"
)

const prime = 127

func hash(i interface{}) int {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(i)

	if err != nil {
		logs.Error("cannot encode interface")
	}
	byteValues := buf.Bytes()
	var count int
	for _, v := range byteValues {
		count += prime*count + int(v)
	}

	return count
}
