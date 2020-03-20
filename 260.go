package leetcode

func singleNumber(nums []int) []int {
	p := make(map[int]int)
	p[nums[0]] = 1
	for i := 1; i < len(nums); i++ {
		if _, ok := p[nums[i]]; ok {
			delete(p, nums[i])
		} else {
			p[nums[i]] = 1
		}
	}
	var q []int
	for k := range p {
		q = append(q, k)
	}
	return q
}
