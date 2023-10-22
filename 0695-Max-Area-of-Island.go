/*
 - Starting at every cell, perform DFS and maintain the maxAreaSoFar
 - Mark a cell visited by setting -1
 */

func maxAreaOfIsland(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    var res int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res = max(res, dfs(grid, i, j))
        }
    }
    return res
}

func dfs(grid [][]int, i, j int) int {
    m, n := len(grid), len(grid[0])
    if i < 0 || i == m || j < 0 || j == n || grid[i][j] != 1 { return 0 }
    grid[i][j] = -1
    var size int
    for _, offset := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
        r, c := i + offset[0], j + offset[1]
        size += dfs(grid, r, c)
    }
    return size + 1
}

func max(x, y int) int { if x > y { return x }; return y }
