// package subfolder 
// import (
// 	"fmt"
// )

// //Public Access 
// func DemoFunction() {
// 	fmt.Println("I am from inside outside the main")

// }

// func demofunction() {
// 	fmt.Println("not callable from the main")
// }


package main import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
        
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.First, p1.person.First)
	fmt.Println(p2.First, p2.person.First)
}

3)input
1 2 3 
4 5 6
7 8 9
output(interchange the first and last coloumn)
3 2 1
6 5 4
9 8 7
-------------------------------
input
1 2 3 
4 5 6
7 8 9
output(interchange the first and last row )
7 8 9
4 5 6
1 2 3
*/

