package main

import "fmt"

func fact(x int) int {
	if x == 0 {
		return 1
	}

	return x * fact(x-1)
}
func main() {
	fmt.Println(fact(9))

	//Fibonacci number
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}
    fmt.Println(fib(7))
}
