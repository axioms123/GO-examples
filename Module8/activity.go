
package main

import (
	"fmt"
	"time"
)

//Activity 1 

func customTime () {
	// Current date and time
   currentTime := time.Date(
      2021, 07, 15, 15, 20, 00, 0, time.Local)
   fmt.Println("Custom date and time:", currentTime)
   
   // Current year
   fmt.Println("Current Year", time.Now().Year())
   
   // Current month
   fmt.Println("Current Month", time.Now().Month())
   

   // Week number of the year 
   fmt.Println("Week number of the year", time.Now().Year())
   
   // Weekday of the week
   var _, wk_num = time.Now().ISOWeek() 
   fmt.Println(wk_num)


   // Day of the year

   
   // Day of the month



   // Day of the week
}


func main() {

	customTime()

}





// 2 Write a function that prompts the user for a year and checks if it is a leap year. The function should return True if the input is a leap year and False otherwise.

