package _00numberofislands

func numIslands(grid [][]byte) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				num++
				dfs(grid, i, j)
			}
		}
	}
	return num
}

func dfs(grid [][]byte, i int, j int) {
	grid[i][j] = 0
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for d := 0; d < len(dir); d++ {
		x := i + dir[d][0]
		y := j + dir[d][1]
		if isValid(grid, x, y) && grid[x][y] == 1 {
			dfs(grid, x, y)
		}
	}
}

func isValid(grid [][]byte, x int, y int) bool {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return false
	}
	return true
}
