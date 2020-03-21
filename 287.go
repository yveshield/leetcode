package main

func findDuplicate(nums []int) int {
	left := 1
	right := len(nums) - 1
	for left < right {
		mid := (left + right) >> 1
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
