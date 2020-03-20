package leetcode

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(s string, left, right int)

	dfs = func(s string, left, right int) {
		if left > right {
			return
		}

		if left == 0 && right == 0 {
			ans = append(ans, s)
			return
		}

		if left > 0 {
			dfs(s+"(", left-1, right)
		}
		if right > 0 {
			dfs(s+")", left, right-1)
		}
	}
	dfs("", n, n)
	return ans
}
