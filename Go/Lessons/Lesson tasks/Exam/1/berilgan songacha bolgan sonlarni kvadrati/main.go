package main

import "fmt"

func main() {
	get := 5
	index := 0
	var result int 
	for index < get {

		result = index * (get * index)
		fmt.Println(result)
		index++	
	}
}
