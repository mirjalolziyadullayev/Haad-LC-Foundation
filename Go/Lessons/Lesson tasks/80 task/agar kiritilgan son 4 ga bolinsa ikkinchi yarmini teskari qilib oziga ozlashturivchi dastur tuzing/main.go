package main

import (
	"fmt"
)

func main() {
	var get string
	fmt.Println("string kiriting:")
	fmt.Scan(&get)
	if len(get)%4 == 0 {
		var str string
        
		middle := len(get) / 2
        
		str = str + get[middle:]
        
		fmt.Println(str)
        
		got := get[:middle]
        
		fmt.Println(string(got)+ reversedString(str))
		
	}
}

func reversedString(str string) (reversed string) {
	for i := len(str)-1; i >= 0; i-- {
		reversed = reversed + string(str[i])
	}
	return reversed
}