package leetcode

func nextGreaterElement(n int) int {
	p := []int{}
	for n > 0 {
		p = append(p, n%10)
		n = n / 10
	}
	found := false
	for i := 0; i < len(p); i++ {
		for j := 0; j < i; j++ {
			if p[i] < p[j] {
				found = true
				p[i], p[j] = p[j], p[i]
				il := i - 1
				for s := 0; s <= il/2; s++ {
					if p[s] < p[il-s] {
						p[s], p[il-s] = p[il-s], p[s]
					}
				}
				break
			}
		}
		if found {
			break
		}
	}
	if found {
		max := 1 << 31
		l := len(p)
		r := 0
		for t := l - 1; t >= 0; t-- {
			r = r*10 + p[t]
			if r > max {
				return -1
			}
		}
		return r
	}

	return -1
}
