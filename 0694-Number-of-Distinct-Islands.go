func numDistinctIslands(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    seen := make([][]bool, m)
    for i := range seen {
        seen[i] = make([]bool, n)
    }
    islands := map[string]bool{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            paths := []byte{}
            dfs(i, j, i, j, grid, seen, &paths)
            if len(paths) != 0 {
                if islands[string(paths)] { continue }
                islands[string(paths)] = true
            }
        }
    }
    return len(islands)
}

func dfs(i, j, origRow, origCol int, grid [][]int, seen [][]bool, paths *[]byte) {
    m, n := len(grid), len(grid[0])
    if i < 0 || i >= m || j < 0 || j >= n || seen[i][j] || grid[i][j] == 0 { return }
    seen[i][j] = true
    *paths = append(*paths, fmt.Sprintf("%d,%d,", i - origRow, j - origCol)...)
    dfs(i+1, j, origRow, origCol, grid, seen, paths)
    dfs(i, j+1, origRow, origCol, grid, seen, paths)
    dfs(i-1, j, origRow, origCol, grid, seen, paths)
    dfs(i, j-1, origRow, origCol, grid, seen, paths)
}

---

func numDistinctIslands(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    seen := make([][]bool, m)
    for i := range seen {
        seen[i] = make([]bool, n)
    }
    islands := map[string]bool{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            paths := []byte{}
            dfs(i, j, 'S', grid, seen, &paths)
            if len(paths) != 0 {
                if islands[string(paths)] { continue }
                islands[string(paths)] = true
            }
        }
    }
    return len(islands)
}

func dfs(i, j int, ch byte, grid [][]int, seen [][]bool, paths *[]byte) {
    m, n := len(grid), len(grid[0])
    if i < 0 || i >= m || j < 0 || j >= n || seen[i][j] || grid[i][j] == 0 { return }
    seen[i][j] = true
    *paths = append(*paths, ch)
    dfs(i+1, j, 'U', grid, seen, paths)
    dfs(i, j+1, 'R', grid, seen, paths)
    dfs(i-1, j, 'D', grid, seen, paths)
    dfs(i, j-1, 'L', grid, seen, paths)
    *paths = append(*paths, 'E')
}
