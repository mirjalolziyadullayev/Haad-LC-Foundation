package main

import (
	"fmt"
)

type Book struct {
	Name string
	Author string
	Price float32
	Currency string
}
func (b Book) display() {
	fmt.Println("______________________________\n")
	fmt.Println("Book name is:", b.Name)
	fmt.Println("-----------------------")
	fmt.Println("Author:", b.Author)
	fmt.Println("-----------------------")
	fmt.Println("Price:", b.Currency + fmt.Sprint(b.Price))
	fmt.Println("______________________________\n")

}
func (d *Book) discount(in float32) {
	if in  > 100 {
		fmt.Println("bunaqa son kirita olmaysiz")
		return
	}
	d.Price -= (d.Price * in / 100)
}
func main() {
	book1:= Book{
		Name: "The Science of Interstellar",
		Author: "Kip Thorne",
		Price: 18,
		Currency: "$",
	}

	book1.display()
	book1.discount(10)
	book1.display()
}
