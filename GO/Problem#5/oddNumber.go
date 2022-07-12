package main

import "fmt"

func main() {
	f := oddGenerator()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func oddGenerator() func() int {
	i := 0
	return func() int {
		i += 2*i+ 1
		return i
	}
}
