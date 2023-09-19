package main

import "fmt"

func main() {
	fmt.Println(oddNums(16))
}
func oddNums(n int) [] int {
	var res []int
	count :=1

	for len(res) < n {
		res = append(res, count)
		count = count * 2
	}
	return res
}