package main

import "fmt"

func main() {
	plusminus(0)
}
func plusminus(in int) {
	if in>0 {
		fmt.Println("true")
	} else if in == 0{
		fmt.Println("musbatham manfiyham emas")
	} else {
		fmt.Println("false") 
	}
}