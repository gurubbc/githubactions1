package main

import (
	"fmt"
  "testing"
)

func TestAdd(t *testing.T) {
	if mycalculator.Add1(2, 2) != 4 {
		t.Fail()
	}
}

