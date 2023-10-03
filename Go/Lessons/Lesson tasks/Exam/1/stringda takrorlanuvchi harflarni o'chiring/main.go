package main

import (
	"fmt"
)

func main() {
	fmt.Println(deleteDuplicates("getthis"))
}
func deleteDuplicates(matn string) (new string)  {
	i := 0
	j := 0

	for j < len(matn) {

		if matn[i] == matn[j] {
			j+=1
		} else if matn[j]!= matn[i] || j == len(matn) -1 {
			new += string(matn[i])
			i = j
			j++
		} 
	}
	new += string(matn[j-1])
	return new
}