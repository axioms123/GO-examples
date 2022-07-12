package main

import {
	"fmt"
    "net/http"
	"example.com/greetings"
}


// Define type of response 
type Hello struct{}

//let that type implement the ServeHTTP method (defined in the interface http.Handler)
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}


func main() {
	var h Hello 
	http.ListenAndServe("localhost: 4000", h)
}


 
