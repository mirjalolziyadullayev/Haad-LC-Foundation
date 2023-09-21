package main

import "fmt"

type rectangleGeometry struct {
	width int
	height int
}
func (p *rectangleGeometry) perimeter() {
	resP:= (p.width * 2) + (p.height * 2)
	fmt.Println(resP)
}
func (s *rectangleGeometry) Surface() {
	resS:= s.width * s.height 
	fmt.Println(resS)
}
func main() {
	object:= rectangleGeometry {
		width: 7,
		height: 6,
	}
	
	object.Surface()
	object.perimeter()
}