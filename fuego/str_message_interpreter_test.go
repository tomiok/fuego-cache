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

	operation, _ := msg.Compute(fuegoCache)
	res := operation.Apply()
	if res.Response != "OK" && fuegoCache.Count() != 1 {
		t.Fail()
	}

	msg = NewFuegoMessage(correctMessageGet)

	operation, _ = msg.Compute(fuegoCache)
	res = operation.Apply()

	if res.Response != "1" {
		t.Fail()
	}
}

func Test_Compute_incorrectMessage(t *testing.T) {
	fuegoCache := NewCache(defaultConfigs())
	msg := NewFuegoMessage(incorrectMessage)
	operation, err := msg.Compute(fuegoCache)

	if operation != nil && err == nil {
		t.Fail()
	}
}

func Test_getInQuotes(t *testing.T) {
	s := "hey dude \"how are you\""
	expected := "how are you"
	res := getInQuotes(s)

	if res != expected {
		t.Fail()
	}
}
