package main

func nextGreaterElements(nums []int) []int {
	p := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		found := false
		p[i] = -1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				p[i] = nums[j]
				found = true
				break
			}
		}
		if !found {
			for k := 0; k < i; k++ {
				if nums[k] > nums[i] {
					p[i] = nums[k]
					break
				}
			}
		}

	}
	return p
}
