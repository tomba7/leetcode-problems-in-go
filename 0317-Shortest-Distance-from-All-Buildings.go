func shortestDistance(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    distance := make([][]int, m)
    reach := make([][]int, m)
    for i := range distance {
        distance[i] = make([]int, n)
        reach[i] = make([]int, n)
    }
    var bldgCount int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                bldgCount++
                calculateDistance(i, j, grid, distance, reach)
            }
        }
    }
    var res int = math.MaxInt32
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 && reach[i][j] == bldgCount {
                res = min(res, distance[i][j])
            }
        }
    }
    if res == math.MaxInt32 {
        return -1
    }
    return res
}

func calculateDistance(i, j int, grid, distance, reach [][]int) {
    m, n := len(grid), len(grid[0])
    q := list.New()
    q.PushBack([]int{i, j})
    seen := make([][]bool, m)
    for i := range seen {
        seen[i] = make([]bool, n)
    }
    seen[i][j] = true
    var dist int
    levelSize := q.Len()
    for q.Len() != 0 {
        curr := q.Remove(q.Front()).([]int)
        row, col := curr[0], curr[1]
        levelSize--
        for _, offset := range offsets {
            nextRow, nextCol := row + offset[0], col + offset[1]
            if valid(nextRow, nextCol, grid) && !seen[nextRow][nextCol] {
                seen[nextRow][nextCol] = true
                distance[nextRow][nextCol] += dist + 1
                reach[nextRow][nextCol]++
                q.PushBack([]int{nextRow, nextCol})
            }
        }
        if levelSize == 0 {
            dist++
            levelSize = q.Len()
        }
    }
}

func valid(row, col int, grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    return 0 <= row && row < m && 0 <= col && col < n && grid[row][col] == 0
}

func min(x, y int) int { if x < y { return x }; return y }

var offsets = [][]int{
    {0, 1}, {0, -1}, {1, 0}, {-1, 0},
}
