package main

import "fmt"

type account struct {
	number string 
	city string 
	balance float64 
}

// method uses value from account struct
func(a *account) HaveEnoughBalance(value float64) bool{
   if a.balance >= value{
      return true
   }
   return false
}

// Withdraw amount from account 
func (a *account) withdraw(value float64) {
	if a.HaveEnoughBalance(value) {
		y :=  &a.balance
		*y = a.balance - value
	}
}

// add amount to account 
func (a *account) add(value float64) {
	a.balance = a.balance + value
}




func main() {
   drew := account{"C21345345345355", "New York", 15470.09}

   fmt.Println(drew)
   drew.withdraw(500)
   drew.add(2000)
   drew.withdraw(3000000)

   fmt.Println(drew.balance)
}
