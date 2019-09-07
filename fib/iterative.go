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

// IterativeGenFib returns the nth value using GeneratorFib
func IterativeGenFib(n int) int {
	fib := GeneratorFib()
	r := fib()
	for i := 0; i < n; i++ {
		r = fib()
	}
	return r
}

// GeneratorFib returns a function that returns incremental values from fib
func GeneratorFib() func() int {
	a, b := 0, 1
	return func() int {
		r := b
		b = a + b
		a = r
		return r
	}
}
