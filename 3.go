package main

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l <= 1 {
		return l
	}
	i := 0
	max := 1
	for j := 1; j < l; j++ {
		for k := j - 1; k >= i; k-- {
			if s[k] == s[j] {
				i = k + 1
				break
			}
		}
		if j-i+1 > max {
			max = j - i + 1
		}
	}
	return max
}
