package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indices := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indices[0] = i
				indices[1] = j
			}
		}
	}
	return indices
}

func main() {
	x := []int{1, 2, 3, 4, 14, 19, 10, 7}
	fmt.Println(twoSum(x, 9))
}
