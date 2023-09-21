package main

import (
	"fmt"
)

func main() {
	fmt.Println(reversedInteger([]int{1,2,3,4,5,6}))
}

func reversedInteger(chart[] int) (reversed[] int) {
	
	for i := len(chart)-1; i >= 0; i-- {
		reversed = append(reversed, chart[i])
	}
	return reversed
}