//The Fibonacci sequence is defined as: 0, 
//fib(0) = fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). 
//Write a recursive function which can find fib(n).

package main

import (
	"fmt"
)




func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
