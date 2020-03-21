package main

func hIndex(citations []int) int {
	citations = Shell(citations)
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

func Shell(items []int) []int {
	n := len(items)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			j := i
			for j >= h && items[j] < items[j-h] {
				items[j], items[j-h] = items[j-h], items[j]
				j -= h
			}
		}
		h /= 3
	}
	return items
}
