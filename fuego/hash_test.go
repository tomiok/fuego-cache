package cache

import (
	"github.com/tomiok/fuego-cache/internal"
	"testing"
)

func Test_hashing(t *testing.T) {
	val1 := "1"
	val2 := "2"
	val3 := "3"

	h1 := internal.ApplyHash(val1)
	h2 := internal.ApplyHash(val2)
	h3 := internal.ApplyHash(val3)

	if h1 == h2 {
		t.Fail()
	}

	if h1 == h3 {
		t.Fail()
	}

	if h2 == h3 {
		t.Fail()
	}
}
