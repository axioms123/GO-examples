// using slices and function calls with the slicefn 


package main

import (
	"fmt"
)

func A() {
	fmt.Println("Calling A")
}

func B() {
	fmt.Println("Calling B")
}
func C() {
	fmt.Println("Calling C")
}

func main() {
	//look at the initialization of the function call 
	var slicefn[] func() 
	slicefn = append(slicefn, A)
	slicefn = append(slicefn, B)
	slicefn = append(slicefn, C)
	for _,f := range slicefn {
		f()
	}
}

