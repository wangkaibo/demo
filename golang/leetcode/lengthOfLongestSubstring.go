package leetcode

func LengthOfLongestSubstring(s string) int {
	l := len(s)
	temp := map[string]int{}
	start, end := 0, 0
	max := 0
	tempChar := ""
	for end < l {
		tempChar = string(s[end])
		endIndex, endExist := temp[tempChar]
		if endExist && endIndex >= start {
			start = endIndex + 1
		}
		temp[tempChar] = end
		end++
		if max < end - start {
			max = end - start
		}
	}
	return max;
}