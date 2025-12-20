package main

import "fmt"

func fact(n int) int {
	if n == 0 {
	  return 1
	}

	return n * fact(n - 1)
}

func main() {
	fmt.Println("fact 7:", fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 { return n }
		return fib(n - 1) + fib(n - 2)
	}

	fmt.Println(fib(7))
}

// maffs
// n = 7
// fib(5) + fib(6)
// fib(3) + fib(4) + fib(4) + fib(5)
// 1 + fib(2) + fib(3) + fib(2) + fib(3) + fib(3) + fib(4)
// 1 + 0 + 1 + 1 + 0 + 1 + 0 + 1 + 1 + 1 + fib(2) + 1 + fib(2) + fib(2) + fib(3)
// 7 + 1 + 1 + 1 + 1 + 1 + fib(2)
// 12 + 1
// 13