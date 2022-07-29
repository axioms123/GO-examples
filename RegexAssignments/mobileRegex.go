//Assignment 1) Find a valid mobile no 10 digit long and begin with your country code/


package main

import (
	"fmt"
	"regexp"

)

func main() {
	number := regexp.MustCompile(`^[\+?1]?\nnn*$`)
	phoneNumber := regexp.MustCompile(`[^0-9]*1[34578][0-9]{9}[^0-9]*`)

	string := "my phone number is found below 13435764545"
	fmt.Printf("%q\n", phoneNumber.FindAllString(string, -1))
	fmt.Printf("%q\n", number.FindAllString(string, -1))

}
