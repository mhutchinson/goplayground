package fib

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
