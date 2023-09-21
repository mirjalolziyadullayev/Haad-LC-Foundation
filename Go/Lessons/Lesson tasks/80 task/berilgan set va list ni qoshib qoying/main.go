package main

import "fmt"

func main() {
	set1:= []string{"blue", "yellow", "black"}
	set2:= []string{"green", "maroon","brown"}


	for i := 0; i < len(set1); i++ {
		set2 = append(set2, set1[i])
	}
	fmt.Println(set2)
}