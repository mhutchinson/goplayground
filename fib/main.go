package fib

import "fmt"

// RunMain prints out some fib numbers for a quick demo
func RunMain() {
	n := 30
	fmt.Println("Recursive", RecursiveFib(n))
	fmt.Println("Iterative", IterativeFib(n))
}
