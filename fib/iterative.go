package fib

import "iter"

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

// InfiniteFibSeq returns an iterator that yields Fibonacci numbers indefinitely.
func InfiniteFibSeq() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if !yield(b) {
				return
			}
			a, b = b, a+b
		}
	}
}

// IterativeSeqFib returns the nth value using InfiniteFibSeq
func IterativeSeqFib(n int) int {
	if n < 0 {
		return 0
	}
	i := 0
	for val := range InfiniteFibSeq() {
		if i == n {
			return val
		}
		i++
	}
	return 0
}

