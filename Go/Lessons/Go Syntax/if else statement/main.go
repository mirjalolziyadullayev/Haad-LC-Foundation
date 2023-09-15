package main

import "fmt"

func main() {
	var degree int = 100

	if degree >= 100 {
		fmt.Println("Water will boil!")
	} else if degree > 0 {
		fmt.Println("Water will not freeze!")
	} else {
		fmt.Println("Water will freeze")
	}
}