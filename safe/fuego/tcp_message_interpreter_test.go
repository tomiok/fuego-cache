package cache

import (
	"testing"
)

const (
	correctMessageAdd = "add 1 1"
	correctMessageGet = "get 1"
	incorrectMessage  = "incorrect"
)

func Test_NewTCPMessageCorrectMessage(t *testing.T) {
	m := NewFuegoMessage(correctMessageAdd)
	if m == nil {
		t.Fail()
	}
}

func Test_NewTCPMessageIncorrectMessage(t *testing.T) {
	m := NewFuegoMessage(incorrectMessage)
	if m == nil {
		t.Fail()
	}
}

func TestMessage_Compute(t *testing.T) {
	fuegoCache := NewCache()
	msg := NewFuegoMessage(correctMessageAdd)

	res := msg.Compute(fuegoCache).Apply()

	if res.Response != "OK" && fuegoCache.Count() != 1 {
		t.Fail()
	}

	msg = NewFuegoMessage(correctMessageGet)

	res = msg.Compute(fuegoCache).Apply()

	if res.Response != "1" {
		t.Fail()
	}
}
