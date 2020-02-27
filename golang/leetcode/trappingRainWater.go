package leetcode

func Trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	// 找到最大值
 	maxIndex := 0
	for i := 0; i < len(height); i++ {
		if height[maxIndex] < height[i] {
			maxIndex = i
		}
	}
	var leftWater int
	var rightWater int
	leftMax := height[0]
	// 从左往右爬楼，爬到最高
	for j := 0; j < maxIndex; j++ {
		if height[j] < leftMax {
			leftWater += leftMax - height[j]
		} else {
			leftMax = height[j]
		}
	}
	// 从右往左爬楼，爬到最高
	rightMax := height[len(height) - 1]
	for k := len(height) - 1; k > maxIndex; k-- {
		if height[k] < rightMax {
			rightWater += rightMax - height[k]
		} else {
			rightMax = height[k]
		}
	}
	return leftWater + rightWater
}