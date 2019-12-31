package demo

func singleNumber(nums []int) int {
	res := nums[0]
	for index := 1; index < len(nums); index++ {
		res = res ^ nums[index]
	}

	return res
}
