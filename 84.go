package main

func largestRectangleArea(heights []int) int {
	l := len(heights)
	max := 0
	for i := 0; i < l; i++ {
		area := 0
		for j := i; j >= 0; j-- {
			if heights[j] >= heights[i] {
				area += heights[i]
			} else {
				break
			}
		}
		for j := i + 1; j < l; j++ {
			if heights[j] >= heights[i] {
				area += heights[i]
			} else {
				break
			}
		}
		if area > max {
			max = area
		}
	}
	return max
}
