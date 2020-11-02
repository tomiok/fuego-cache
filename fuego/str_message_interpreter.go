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
	readArraySize  = 2
	writeArraySize = 3
)

type messageType string

type Message struct {
	InputMessage string
	ErrResponse  string
}

type MessageOperator struct {
	cacheOperation messageType // READ, WRITE operation
	command        string      // set, get, etc...
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
	operation, err := fetchMessage(m.InputMessage)

	if err != nil {
		return nil, errors.New("") //TODO finish this error
	}

	//read - only support GET stand alone, no with multiple keys
	if operation.cacheOperation == readOperation {
		command := operation.command
		if strings.ToLower(command) != get {
			return nil, errors.New("")
		}
		return &ReadOperation{
			Key:       operation.key,
			DoGet:     cache.GetOne,
		}, nil
	}

	//write
	if operation.cacheOperation == writeOperation {
		// TODO finish this
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

	if l != writeArraySize { // should be 3 since is operation, key, value
		return nil, errors.New("the message is not formatted properly")
	}
	value := getQuotedMessage(parsed[1])

	operator, key := getOperationAndKey(parsed[0])
	return &MessageOperator{
		cacheOperation: writeOperation,
		command:        operator,
		key:            key,
		value:          value,
	}, nil
}

func getReadMessage(s string) (*MessageOperator, error) {
	msg := strings.TrimSpace(s)
	parsed := strings.SplitN(msg, space, readArraySize)

	if len(parsed) != readArraySize {
		return nil, errors.New("the message is not formatted properly")
	}

	key := strings.TrimSpace(parsed[1])

	return &MessageOperator{
		cacheOperation: readOperation,
		command:        parsed[0],
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
