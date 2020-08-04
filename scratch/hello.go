package main

import (
	"fmt"
	"errors"
)


func fib(n int) (int, error) {
	// Returns the nth Fibonacci number, where 1st = 0 and 2nd = 1
	if n <= 0 {
		return -1, errors.New("args must be positive")
	}
	if n == 1 {
		return 0, nil
	}
	if n == 2 {
		return 1, nil
	}

	fib_1, _ := fib(n-1)
	fib_2, _ := fib(n-2)

	return fib_1 + fib_2, nil
}


func main() {
	result, err := fib(8)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
