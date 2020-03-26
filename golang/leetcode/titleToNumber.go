package leetcode

import "math"

func TitleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += (int(s[i]) - 64) * int(math.Pow(float64(26), float64(len(s) - i - 1)))
	}

	return res
}