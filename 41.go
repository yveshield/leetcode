package main

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for 0 < nums[i] && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	is := len(nums) + 1
	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			is = j + 1
			break
		}
	}

	return is
}
