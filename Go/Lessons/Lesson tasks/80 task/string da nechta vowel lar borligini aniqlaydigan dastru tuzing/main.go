package main

import (
	"fmt"
	"strings"
)

func main() {
	var matn string

	fmt.Println("matn kiriting:")
	fmt.Scan(&matn)

	var count int
	for i:=len(matn)-1; i >= 0; i-- {

		if strings.Contains("oeiau", string(matn[i])) {
			count++
		}

	}
	fmt.Println(matn, count)
}
