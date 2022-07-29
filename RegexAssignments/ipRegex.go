/* Assignment 2) Search for IP address of range 192.160.1 - 192.170.1
192.160.1.1    - 192.170.1.255
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	ip := regexp.MustCompile(`192([1][6]\d+).1\b([01]?[0-9][0-9]?|2[0-4][0-9]|25[0-5])`)

	string := "This is the IP address is 192.165.1.1"
	fmt.Printf("%q\n", ip.FindAllString(string, -1))

}










