package main

import (
	"fmt"
)

func main() {
	fmt.Println(quickMaths(some(4)))
	fmt.Println(quickMaths(none[int]()))

	fmt.Println(quickMathsStatic(some(4)))
	fmt.Println(quickMathsStatic(none[int]()))

	fmt.Println("Done")
}

func quickMaths(v option[int]) int {
	square := func(i int) int { return i * i }
	add1 := func(i int) int { return i + 1 }
	return v.bind(square).bind(add1).bind(square).get()
}

func quickMathsStatic(v option[int]) string {
	square := func(i int) int { return i * i }
	add1 := func(i int) int { return i + 1 }
	toStr := func(i int) string { return fmt.Sprintf("%d", i) }
	r1 := bind(v, square)
	r2 := bind(r1, add1)
	r3 := bind(r2, square)
	return bind(r3, toStr).get()
}

func some[T any](v T) option[T] {
	return option[T]{
		result: &v,
	}
}

func none[T any]() option[T] {
	return option[T]{
		result: nil,
	}
}

type option[T any] struct {
	result *T
}

// bind applies the given function and returns another option,
// taking care not to blow up if this option is none.
// My original hope was to make this support taking a function
// that took T and returned a new type, but go _methods_ do not
// support generics!
func (o option[T]) bind(f func(T) T) option[T] {
	if o.result == nil {
		return none[T]()
	}
	return some(f(o.get()))
}

func (o option[T]) get() T {
	if o.result == nil {
		var r T
		return r
	}
	return *o.result
}

// bind applies the given function and returns another option,
// taking care not to blow up if this option is none.
func bind[T, R any](o option[T], f func(T) R) option[R] {
	if o.result == nil {
		return none[R]()
	}
	return some(f(o.get()))
}
