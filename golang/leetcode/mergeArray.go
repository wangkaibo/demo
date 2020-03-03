package leetcode

func Merge(A []int, m int, B []int, n int) {
	i, pa, pb := m + n - 1, m - 1, n - 1
	for pa >= 0 || pb >= 0 {
		if pa == -1 {
			A[i] = B[pb]
			pb--
		} else if pb == -1 {
			A[i] = A[pa]
			pa--
		} else if A[pa] < B[pb] {
			A[i] = B[pb]
			pb--
		} else {
			A[i] = A[pa]
			pa--
		}
		i--
	}
}