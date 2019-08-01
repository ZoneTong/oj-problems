// 数独
// 数独是一个我们都非常熟悉的经典游戏，运用计算机我们可以很快地解开数独难题，现在有一些简单的数独题目，请编写一个程序求解。

// 输入描述:
// 输入9行，每行为空格隔开的9个数字，为0的地方就是需要填充的。

// 输出描述:
// 输出九行，每行九个空格隔开的数字，为解出的答案。
package main

import "fmt"

func main() {

LOOP:
	for {
		var m [9][9]int
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				n, _ := fmt.Scan(&m[i][j])
				if n == 0 {
					break LOOP
				}
			}
		}

		m, _ = fillnext(0, m)
		// fmt.Println("zht", b)
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Print(m[i][j])
				if j != 8 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}

		// fmt.Println(validWhole(m))
	}
}

func fillnext(n int, m [9][9]int) ([9][9]int, bool) {
	for n < 81 && m[n/9][n%9] != 0 {
		n++
	}
	if n == 81 {
		return m, true
	}

	i, j := n/9, n%9

	for k := 1; k <= 9; k++ {
		// fmt.Println("zht", i, j, k, m)
		m[i][j] = k
		if validInput(i, j, m) {
			m, ok := fillnext(n+1, m)
			if ok {
				return m, true
			}
		}
		m[i][j] = 0
	}
	return m, false
}

func validInput(row, col int, m [9][9]int) bool {
	var row09 [10]int
	for i := 0; i < 9; i++ {
		num := m[row][i]
		if num == 0 {
			continue
		}
		if row09[num] == 1 {
			return false
		}
		row09[num] = 1
	}

	var col09 [10]int
	for i := 0; i < 9; i++ {
		num := m[i][col]
		if num == 0 {
			continue
		}
		if col09[num] == 1 {
			return false
		}
		col09[num] = 1
	}

	var grid [10]int
	var rowbase, colbase = row / 3, col / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := m[i+rowbase*3][j+colbase*3]
			if num == 0 {
				continue
			}
			if grid[num] == 1 {
				return false
			}
			grid[num] = 1
		}
	}

	return true
}

func validWhole(m [9][9]int) bool {
	var n int
	for n < 81 {
		if !validInput(n/9, n%9, m) {
			return false
		}
		n++
	}
	return true
}
