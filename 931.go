package main

func minFallingPathSum(A [][]int) int {
	row := len(A)
	col := len(A[0])
	IntMax := int(^uint(0) >> 1)
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			tmp := A[i][j]
			A[i][j] = IntMax
			for k := j - 1; k <= j+1; k++ {
				if k < 0 || k > col-1 {
					continue
				}
				if A[i-1][k]+tmp < A[i][j] {
					A[i][j] = A[i-1][k] + tmp
				}
			}
		}
	}
	min := A[row-1][0]
	for l := 0; l < col; l++ {
		if A[row-1][l] < min {
			min = A[row-1][l]
		}
	}
	return min
}
