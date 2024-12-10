package day_10

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/utils"
)

func (d Day) secondPart() any {

	grid := d.readGrid()

	n, m := utils.GridSize(grid)

	memo := utils.MakeGrid[int](n, m, -1)

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if grid[i][j] == 9 {
			return 1
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		res := 0
		for k := 0; k < 4; k++ {
			if utils.IsOutOfBounds(i+dy[k], j+dx[k], n, m) || grid[i+dy[k]][j+dx[k]] != grid[i][j]+1 {
				continue
			}
			res += dfs(i+dy[k], j+dx[k])
		}

		memo[i][j] = res
		return res
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res += dfs(i, j)
			}
		}
	}

	return res
}
