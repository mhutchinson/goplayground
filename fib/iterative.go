package fib

// IterativeFib returns the nth value from the fibonnaci sequence
func IterativeFib(n int) int {
	if n < 0 {
		return 0
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		t := a + b
		a = b
		b = t
	}
	return b
}
