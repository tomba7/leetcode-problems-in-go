// TLE (Recursive DFS Solution)
func longestIncreasingPath(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])
    var res int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res = max(res, dfs(matrix, i, j)) 
        }
    }
    return res
}

func dfs(matrix [][]int, i, j int) int {
    m, n := len(matrix), len(matrix[0])
    var res int
    for _, offset := range offsets {
        x, y := i + offset[0], j + offset[1]
        if x < 0 || x >= m || y < 0 || y >= n || matrix[x][y] <= matrix[i][j] {
            continue
        }
        res = max(res, dfs(matrix, x, y))
    }
    return res + 1 
}

func max(x, y int) int { if x > y { return x }; return y }
var offsets = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
