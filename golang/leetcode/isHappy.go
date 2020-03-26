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