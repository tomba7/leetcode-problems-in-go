func minPathSum(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    dp := newTable(n, m)
    dp[n][m-1], dp[n-1][m] = 0, 0
    for i := n-1; i >= 0; i-- {
        for j := m-1; j >= 0; j-- {
            dp[i][j] = grid[i][j] + min(dp[i+1][j], dp[i][j+1])
        }
    }
    return dp[0][0]
}

func newTable(n, m int) [][]int {
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }
    for i := 0; i <= n; i++ {
        dp[i][m] = math.MaxInt32
    }
    for j := 0; j <= m; j++ {
        dp[n][j] = math.MaxInt32
    }
    return dp
}

func min(x, y int) int { if x < y { return x }; return y }

---

// If you're allowed to modify the input grid
func minPathSum(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    for i := n-2; i >= 0; i-- {
        grid[i][m-1] += grid[i+1][m-1]
    }
    for j := m-2; j >= 0; j-- {
        grid[n-1][j] += grid[n-1][j+1]
    }
    for i := n-2; i >= 0; i-- {
        for j := m-2; j >= 0; j-- {
            grid[i][j] += min(grid[i+1][j], grid[i][j+1])
        }
    }
    return grid[0][0]
}

func min(x, y int) int { if x < y { return x }; return y }

---

// Memoization
func minPathSum(grid [][]int) int {
    cache := make([][]int, len(grid))
    for i := range cache {
        cache[i] = make([]int, len(grid[0]))
        for j := range cache[i] {
            cache[i][j] = -1
        }
    }
    return helper(grid, 0, 0, cache)
}

func helper(grid [][]int, r, c int, cache [][]int) int {
    n, m := len(grid), len(grid[0])
    if r == n || c == m {
        return math.MaxInt32
    } else if r == n-1 && c == m-1 {
        return grid[r][c]
    }
    if cache[r][c] == -1 {
        cache[r][c] = min(
            helper(grid, r, c+1, cache),
            helper(grid, r+1, c, cache),
        ) + grid[r][c]
    }
    return cache[r][c]
}

func min(x, y int) int { if x < y { return x }; return y }
