package main

func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	//col := i+1, 行数加1
	IntMax := int(^uint(0) >> 1)
	for i := 1; i < row; i++ {
		col := i + 1
		for j := 0; j < col; j++ {
			tmp := triangle[i][j]
			triangle[i][j] = IntMax
			for k := j - 1; k <= j; k++ {
				//当行最大下标为col-1，上一行更少一项，为col-2
				if k < 0 || k > col-2 {
					continue
				}
				if triangle[i-1][k]+tmp < triangle[i][j] {
					triangle[i][j] = triangle[i-1][k] + tmp
				}
			}
		}
	}
	min := triangle[row-1][0]
	col := row
	for l := 0; l < col; l++ {
		if triangle[row-1][l] < min {
			min = triangle[row-1][l]
		}
	}
	return min
}
