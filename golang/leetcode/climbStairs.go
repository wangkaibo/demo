package leetcode

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a,b := 1, 2
	for i := 3; i < n; i++ {
		a, b = b, a + b
	}

	return a + b
}