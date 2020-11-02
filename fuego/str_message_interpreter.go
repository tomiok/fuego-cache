package cache

import (
	"errors"
	"strings"
)

const (
	space          = " "
	doubleQuote    = "\""
	intro          = "\n"
	readOperation  = "READ"
	writeOperation = "WRITE"
)

type messageType string

type Message struct {
	InputMessage string
	ErrResponse  string
}

type MessageOperator struct {
	cacheOperation messageType // READ, WRITE operation
	operator       string      // set, get, etc...
	key            string      // the key for the cache
	value          string      // the value for the cache entry, optional for read operations
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
// 3 strings. OPERATION, KEY, VALUE, in that order when is a WRITING message, for the READ message, only operation
// and key are necessary. The input message should take restrictively some rules. Those
// rules are: {operation} {key} {double quoted value}, only one space is the separation and the value is always quoted.
func fetchMessage(msg string) (*MessageOperator, error) {
	parsed := strings.SplitAfter(msg, doubleQuote)
	l := len(parsed)
	if l == 1 {
		return getReadMessage(parsed[0])
	}

	if l != 3 { // should be 3 since is operation, key, value
		return nil, errors.New("")
	}
	value := getQuotedMessage(parsed[1])

	operator, key := getOperationAndKey(parsed[0])
	return &MessageOperator{
		cacheOperation: writeOperation,
		operator:       operator,
		key:            key,
		value:          value,
	}, nil
}

// TODO finish this
func getReadMessage(s string) (*MessageOperator, error) {
	msg := strings.TrimSpace(s)
	parsed := strings.SplitN(msg, space, 2)

	if len(parsed) != 2 {
		return nil, errors.New("")
	}

	key := strings.TrimSpace(parsed[1])
	return &MessageOperator{
		cacheOperation: readOperation,
		operator:       parsed[0],
		key:            key,
	}, nil
}

func getQuotedMessage(s string) string {
	return strings.TrimSuffix(s, doubleQuote)
}

func getOperationAndKey(s string) (string, string) {
	ss := strings.Split(s, space)
	return ss[0], ss[1]
}
