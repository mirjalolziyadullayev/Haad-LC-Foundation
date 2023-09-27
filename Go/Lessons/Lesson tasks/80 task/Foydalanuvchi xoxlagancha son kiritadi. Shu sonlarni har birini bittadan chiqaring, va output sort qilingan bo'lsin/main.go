package main

import "fmt"

func main() {
	var get int = 12356
	var got [] int
	fmt.Println("Foydalanuvchi son kiritsin:")
	fmt.Scan(&get)

	for get > 0 {
		//boshidan qo'shib kelish uchun
		flag := []int{get % 10}
		got = append(flag, got...)

		//ohridan qoshadi
		// got = append(got, get%10)

		get = get / 10
		fmt.Println(get, got)
	}
	
	fmt.Println(got)
}