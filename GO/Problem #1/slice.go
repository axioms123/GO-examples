//1)sum is a function which takes a slice of numbers and adds them together.
//What would its function signature look like in Go?

package main

import (
	"fmt"
)

func sumofDigits(number int) int {
	remainder := 0
	sum := 0
	for number != 0 {
		remainder = number % 10
		sum += remainder
		number = number / 10
	}
	return sum
}

func main() {
	var number int
	fmt.Print("Enter Number:")
	fmt.Scanln(&number)
	fmt.Printf("The sum of digits of %d = %d\n", number, sumofDigits(number))
}
