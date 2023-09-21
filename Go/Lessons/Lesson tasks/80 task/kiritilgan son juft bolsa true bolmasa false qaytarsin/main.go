package main

import "fmt"

func main() {

	fmt.Println(oddeven(6))
}
func oddeven(in int) bool {
	if in%2 == 0 {
		return true
	} else {
		return false
	}
}