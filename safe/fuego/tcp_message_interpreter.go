package cache

import "strings"

const space = " "

type Message struct {
	InputMessage string
	//	Validate(op, k, v string) error
}

func NewFuegoMessage(msg string) *Message {
	return &Message{
		InputMessage: msg,
	}
}

func (m *Message) Compute() FuegoOps {
	operation := strings.Split(m.InputMessage, space)

	if len(operation) == 0 {
		return nil
	}

	//TODO finish this implementation
	return nil
}
