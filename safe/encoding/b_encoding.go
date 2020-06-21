package encoding

import (
	"bytes"
	"encoding/gob"
	"github.com/tomiok/fuego-cache/logs"
)

func Encode(i interface{}) bytes.Buffer {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(i)

	if err != nil {
		logs.Error("cannot encode the interface: " + err.Error())
		return bytes.Buffer{}
	}

	return b
}
