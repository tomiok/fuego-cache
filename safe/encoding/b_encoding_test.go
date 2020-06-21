package encoding

import (
	"testing"
)

func Test_b_encoding(t *testing.T) {
	v1 := 1
	v2 := "1"
	buff1 := Encode(v1)
	buff2 := Encode(v2)

	equals := true
	if buff1.Len() == buff2.Len() {
		b1 := buff1.Bytes()
		b2 := buff2.Bytes()
		for i := 0; i < buff1.Len(); i++ {
			if b1[i] != b2[i] {
				equals = false
			}
		}
	}else {
		equals = false
	}

	if equals {
		t.Fail()
	}

}
