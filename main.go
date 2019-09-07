package main

import (
	"fmt"

	"github.com/mhutchinson/goplayground/fib"
)

func main() {
	Fib()
}

// Fib does fib
func Fib() {
	n := 30
	fmt.Println("Recursive", fib.RecursiveFib(n))
	fmt.Println("ParallelRecursive", fib.ParallelRecursiveFib(n))
	fmt.Println("Iterative", fib.IterativeFib(n))
}
