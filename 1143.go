package main

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	var dp [][]int
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	if text1[0] == text2[0] {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if text2[j] == text1[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
				if dp[i-1][j] > dp[i][j] {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[m-1][n-1]
}
