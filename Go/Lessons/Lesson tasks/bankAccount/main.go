package main

import "fmt"
type bankAccount struct {
	accountID float32
	holderName string
	accountBalance float64
}
func main() {
	res := bankAccount{
		accountID: 1000,
		holderName: "Mirjalol",
		accountBalance: 10.100,
	}
	res.display()
}

func (show bankAccount) display() {

	fmt.Println("AccountID is:", show.accountID)
	fmt.Println("Holder name is:", show.holderName)
	fmt.Println("Account balance is:", show.accountBalance)
}
func (depozit *bankAccount) depozit() {
	
	depozit += depozit.accountBalance
}