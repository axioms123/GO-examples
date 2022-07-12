package main

import (
	"fmt"
	"math"
)

func main() {

	var number float64
	fmt.Println("Enter your the number: ")
	fmt.Scanf("%s", &number)
	fmt.Println(int(math.Round(number)))
}


