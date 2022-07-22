
//Create a slice of 20 of type int and take 20 number create
//4 goroutines to take average of every 5 elements
//ex[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20] ,
//so 1 goroutine should give average of first 5 elements
//and next for another 5 so on

package main

import (
	"fmt"

)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
