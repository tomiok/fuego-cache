package cache

import (
	"errors"
	"strings"
)

const space = " "

type Message struct {
	InputMessage string
	ErrResponse  string
}

func NewFuegoMessage(msg string) *Message {
	return &Message{
		InputMessage: strings.TrimSuffix(msg, "\n"),
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
