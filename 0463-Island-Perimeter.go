/*
 [0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]
 
 - Scan the grid row by row
 - For every grid[i][j] == 1, compute the perimeter and add to result
   * Increment the perimeter by 1 if the neigboring cell is water
   * This include (a) out of bounds cell and (b) cells whose values are '0'
 */
func islandPerimeter(grid [][]int) int {
    var res int
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                res += cellPerimeter(grid, i, j)
            }
        }
    }
    return res
}

func cellPerimeter(grid [][]int, i, j int) int {
    m, n := len(grid), len(grid[0])
    var count int
    for _, offset := range offsets {
        r, c := i + offset[0], j + offset[1]
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] == 0 { count++ }
    }
    return count
}

var offsets = [][]int{
    {0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

---

func islandPerimeter(grid [][]int) int {
    islands, neighbors := 0, 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                islands++
                if i < len(grid) - 1 && grid[i + 1][j] == 1 {neighbors++}
                if j < len(grid[0]) - 1 && grid[i][j + 1] == 1 {neighbors++}
            }
        }
    }
    return islands * 4 - neighbors * 2
}
