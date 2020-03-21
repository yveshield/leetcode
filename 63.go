package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				if obstacleGrid[i][j] == 1 {
					return 0
				}
				obstacleGrid[i][j] = 1
			} else if i == 0 {
				if obstacleGrid[i][j] == 1 || obstacleGrid[i][j-1] == 0 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = 1
				}
			} else if j == 0 {
				if obstacleGrid[i][j] == 1 || obstacleGrid[i-1][j] == 0 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = 1
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
				}
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}
