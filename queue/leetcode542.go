package queue

type Position struct {
	x int
	y int
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func updateMatrix(matrix [][]int) [][]int {
	level := make([]*Position, 0)
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[i]))
		for k := 0; k < len(matrix[i]); k++ {
			if matrix[i][k] == 0 {
				level = append(level, &Position{x: i, y: k})
				visited[i][k] = true
			}
		}
	}
	distance := 1
	for len(level) > 0 {
		nextLevel := make([]*Position, 0)
		for len(level) > 0 {
			pop := level[len(level)-1]
			level = level[:len(level)-1]
			for i := 0; i < len(directions); i++ {
				x := pop.x + directions[i][0]
				y := pop.y + directions[i][1]
				if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[x]) || visited[x][y] {
					continue
				}
				matrix[x][y] = distance
				visited[x][y] = true
				nextLevel = append(nextLevel, &Position{x: x, y: y})
			}
		}
		level = nextLevel
		distance++
	}
	return matrix
}
