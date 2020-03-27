package leetcode

import "math"

func TitleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += (int(s[i]) - 64) * int(math.Pow(float64(26), float64(len(s) - i - 1)))
	}

	/*
	ABC
	1
	1 * 26^1 + 2*26^0
	(1*26 + 2) * 26 + 3 = 1*26^2 + 2*26^1 + 3*26^0
	*/

	return res
}