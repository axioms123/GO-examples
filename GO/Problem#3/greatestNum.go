//Write a function with one variadic parameter that finds the 
//greatest number in a list of numbers.


package main 

import (
	"fmt"
)

func greatestNum(nums ...int) {
	largestNum := 0
	 for _, num := range nums {
		 if num > largestNum{
			 largestNum = num
		 }
    }
	fmt.Println(largestNum)
}

func main() {
	greatestNum(1, 2, 4)
}
