package main

func coinChange(coins []int, amount int) int {
	IntMax := int(^uint(0) >> 1)
	dp := make([]int, amount+1)
	for j := 1; j <= amount; j++ {
		dp[j] = IntMax
		for i := 0; i < len(coins); i++ {
			if coins[i] <= j && dp[j-coins[i]] < IntMax && (dp[j-coins[i]]+1) < dp[j] {
				dp[j] = dp[j-coins[i]] + 1
			}
		}
	}
	if dp[amount] == IntMax {
		dp[amount] = -1
	}
	return dp[amount]
}
