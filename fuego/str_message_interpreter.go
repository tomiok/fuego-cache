package cache

import (
	"errors"
	"strings"
)

const (
	space     = " "
	backSlash = "\""
	intro     = "\n"
)

type Message struct {
	InputMessage string
	ErrResponse  string
}

func NewFuegoMessage(msg string) *Message {
	return &Message{
		InputMessage: strings.TrimSuffix(msg, intro),
		ErrResponse:  responseNil,
	}
}

func (m *Message) Compute(cache *cache) (FuegoOps, error) {
	operation := strings.Split(m.InputMessage, space)
	l := len(operation)
	if l == 0 {
		return nil, errors.New("")
	}

	//read
	if l == 2 {
		action := operation[0]
		if strings.ToUpper(action) != get {
			return nil, errors.New("")
		}
		return &ReadOperation{
			Operation: operation[0],
			Key:       operation[1],
			DoGet:     cache.GetOne,
		}, nil
	}

	//write
	if l == 3 {
		action := operation[0]
		if strings.ToUpper(action) != set {
			return nil, errors.New("")
		}
		return &WriteOperation{
			Operation: operation[0],
			Key:       operation[1],
			Value:     operation[2],
			DoAdd:     cache.SetOne,
		}, nil
	}

	return nil, errors.New("operation not supported")
}

// fetchMessage is the function that takes the input from the CLI and separate the message in
// 3 strings. OPERATION, KEY, VALUE, in that order. The inpunt message should take restrictively some rules. Those
// rules are: {operation} {key} {double quoted value}, only one space is the separation and the value is always quoted.
func fetchMessage(msg string) (string, string, string) {
	ss := strings.SplitAfter(msg, backSlash)
	l := len(ss)
	if l == 0 {
		return "", "", ""
	}

	value := ss[1]
	value = strings.TrimSuffix(value, backSlash)
	operation, key := getOperationAndKey(ss[0])
	return operation, key, value
}

func getOperationAndKey(s string) (string, string) {
	ss := strings.Split(s, space)
	return ss[0], ss[1]
}
