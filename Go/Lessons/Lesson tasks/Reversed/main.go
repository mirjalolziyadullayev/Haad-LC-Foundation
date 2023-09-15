package main

import (
	"fmt"
)

func main() {
	fmt.Println(reversedString("bobur"))
}

func reversedString(chart string) (reversed string) {
	for i := len(chart)-1; i >= 0; i-- {
		reversed = reversed + string(chart[i])
	}
	return reversed
}