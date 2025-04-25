//2d array

package main

import "fmt"

func main() {
	mat := make([][]int, 0)
	col := 3
	row := 3

	for i := 0; i < row; i++ {
		rows := make([]int, col)
		for j := 0; j < col; j++ {
			rows[j] = i * j
		}
		mat = append(mat, rows)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}
}
