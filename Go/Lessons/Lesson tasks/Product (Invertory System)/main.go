package main

import "fmt"

type Product struct {
	Name string
	Quantity int
	Price float32
}
func (d Product) display() {
	fmt.Println("__________________________________________\n")
	fmt.Println("Name:", d.Name)
	fmt.Println("--------------------------------")
	fmt.Println("Quantity:", d.Quantity)
	fmt.Println("--------------------------------")
	fmt.Println("Price:", d.Price)
	fmt.Println("__________________________________________\n")
}
func (inc *Product) incQuantity(in int) {
	if in <= 0 {
		fmt.Println("0 va 0 dan kam miqdor kiritib bolmaydi")
		return
	}
	inc.Quantity += in
}
func (inc *Product) decQuantity(in int) {
	if in <= 0 {
		fmt.Println("0 va 0 dan kam miqdor kiritib bolmaydi")
		return
	}
	inc.Quantity -= in
}
func main() {
	product1:= Product{
		Name: "Thing",
		Quantity: 10,
		Price: 300,
	}

	product1.display()
	product1.decQuantity(23)
	product1.display()
}