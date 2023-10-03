package main

import "fmt"

func main() {

		var a, b, c  int= 3,2,1

		a = a + b;
		b = a - b;
		a = a - b;
 
		b = b + c;
		c = b - c;
		b = b - c;

		fmt.Println(a,b,c)
}
