package main

func hIndex(citations []int) int {
	l := len(citations)
	h := 0
	for i := l - 1; i >= 0; i-- {
		if citations[i] >= l-i {
			h = l - i
		} else {
			break
		}
	}
	return h
}
