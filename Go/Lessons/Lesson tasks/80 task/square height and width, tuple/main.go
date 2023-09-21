package main

import (
	"fmt"
)

func main() {

	str:="()"

	fmt.Println( string( str[0]) + fmt.Sprint( perimeter(4,6), ",", surface(4,6) )  + string(str[1]) )

}
func perimeter(height int, width int) (result int) {
	result = (height * 2) + (width * 2)
	return result
}
func surface(height int, width int) (result int) {
	result = height + width
	return result
}