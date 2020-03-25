package leetcode

import (
	"strconv"
)

func AddStrings(num1 string, num2 string) string {
	res := ""
	flag, i, j := 0, len(num1) - 1,len(num2) - 1
	a, b, sum := 0,0,0
	for i >= 0 || j >= 0 {
		if i >= 0 {
			a,_ = strconv.Atoi(string(num1[i]))
		} else {
			a = 0
		}
		if j >= 0 {
			b,_ = strconv.Atoi(string(num2[j]))
		} else {
			b = 0
		}
		sum = a + b + flag
		if sum > 9 {
			flag = 1
			sum = sum % 10
		} else {
			flag = 0
		}
		str := strconv.Itoa(sum)
		res = str + res
		i--
		j--
	}

	if flag == 1 {
		res = "1" + res
	}

	return res
}