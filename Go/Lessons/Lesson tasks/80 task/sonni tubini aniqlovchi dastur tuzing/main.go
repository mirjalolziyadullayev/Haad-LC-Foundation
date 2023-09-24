package main

import (
	"fmt"
)

func main() {
	son:= 13
	count:= 1
	var result bool

	for i:= 1; i < son; i++ {
		if son%i==0{
			count++
		}
	}
	if count == 2{
		result = true
	} else {
		result = false 
	}
	fmt.Println(result)
}