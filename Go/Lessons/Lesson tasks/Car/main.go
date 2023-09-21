package main

import "fmt"
type Car struct {
	Name string
	Model int
	RentPerDay float64
} 
func (d Car) display() {
	fmt.Println("_____________________________________")
	fmt.Println("Name:", d.Name)
	fmt.Println("------------------------------")
	fmt.Println("Model:", d.Model)
	fmt.Println("------------------------------")
	fmt.Println("Rent per day:", d.RentPerDay)
	fmt.Println("_____________________________________")
}
func (r *Car) day(in float64) {
	if in <= 0 {
		fmt.Println("0 yoki 0 dan kichik son qabil qilinmaydi")
		return
	}
	r.RentPerDay *= in
}
func main() {
	car1:= Car{
		Name: "NIX",
		Model: 1,
		RentPerDay: 30,
	}

	car1.display()
	car1.day(-3)
	car1.display()

}