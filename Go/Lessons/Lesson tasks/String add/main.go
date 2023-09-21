package main

import (
	"fmt"
)

func main() {
	p2:="THEREALMAN"

	var center int
	center = len(p2) /2

	fmt.Println( string( p2[:center] ) + "slice" + string( p2[center:] ) )
}