func uniquePathsIII(grid [][]int) int {
    var r, c int
    k := 1
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                r, c = i, j
            } else if grid[i][j] == 0 {
                k++
            }
        }
    }
    return uniquePaths(grid, r, c, k)
}

func uniquePaths(grid [][]int, r, c, k int) int {
    m, n := len(grid), len(grid[0])
    if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] < 0 {
        return 0
    } else if grid[r][c] == 2 {
        if k == 0 { return 1 }
        return 0
    }
    grid[r][c] = -2
    total := uniquePaths(grid, r+1, c, k-1) +
             uniquePaths(grid, r-1, c, k-1) +
             uniquePaths(grid, r, c+1, k-1) +
             uniquePaths(grid, r, c-1, k-1)
    grid[r][c] = 0
    return total
}
