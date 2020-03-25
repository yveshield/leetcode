package main

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	if s[0] == '0' {
		dp[1] = 0
	}
	if len(s) <= 1 {
		return dp[1]
	}
	for i := 2; i <= len(s); i++ {
		a := s[i-2] - '0'
		b := s[i-1] - '0'
		n := a*10 + b
		// 至少一个0独处，无法解码
		if s[i-1] == '0' && s[i-2] == '0' {
			return 0
		} else if s[i-2] == '0' { // i-1必需独立，增加i位时，组合数不变
			dp[i] = dp[i-1]
		} else if s[i-1] == '0' {
			if n > 26 {
				return 0
			}
			dp[i] = dp[i-2] //abc10+x，10和x必需独立
		} else if n > 26 { //abc34+x
			dp[i] = dp[i-1]
		} else { //abc23+x
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[len(s)]
}
