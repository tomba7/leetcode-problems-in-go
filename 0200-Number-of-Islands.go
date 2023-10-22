func numIslands(grid [][]byte) int {
    n, m := len(grid), len(grid[0])
    var res int
    visited := make([][]bool, n)
    for i := range visited { visited[i] = make([]bool, m) }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            numCells := dfs(grid, i, j, visited)
            if numCells != 0 { res++ }
        }
    }
    return res
}

func dfs(grid [][]byte, i, j int, visited [][]bool) int {
    n, m := len(grid), len(grid[0])
    if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == '0' || visited[i][j] { return 0 }
    visited[i][j] = true
    return dfs(grid, i+1, j, visited) + dfs(grid, i-1, j, visited) +
           dfs(grid, i, j+1, visited) + dfs(grid, i, j-1, visited) + 1
}

---

/* 3:44
 - Traverse the grid row by row and perform DFS
 - If a land has been visited skip starting DFS
 - Count the number of DFS
 - This is our result
   ["1","1","0","0","0"],
   ["1","1","0","0","0"],
   ["0","0","1","0","0"],
   ["0","0","0","1","1"]
   
   ["-1","-1", "0", "0", "0"],
   ["-1","-1", "0", "0", "0"],
   [ "0", "0","-1", "0", "0"],
   [ "0", "0", "0","-1","-1"]
 */
func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    var res int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != '1' { continue }
            dfs(grid, i, j)
            res++
        }
    }
    return res
}

func dfs(grid [][]byte, i, j int) {
    grid[i][j] = '2'
    for _, offset := range offsets {
        r, c := i + offset[0], j + offset[1]
        if !valid(grid, r, c) || grid[r][c] != '1' { continue }
        dfs(grid, r, c)
    }
}

func valid(grid [][]byte, i, j int) bool {
    return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
}

var offsets = [][]int{
    {1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

---

// If you are allowed to modify the input grid
func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {return 0}
    islandCount := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                visitIsland(grid, i, j)
                islandCount++
            }
        }
    }
    return islandCount
}

func visitIsland(grid [][]byte, i, j int) {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {return}
    grid[i][j] = '0'
    visitIsland(grid, i + 1, j)
    visitIsland(grid, i - 1, j)
    visitIsland(grid, i, j + 1)
    visitIsland(grid, i, j - 1)
}

---

func numIslands(grid [][]byte) int {
    n, m := len(grid), len(grid[0])
    var res int
    visited := make(map[int]bool)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if valid(grid, i, j, visited) {
                dfs(grid, i, j, visited)
                res++
            }
        }
    }
    return res
}

func dfs(grid [][]byte, i, j int, visited map[int]bool) {
    visited[index(i, j, len(grid[0]))] = true
    if valid(grid, i+1, j, visited) { dfs(grid, i+1, j, visited) }
    if valid(grid, i-1, j, visited) { dfs(grid, i-1, j, visited) }
    if valid(grid, i, j+1, visited) { dfs(grid, i, j+1, visited) }
    if valid(grid, i, j-1, visited) { dfs(grid, i, j-1, visited) }
}

func valid(grid [][]byte, i, j int, visited map[int]bool) bool {
    n, m := len(grid), len(grid[0])
    if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == '0' || visited[index(i, j, m)] {
        return false
    }
    return true
}

func index(i, j, numCols int) int {
    return i * numCols + j
}

---

func numIslands(grid [][]byte) int {
    n, m := len(grid), len(grid[0])
    var res int
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, m)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if valid(grid, i, j, visited) {
                dfs(grid, i, j, visited)
                res++
            }
        }
    }
    return res
}

func dfs(grid [][]byte, i, j int, visited [][]bool) {
    visited[i][j] = true
    if valid(grid, i+1, j, visited) { dfs(grid, i+1, j, visited) }
    if valid(grid, i-1, j, visited) { dfs(grid, i-1, j, visited) }
    if valid(grid, i, j+1, visited) { dfs(grid, i, j+1, visited) }
    if valid(grid, i, j-1, visited) { dfs(grid, i, j-1, visited) }
}

func valid(grid [][]byte, i, j int, visited [][]bool) bool {
    n, m := len(grid), len(grid[0])
    if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == '0' || visited[i][j] {
        return false
    }
    return true
}
