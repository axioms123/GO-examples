//package main 

//import (
//	"fmt" 
//	"math"
//	"reflect"
//	"unsafe"
//)


package main

import "fmt"

func main() {
    var a int = 10
    var b int32 = 20
   // var c byte = 15
   // var d float32 = 0.05

	var x int32 = int32(a) + b
	fmt.Println(x)

	
    // fmt.Println(b + c) // int32 & byte
    // fmt.Println(a + c) // int & byte
    // fmt.Println(a + d) // int & float32
    // fmt.Println(b * d) // int32 & float32
   // fmt.Println(c / d) // byte & float21
}
