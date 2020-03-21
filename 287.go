package main

func findDuplicate(nums []int) int {
	dp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := dp[nums[i]]; ok {
			return nums[i]
		}
		dp[nums[i]] = 1
	}
	return 0
}
