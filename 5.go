package main

func longestPalindrome(s string) string {
	begin := 0
	end := 0
	size := len(s)
	for i := 0; i < size; i++ {
		len1 := expand(s, i, i)
		if i < size-1 && s[i] == s[i+1] {
			len2 := expand(s, i, i+1)
			if len2 > len1 {
				len1 = len2
			}
		}
		if len1 > end-begin {
			begin = i - (len1-1)/2
			end = begin + len1
		}
	}
	return s[begin:end]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}
