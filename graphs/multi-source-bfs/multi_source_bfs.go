package multisourcebfs

import "container/list"

// grid -| adjacency matrix where grid[i][j] has only three values
// 0 - space
// 1 - fresh fruit
// 2 - rotten fruit
// Every minute, if a fresh fruit is horizontally or vertically adjacent to a rotten fruit, then the fresh fruit also becomes rotten.
// Return the minimum number of minutes that must elapse until there are zero fresh fruits remaining. If this state is impossible within the grid, return -1.
// references: https://leetcode.com/problems/rotting-oranges/description/

func MultiSourceBFS(grid [][]int) (minimumMinutes int) {
	rows, cols := len(grid), len(grid[0])
	q := list.New()

	freshFruits := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] == 1 {
				freshFruits++
			}
			if grid[r][c] == 2 {
				q.PushBack([2]int{r, c})
			}
		}
	}

	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for q.Len() > 0 && freshFruits > 0 {
		qLen := q.Len()
		for range qLen {
			node := q.Remove(q.Front()).([2]int)
			row, col := node[0], node[1]

			for _, dir := range dirs {
				r := row + dir[0]
				c := col + dir[1]
				if r < 0 || r >= rows ||
					c < 0 || c >= cols || grid[r][c] == 0 {
					continue
				}
				if grid[r][c] == 1 {
					q.PushBack([2]int{r, c})
					freshFruits--
				}
				grid[r][c] = 2
			}
		}
		minimumMinutes++
	}

	if freshFruits == 0 {
		return
	}

	return -1
}
