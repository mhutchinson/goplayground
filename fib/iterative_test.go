package fib

import (
	"testing"
)

func TestIterative(t *testing.T) {
	RunTests(t, IterativeFib)
}
