package cache

import (
	"testing"
)

const (
	correctMessageSet = "set 1 1"
	correctMessageGet = "get 1"
	incorrectMessage  = "incorrect"
)

func Test_NewTCPMessageCorrectMessage(t *testing.T) {
	m := NewFuegoMessage(correctMessageSet)
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
	fuegoCache := NewCache(defaultConfigs())
	msg := NewFuegoMessage(correctMessageSet)

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

func Test_Compute_incorrectMessage(t *testing.T) {
	fuegoCache := NewCache(defaultConfigs())
	msg := NewFuegoMessage(incorrectMessage)
	res := msg.Compute(fuegoCache)

	if res != nil {
		t.Fail()
	}

}
