package leetcode

func RomanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	if len(s) == 1 {
		return m[s]
	}
	i := 0
	res := 0
	for i < len(s) {
		if i == len(s) - 1 {
			res += m[string(s[i])]
		} else {
			a := m[string(s[i])]
			b := m[string(s[i + 1])]
			if a >= b {
				res += a
			} else {
				res -= a
			}
		}
		i++
	}

	return res
}

func IntToRoman(num int) string {
	m := map[int]string{
		1000: "M",
		900: "CM",
		500: "D",
		400: "CD",
		100: "C",
		90: "XC",
		50: "L",
		40: "XL",
		10: "X",
		9: "IX",
		5: "V",
		4: "IV",
		1: "I",
	}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""
	i := 0
	for i < 13 {
		if num >= nums[i]{
			num -= nums[i]
			res += m[nums[i]]
		} else {
			i++
		}
	}
	return res
}

func IntToRoman2(num int) string {
	m := map[int]string{
		1000: "M",
		900: "CM",
		500: "D",
		400: "CD",
		100: "C",
		90: "XC",
		50: "L",
		40: "XL",
		10: "X",
		9: "IX",
		5: "V",
		4: "IV",
		1: "I",
	}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""
	for _, key := range(nums) {
		if num / key != 0 {
			count := num / key
			for i := 0; i < count; i++{
				res += m[key]
			}
		}
		num = num % key
	}
	return res
}