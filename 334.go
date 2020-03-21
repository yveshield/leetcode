package main

const MaxInt = int(^uint(0) >> 1)

func increasingTriplet(nums []int) bool {
	m1 := MaxInt
	m2 := MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] <= m1 {
			m1 = nums[i]
		} else if nums[i] <= m2 {
			m2 = nums[i]
		} else {
			return true
		}
	}
	return false
}
