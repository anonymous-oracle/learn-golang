package main

import "testing"

func TestMySum(t *testing.T) {
	if x := mySum(2, 3); x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}
