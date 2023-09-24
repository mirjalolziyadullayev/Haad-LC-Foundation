package main

import "fmt"

func main() {

	var a,b,c int

	fmt.Println("a soni uchun qiymat kiriting:")
	fmt.Scan(&a)

	fmt.Println("a soni uchun qiymat kiriting:")
	fmt.Scan(&b)
	
	fmt.Println("a soni uchun qiymat kiriting:")
	fmt.Scan(&c)

	result:= a+b+c
	fmt.Println("result is:",result)
}