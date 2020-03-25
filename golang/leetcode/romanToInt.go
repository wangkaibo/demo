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