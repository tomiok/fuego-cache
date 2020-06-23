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

func Test_AddGetOne(t *testing.T) {
	fuegoCache := NewCache()
	e, err := ToEntry(1, "1")

	if err != nil {
		t.Fail()
	}
	res := fuegoCache.AddOne(e)

	if !res {
		t.Log("cannot add")
		t.Fail()
	}

	value := fuegoCache.GetOne(1)

	if value != "1" {
		t.Log("cannot read " + value)
		t.Fail()
	}
}
