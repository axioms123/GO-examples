// Upside Down Triangle with Loops
package main

import "fmt"

func main() {
	for i := 7; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Print("\n")
	}
}
