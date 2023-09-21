package main

import (
	"fmt"
)

type Libriary struct {
	Name string
	Books []string 
	Members []string
}

var unavialable string
var index int

func (d Libriary) display() {
	fmt.Println("________________________________________\n")
	fmt.Println("Library:", d.Name)
	fmt.Println("----------------------------")	
	fmt.Println("Books:", d.Books)
	fmt.Println("----------------------------")
	fmt.Println("Member:", d.Members)
	fmt.Println("________________________________________\n")
}

func (a *Libriary) addBook(in string) {
	a.Books = append(a.Books, in)
}

func (l *Libriary) lendBook(in int) {
	index = in-1
	unavialable = string(l.Books[in])

	l.Books = append(l.Books[:in], l.Books[in+1:]...)
}

func (av Libriary) availableBook(in string) bool {



	if av.Books[index] == unavialable{
		fmt.Println("not available")
		return false

	} else {
		fmt.Println("available")
		return true
	}
}

func main() {

	Libriary1:= Libriary {
		Name: "BBC Libriary",
		Books: []string{"harry potter","the science of interstellar"},
		Members: []string{"Kevin","Peter"}, 
	}
	
	Libriary1.display()
	Libriary1.addBook("Kip Thorne's life")
	Libriary1.display()
	Libriary1.lendBook(2)
	Libriary1.display()
	Libriary1.availableBook("Kip Thorne's life")

}