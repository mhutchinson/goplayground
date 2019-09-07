package fib

import (
	"sync"
)

// RecursiveFib returns the nth value from the fibonnaci sequence
func RecursiveFib(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return 1
	}
	return RecursiveFib(n-1) + RecursiveFib(n-2)
}

// ParallelRecursiveFib returns the nth value from the fibonnaci sequence
func ParallelRecursiveFib(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return 1
	}
	var wg sync.WaitGroup
	wg.Add(2)
	var a, b int
	go func() {
		defer wg.Done()
		a = ParallelRecursiveFib(n - 1)
	}()
	go func() {
		defer wg.Done()
		b = ParallelRecursiveFib(n - 2)
	}()
	wg.Wait()
	return a + b
}
