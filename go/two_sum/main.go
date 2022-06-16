package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	indices := []int{}
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		ct := target - nums[i]
		if val, ok := hash[ct]; ok {
			return []int{val, i}
		}
		hash[nums[i]] = i
	}
	return indices
}

func main() {
	x := []int{1, 2, 3, 4, 14, 19, 10, 7}
	fmt.Println(TwoSum(x, 9))
}
