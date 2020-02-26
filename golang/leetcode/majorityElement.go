package leetcode

func MajorityElement(nums []int) int {
	res := nums[0]
	c := 0
	for index := 0; index < len(nums); index++ {
		if res == nums[index] {
			c++
		} else {
			c--
		}
		if c == 0 {
			res = nums[index+1]
		}
	}

	return res
}
