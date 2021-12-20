package main

import "testing"

func TestSum(t *testing.T) {
	var v = Sum(1, 2)
	if v != 3 {
		t.Error("Expected", 3, "Got", v)
	}
}
