package main 

import (
	"fmt"
)

//func main() {
//	s := make([]int, 1,4)
//	for i:=0; i<=8; i++ {
//		fmt.Printf("%d, %d\n", len(s), cap(s))
//		s = append(s,i)	}
//	fmt.Scanln("what is your name");
//	//var name string = &input 
//}

func main() {
	var threeDimensionalArray = make([][][]int, 3)
	var array1 = make([]int,3)
	var array2 = make([]int,3)
	var array3 = make([]int,3)
	println("Enter three ints: ");
	for i:=0; i<3; i++ {
		fmt.Scanf("%d", &array1[i]);
	}
	println("Enter the next three ints: ");
	for i:=0; i<3; i++ {
		fmt.Scanf("%d", &array2[i]);
	}
	println("Enter the last three ints: ");
	for i:=0; i<3; i++ {
		fmt.Scanf("%d", &array3[i]);
	}
	threeDimensionalArray = append(threeDimensionalArray, array1)
	threeDimensionalArray = append(threeDimensionalArray, array2)
	threeDimensionalArray = append(threeDimensionalArray, array3)
}
