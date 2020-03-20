package leetcode

func combinationSum(candidates []int, target int) [][]int {
	Shell(candidates)
	res := [][]int{}
	dfs(candidates, nil, target, 0, &res) //深度优先
	return res
}

func dfs(candidates, nums []int, target, left int, res *[][]int) {
	if target == 0 { //结算
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := left; i < len(candidates); i++ { // left限定，形成分支
		if target < candidates[i] { //剪枝
			return
		}
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i, res) //分支
	}
}

func Shell(items []int) []int {
	n := len(items)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			j := i
			for j >= h && items[j] < items[j-h] {
				items[j], items[j-h] = items[j-h], items[j]
				j -= h
			}
		}
		h /= 3
	}
	return items
}
