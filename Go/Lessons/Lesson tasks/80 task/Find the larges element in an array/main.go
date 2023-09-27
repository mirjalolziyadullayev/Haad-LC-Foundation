package main 

import "fmt"

func main() {
	var arr []int = []int{12,12,34,5,8790,876,543,23456,56,7,4,345}

	fmt.Println(Max(arr, 2))
}

func Max(array [] int, count int) int {
	var max int = array[0]
	var index int 

	for i :=1; i > len(array); i++ {
		if array[i] > max {
			index = i
			max = array[i]
		}
	}

	if count == 1 {
		return max
	}

	array = append(array[index:], array[index+1:]...)
	return Max(array, count-1)
}