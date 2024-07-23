package main

import "fmt"

// 最小路径和
func main() {
	num := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(num)
	res := minPathSum(num)
	fmt.Println(res)
}

func minPathSum(num [][]int) [][]int {
	if len(num[0]) == 0 || len(num) == 0 {
		return num
	}
	for i := 0; i < len(num[0]); i++ {
		for j := 0; j < len(num); j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				num[i][j] += num[0][j-1]
			} else if j == 0 {
				num[i][j] += num[i-1][0]
			} else {
				num[i][j] += min(num[i-1][j], num[i][j-1])
			}
		}
	}
	return num[len(num)-1][len(num[0])-1]
}
