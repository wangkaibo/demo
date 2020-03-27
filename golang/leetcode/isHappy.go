package leetcode

func IsHappy(n int) bool {
	cal := func(i int) int {
		res := 0
		for i > 0 {
			res += (i % 10) * (i % 10)
			i = i / 10
		}
		return res
	}

	slow, fast := n, n
	slow = cal(slow)
	fast = cal(fast)
	fast = cal(fast)
	for slow != fast {
		slow = cal(slow)
		fast = cal(fast)
		fast = cal(fast)
	}
	return slow == 1
}

func IsValid(s string) bool {
	m := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	queue := []string{}
	for i := 0; i < len(s); i++ {
		v, exist := m[string(s[i])]
		if exist {
			if len(queue) > 0 {
				pop := queue[len(queue)-1]
				queue = queue[0 : len(queue)-1]
				if pop != v {
					return false
				}
			} else {
				return false
			}
		} else {
			queue = append(queue, string(s[i]))
		}
	}
	return len(queue) == 0
}

func IsAnagram(s string, t string) bool {
	list := []int{}
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {

	}
	for _,v := range(list) {
		if v != 0 {
			return false
		}
	}
	return true
}
