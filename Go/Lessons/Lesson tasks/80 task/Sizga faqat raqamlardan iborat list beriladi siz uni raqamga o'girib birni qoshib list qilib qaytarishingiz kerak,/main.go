package main 

import "fmt"

func main() {
 	num :=[] int {1,2,3,5,6}
	var arrtos string
	var endindex int
	for i := 0; i < len(num); i++ {
		
		arrtos += fmt.Sprint(num[i])
		endindex++
	}


	fmt.Println(fmt.Sprint(arrtos[:endindex-1])+string(arrtos[endindex-1]+1))
}