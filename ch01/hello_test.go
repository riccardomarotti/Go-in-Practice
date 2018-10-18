package main

import "testing"

func TestGetName(t *testing.T) {
	name := getName()

	if name != "World!" {
		t.Error("Unexpected value from getName()")
	}
}
