package main

func minDistance(word1 string, word2 string) int {
	var dp [][]int
	len1 := len(word1)
	len2 := len(word2)
	for i := 0; i <= len1; i++ {
		dp = append(dp, make([]int, len2+1))
	}
	for i := 1; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1] + 1
				if dp[i-1][j]+1 < dp[i][j] {
					dp[i][j] = dp[i-1][j] + 1
				}
				if dp[i-1][j-1]+1 < dp[i][j] {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
		}
	}
	return dp[len1][len2]
}
