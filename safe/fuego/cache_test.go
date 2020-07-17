package cache

import (
	"testing"
)

func Test_cacheConstructor(t *testing.T) {
	c := NewCache()
	if c == nil {
		t.Fail()
	}
}

func Test_SetAndGetOne(t *testing.T) {
	fuegoCache := NewCache()
	e, err := ToEntry(1, "1")

	if err != nil {
		t.Fail()
	}
	res := fuegoCache.SetOne(e)

	if res != "ok" {
		t.Log("cannot add")
		t.Fail()
	}

	value := fuegoCache.GetOne(1)

	if value != "1" {
		t.Log("cannot read " + value)
		t.Fail()
	}
}

func Test_DeleteOne(t *testing.T) {
	fuegoCache := NewCache()
	e, err := ToEntry(1, "hello there")

	if err != nil {
		t.Fail()
	}

	fuegoCache.SetOne(e)

	count := fuegoCache.Count()

	if count != 1 {
		t.Fail()
	}

	res := fuegoCache.DeleteOne(2) //should be nil the response since 2 is not a key inserted

	if res != "nil" {
		t.Fail()
	}

	res = fuegoCache.DeleteOne(1)

	if res != "ok" {
		t.Fail()
	}

	count = fuegoCache.Count()

	if count != 0 {
		t.Fail()
	}
}
