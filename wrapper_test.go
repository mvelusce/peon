package main

import (
	"testing"
)

func TestWrapper(t *testing.T) {

	if 1 != 2 {
		t.Error("asd")
	}
}
