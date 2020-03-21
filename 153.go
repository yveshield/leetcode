package main

func findMin(nums []int) int {
	v := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < v {
			v = nums[i]
		}
	}
	return v
}
