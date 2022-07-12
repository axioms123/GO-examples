// Write a function which takes an integer and halves it and returns true if
// it was even or false if it was odd. For example half(1) should return (0, false)
// and half(2) should return (1, true).

package main

import (
	"fmt"
)

func halfOfNumber(number int) bool {
	var half int = number / 2
	if half%2 == 0 {
		return true
	} else {
		return false

	}

}

func main() {
	var number int
	fmt.Println("Enter a number ")
	fmt.Scanln(&number)
	fmt.Printf(" It is %d that %d\n is an even number ", halfOfNumber(number), number)
}
