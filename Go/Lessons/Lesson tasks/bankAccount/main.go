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
		accountBalance: 10,
	}
	res1:= bankAccount{
		accountID: 1001,
		holderName: "Habib",
		accountBalance: 20,
	}

	res.deposit(100)
	res.display()
	
	res1.deposit(100)
	res1.display()
	res1.withdraw(120)
	res1.display()

}
func (show bankAccount) display() {
	fmt.Println("_________________________________________________\n")
	fmt.Println("   Account ID is:", show.accountID)
	fmt.Println("------------------------------")
	fmt.Println("   Holder name is:", show.holderName)
	fmt.Println("------------------------------")
	fmt.Println("   Account balance is:", show.accountBalance)
	fmt.Println("_________________________________________________\n")
}
func (dep *bankAccount) deposit(in float64) {
	if in < 0 {
		fmt.Println("\nManfiy qiymat kiritib bolmaydi!")
		return
	}
	dep.accountBalance += in
}
func (dep *bankAccount) withdraw(out float64) {
	if out >= dep.accountBalance{
		fmt.Println("Mablag' yetarli emas")
		return
	}
	dep.accountBalance -= out
}
