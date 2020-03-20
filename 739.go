package leetcode

func dailyTemperatures(T []int) []int {
	p := make([]int, len(T))
	for k := range p {
		p[k] = 0
	}
	for i := 0; i < len(T); i++ {
		for j := i + 1; j < len(T); j++ {
			if T[i] < T[j] {
				p[i] = j - i
				break
			}
		}
	}

	return p
}
