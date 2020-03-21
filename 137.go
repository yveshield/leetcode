package main

func singleNumber(nums []int) int {
	p := make(map[int]int)
	p[nums[0]] = 1
	for i := 1; i < len(nums); i++ {
		if k, ok := p[nums[i]]; ok {
			if k == 2 {
				delete(p, nums[i])
			} else {
				p[nums[i]]++
			}
		} else {
			p[nums[i]] = 1
		}
	}
	var q int
	for k := range p {
		q = k
	}

	return q
}
