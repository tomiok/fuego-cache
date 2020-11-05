package cache

import (
	"testing"
)

const (
	correctMessageSet = "set 1 \"1\""
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

	operation, err := msg.Compute(fuegoCache)

	if err != nil {
		t.Fatal("err " + err.Error())
	}

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

func Test_fetchMessage_Write(t *testing.T) {
	s := "set hello  \"how are you\""
	expectedOperation := "set"
	expectedValue := "how are you"
	expectedKey := "hello"
	op, err := fetchMessage(s)

	if err != nil {
		t.Fatal()
	}

	if op == nil {
		t.Fail()
	}

	if op.command != expectedOperation {
		t.Fail()
	}

	if op.key != expectedKey {
		t.Fail()
	}

	if op.value != expectedValue {
		t.Fail()
	}
}

func Test_FetchMessage_Read(t *testing.T) {
	msg := "get             hello"

	expectedOperator := "get"
	expectedKey := "hello"

	res, err := getReadMessage(msg)

	if err != nil {
		t.Fatal()
	}

	if res.command != expectedOperator {
		t.Fatal()
	}

	if res.key != expectedKey {
		t.Fatal()
	}
}
