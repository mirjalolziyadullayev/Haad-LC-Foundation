package main 

import "fmt"

func main() {
	var integer int = 23 // integer can identifiy the bits without specific characters 
	fmt.Println(integer)

	var integer1 int8 = -128 // integer with specific character means the bits of the getting numbers, and in this example used int8 wich can get only signed 8 bit numbers
	fmt.Println(integer1)
}