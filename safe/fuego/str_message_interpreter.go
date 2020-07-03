package cache

import "strings"

const space = " "

type Message struct {
	InputMessage string
}

func NewFuegoMessage(msg string) *Message {
	return &Message{
		InputMessage: strings.TrimSuffix(msg, "\n"),
	}
}

func (m *Message) Compute(cache *Cache) FuegoOps {
	operation := strings.Split(m.InputMessage, space)
	l := len(operation)
	if l == 0 {
		return nil
	}

	//read
	if l == 2 {
		action := operation[0]
		if strings.ToUpper(action) != get {
			return nil
		}
		return &ReadOperation{
			Operation: operation[0],
			Key:       operation[1],
			DoGet:     cache.GetOne,
		}
	}

	//write
	if l == 3 {
		action := operation[0]
		if strings.ToUpper(action) != set {
			return nil
		}
		return &WriteOperation{
			Operation: operation[0],
			Key:       operation[1],
			Value:     operation[2],
			DoAdd:     cache.AddOne,
		}
	}

	return nil
}
