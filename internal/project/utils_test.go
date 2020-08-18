package project

import "testing"

func TestTrimPrefix(t *testing.T) {

	res := TrimPrefix("./asd", "./")
	if res != "asd" {
		t.Errorf("expected asd, got %s", res)
	}
}

func TestTrimSuffix(t *testing.T) {

	res := TrimSuffix("asd/", "/")
	if res != "asd" {
		t.Errorf("expectd asd, got %s", res)
	}
}
