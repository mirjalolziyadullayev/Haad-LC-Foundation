package main

import "fmt"

func main() {
	//slice declaring
	slc := []int{0,1,2,3,4,5,6}
	slc = append(slc, 7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22) 

	fmt.Println(slc, len(slc), cap(slc))
}