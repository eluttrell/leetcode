package main

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

func main() {}
