package stack

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func numIslands(grid [][]byte) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		for c := 0; c < len(grid[i]); c++ {
			if grid[i][c] == '1' {
				num++
				visit(grid, i, c)
			}
		}
	}
	return num
}

func visit(grid [][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) || grid[row][col] != '1' {
		return
	}
	grid[row][col] = '*'
	for i := 0; i < len(dirs); i++ {
		visit(grid, row+dirs[i][0], col+dirs[i][1])
	}
}
