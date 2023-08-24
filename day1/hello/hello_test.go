package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	v := hello.SayHi()
	if v != "Hello" {
		t.Errorf("Expect Hello but got %v", v)
	}
}
