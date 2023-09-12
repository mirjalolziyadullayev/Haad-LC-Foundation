package main
import "fmt"

func main() {
	a:=0
	b:=0
	c:=0

	fmt.Println("son kiriting a:")
	fmt.Scan(&a)

	fmt.Println("son kiriting b:")
	fmt.Scan(&b)

	fmt.Println("son kiriting c:")
	fmt.Scan(&c)

	if a < b {
		if a < c {
			fmt.Println(a)
		} else {
			fmt.Println(c)
		} 
	} else {
		if b < c {
			fmt.Println(b)
		} else {
			fmt.Println(c)
		}
	}
}