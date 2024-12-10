package day_10

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/utils"
)

func (d Day) firstPart() any {

	grid := d.readGrid()

	n, m := utils.GridSize(grid)

	var dfs func(i, j int, reached map[int]bool)

	dfs = func(i, j int, reached map[int]bool) {
		if grid[i][j] == 9 {
			reached[i*m+j] = true
		}

		for k := 0; k < 4; k++ {
			if utils.IsOutOfBounds(i+dy[k], j+dx[k], n, m) || grid[i+dy[k]][j+dx[k]] != grid[i][j]+1 {
				continue
			}
			dfs(i+dy[k], j+dx[k], reached)
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				reached := make(map[int]bool)
				dfs(i, j, reached)
				res += len(reached)
			}
		}
	}

	return res
}
