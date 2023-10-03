package main

import "fmt"

func commonElements(arr1, arr2 []int) ( elements [] int) {
	for i :=0; i < len(arr1); i++ {
		fmt.Println(i)
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				elements = append(elements, arr1[i])	
			}
		}
	} 
	
	return elements
}

func main() {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{4, 5, 6, 7, 8}

	common := commonElements(array1, array2)
	fmt.Println(common)
}