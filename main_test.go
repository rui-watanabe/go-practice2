package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("skip")
	}
	v := IsOne(i)

	if !v {
		t.Error("%v != %v", i, 1)
	}

}
