package main

import "fmt"


func main() {
	var kg, price, discount float64

	fmt.Println("Konfet miqdorini kiriting kg:")
	fmt.Scan(&kg)
	fmt.Println("Konfet narxini kiriting:")
	fmt.Scan(&price)
	fmt.Println("chegirma % ni kiriting:")
	fmt.Scan(&discount)

	var sum float64

	if kg > 10 {
		sum = price * 10 + (kg - 10) * price*discount/100
	} else {
		sum = price * 2
	}
	fmt.Println(sum)
}