package main 

import "fmt"

func main() {
 	num :=[] int {1,2,3}
	var arrtos string
	var endindex int
	for i := 0; i < len(num); i++ {
		endindex++
		arrtos += fmt.Sprint(num[i])
	}

	

	fmt.Println(arrtos)
}