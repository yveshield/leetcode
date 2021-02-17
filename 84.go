package main

func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	lessFromLeft := make([]int, l)
	lessFromRight := make([]int, l)
	lessFromRight[l-1] = l
	lessFromLeft[0] = -1
	for i := 1; i < l; i++ {
		p := i - 1

		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}
		lessFromLeft[i] = p
	}

	for i := l - 2; i >= 0; i-- {
		p := i + 1

		for p < l && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}
		lessFromRight[i] = p
	}

	maxArea := 0
	for i := 0; i < l; i++ {
		max := heights[i] * (lessFromRight[i] - lessFromLeft[i] - 1)
		if max > maxArea {
			maxArea = max
		}
	}
	return maxArea
}