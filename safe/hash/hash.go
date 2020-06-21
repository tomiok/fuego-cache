package hash

import (
	"bytes"
	"encoding/gob"
	"github.com/tomiok/fuego-cache/logs"
	"math"
)

const (
	prime = 127
	M     = math.MaxInt64
)

func hash(i interface{}) int {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(i)

	if err != nil {
		logs.Error("cannot encode interface: " + err.Error())
		return 0
	}
	byteValues := buf.Bytes()
	var index int
	for i, v := range byteValues {
		index += prime*i + int(v)%M
	}

	return index
}
