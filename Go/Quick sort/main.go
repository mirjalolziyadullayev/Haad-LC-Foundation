package main

import "fmt"

func main() {

	fmt.Println(quickSort([]int{12, 53, 23, 5, 64, 7, 78, 96, 5, 535, 36, 364, 7, 34, 3, 32, 6}))

}

func quickSort(nums []int) []int {

	if len(nums) == 1 || len(nums) == 0 {
		return nums
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return []int{nums[1], nums[0]}
		} else {
			return []int{nums[0], nums[1]}
		}
	}

	pivot := nums[0] // 12

	count := 1
	small := []int{}
	big := []int{}
	for count < len(nums) {
		if nums[count] <= pivot {
			small = append(small, nums[count])
		} else {
			big = append(big, nums[count])
		}
		count++
	}
	fmt.Println(small, pivot, big)

	s := quickSort(small)
	b := quickSort(big)
	s = append(s, pivot)

	return append(s, b...)

}
