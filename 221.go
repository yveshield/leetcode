package main

func maximalSquare(matrix [][]byte) int {
	if matrix == nil {
		return 0
	}
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])
	var dp [][]int
	max := 0
	for i := 0; i < row; i++ {
		dd := make([]int, col)
		dp = append(dp, dd)
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					min := dp[i-1][j-1]
					if dp[i-1][j] < min {
						min = dp[i-1][j]
					}
					if dp[i][j-1] < min {
						min = dp[i][j-1]
					}
					dp[i][j] = min + 1
				}
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max
}
