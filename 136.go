package leetcode

func singleNumber(nums []int) int {
	p := make(map[int]int)
	p[nums[0]] = nums[0]
	for i := 1; i < len(nums); i++ {
		if _, ok := p[nums[i]]; ok {
			delete(p, nums[i])
		} else {
			p[nums[i]] = nums[i]
		}
	}
	var k int
	for _, v := range p {
		k = v
	}
	return k
}
