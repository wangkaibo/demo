package leetcode

func ClimbStairs(n int) int {
	dp := map[int]int{
		1:1,
		2:2,
	}
	for i := 3; i <= n; i++ {
		dpi := dp[i - 1] + dp[i - 2]
		dp[i] = dpi
		if i > 3 {
			delete(dp, 1)
			delete(dp, i - 2)
		}
	}
	return dp[n]
}