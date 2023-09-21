package main

import (
	"fmt"
)

func main() {
	getList([]int{1,3,5,6,8,9})
}
func getList(list []int) {

	var reversed []int
	var inttostring string
	var index = 0
	for i := len(list)-1; i >=0; i-- {
		
		reversed = append(reversed, list[i])
		inttostring += fmt.Sprint(reversed[index])
		index++
	}

	fmt.Println("(", inttostring, ")")
}


