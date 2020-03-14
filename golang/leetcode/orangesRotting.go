package leetcode

import (
	"fmt"
)

func OrangesRotting(grid [][]int) int {
	rottens := make([][]int, 0)
	freshs := make([][]int, 0)
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row;i++  {
		for j := 0;j < col;j++ {
			if grid[i][j] == 2 {
				rottens = append(rottens, []int{i,j})
			} else if grid[i][j] == 1 {
				freshs = append(freshs, []int{i,j})
			}
		}
	}
	var rotten []int
	//for len(rotten) > 0 {
	for _,rotten = range rottens {
		up := []int{rotten[0],rotten[1] - 1}
		down := []int{rotten[0],rotten[1] + 1}
		left := []int{rotten[0] - 1,rotten[1]}
		right := []int{rotten[0] + 1,rotten[1]}
		if ([]int{0,1}[0] == down[0] && down[1] == []int{0,1}[1]) {
			fmt.Println(up, down, left, right)
		}
	}

	//}
	fmt.Println(rottens)
	fmt.Println(freshs)
	return -1
}