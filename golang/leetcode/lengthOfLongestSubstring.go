package leetcode

func LengthOfLongestSubstring(s string) int {
	l := len(s)
	temp := map[string]int{}
	start, end := 0, 0
	max := 0
	tempChar := ""
	for end < l {
		tempChar = string(s[end])
		existCharIndex, endExist := temp[tempChar]
		if endExist && existCharIndex >= start {
			start = existCharIndex + 1
		}
		temp[tempChar] = end
		end++
		if max < end - start {
			max = end - start
		}
	}
	return max;
}
