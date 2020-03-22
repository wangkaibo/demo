package leetcode

import "math"

func Reverse(x int) int {
	res := 0
	for x != 0 {
		res = 10 * res + x % 10
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x = x / 10
	}

	return res
}