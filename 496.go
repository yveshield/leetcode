package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	p := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		p[i] = -1
		for j := 0; j < len(nums2)-1; j++ {
			if nums2[j] == nums1[i] {
				for k := j + 1; k < len(nums2); k++ {
					if nums2[k] > nums1[i] {
						p[i] = nums2[k]
						break
					}
				}
				break
			}
		}
	}
	return p
}
