package main

import "fmt"

func main() {
	//this is the program that calculates the surface 
	fmt.Println(zeroFunc(10, 5))

}
func zeroFunc(width int , height int)/*<-intput*/ (string, int)/*<-output*/{

	//function have input and output from itself 

	var result int
	result = width * height

	return string("the surface is:"), result

	//default value of the return is 0
}