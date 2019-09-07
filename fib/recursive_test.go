package fib

import (
	"testing"
)

func TestRecursive(t *testing.T) {
	RunTests(t, RecursiveFib)
}
