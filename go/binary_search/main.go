package main

func binarySearch(nums []int, target int) int {
	lastIdx := len(nums) - 1
	if target < nums[0] || target > nums[lastIdx] {
		return -1
	}
	low := 0
	high := lastIdx
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}
		if low > high {
			return -1
		}
	}
	return -1
}

func main() {}
