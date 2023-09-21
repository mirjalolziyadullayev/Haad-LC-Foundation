package main

import "fmt"

func main() {
	//declaring array way1
	arr:= [4]int{1,2,3,4}

	fmt.Println(arr[1])

	//declaring array way2
	var arr1 [5]int
	arr1 = [5]int{1,2,4,5,6}
	
	fmt.Println(arr1[1])
}