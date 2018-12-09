package main

import "testing"

func TestMain(t *testing.T) {
	if actual, expected := message, "Hello, 世界"; actual != expected {
		t.Errorf("Actual message (%s) does not equal expected message (%s)", actual, expected)
	}
}
