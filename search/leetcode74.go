package search

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := findVertical(matrix, len(matrix[0])-1, target)
	if row == -1 {
		return false
	}

	return findHorizontal(matrix, row, target) != -1
}

func findVertical(matrix [][]int, col, target int) int {
	start, end := 0, len(matrix)
	for mid := (start + end) / 2; start < end; mid = (start + end) / 2 {
		if matrix[mid][col] == target {
			return mid
		}
		if matrix[mid][col] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if start == len(matrix) || (matrix[start][col] < target && start == len(matrix)-1) {
		return -1
	}
	return start
}

func findHorizontal(matrix [][]int, row, target int) int {
	start, end := 0, len(matrix[row])
	for mid := (start + end) / 2; start < end; mid = (start + end) / 2 {
		if matrix[row][mid] == target {
			return mid
		}
		if matrix[row][mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}
