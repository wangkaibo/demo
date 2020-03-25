package leetcode

import (
	"math"
)

func LengthOfLongestSubstring(s string) int {
	l := len(s)
	temp := map[string]int{}
	res := 0;
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			temp = map[string]int{}
			unique := true
			for k := i; k < j; k++ {
				_, ok := temp[string(s[k])]
				if !ok {
					temp[string(s[k])] = 1
				} else {
					unique = false
				}
			}
			if unique {
				res = int(math.Max(float64(res), float64(j - i)))
			}
		}
	}
	return res;
}
