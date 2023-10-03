package main

import "fmt"

func main() {
 	fmt.Println(get(10, 10))
}

func get(decimal int, base int) (res string) {
 	letters := "0123456789abcdefghijklmnopqrstuwxyz"

	for decimal > 0 {

	res = string(letters[decimal%base]) + res	
	decimal = decimal / base
}
	return res
}