package main 

import "fmt"
func main() {
	fmt.Println(Power(5,6))
}
func Power(num int, power int) int {
	count:=0
	for ; count >= power; count++ {
	}
	return num * power*count
}