package hash

import (
	"fmt"
	"testing"
)

func Test_hashing(t *testing.T) {
	val1 := 1
	val2 := "1"
	val3 := struct {
		Val int
	}{Val : 1}

	h1 := Hash(val1)
	h2 := Hash(val2)
	h3 := Hash(val3)

	if h1 == h2 {
		t.Fail()
	}

	if h1 == h3 {
		t.Fail()
	}

	if h2 == h3 {
		t.Fail()
	}
	fmt.Println(h1,h2,h3, m)
}