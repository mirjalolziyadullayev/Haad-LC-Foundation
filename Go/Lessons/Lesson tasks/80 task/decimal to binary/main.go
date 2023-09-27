package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	
	res := DecimalToBinary(num)
	fmt.Println(res)
}
func DecimalToBinary(num int) string {
	var res string 
	for num > 0 {
		flag := num % 2
		res = fmt.Sprint(flag) + res
		num = num / 2
	}
	return res
}