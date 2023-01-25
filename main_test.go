package main

import (

  "testing"
)

func TestAdd(t *testing.T) {
	if Add1(2, 2) != 4 {
		t.Fail()
	}
}

