package hello_test

import (
	"hello"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	v := hello.SayHi()
	if v != "Hello" {
		t.Errorf("Expect Hello but got %v", v)
	}
	assert.Equal(t, "Hello", v)
}
