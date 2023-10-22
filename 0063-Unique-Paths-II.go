func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    var m, n int
    if m = len(obstacleGrid); m == 0 { return 0 }
    if n = len(obstacleGrid[0]); n == 0 { return 0 }

    cache := make([][]int, m + 1)
    for i := 0; i <= m; i++ { cache[i] = make([]int, n + 1) }

    // This could be either 'cache[m - 1][n] = 1` or `cache[m][n - 1] = 1`
    // The bottom line is to make the destination cell `cache[m - 1][n - 1]` equal to 1 
    cache[m][n - 1] = 1
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if obstacleGrid[i][j] != 1 {
                cache[i][j] = cache[i + 1][j] + cache[i][j + 1]
            }
            // If there is an obstacle (i.e == 1), the cell is invalid. Therefore, the number
            // of ways to reach the destination cell from here should be 0. Since
            // it was already initialized to 0, no further action is required.
        }
    }
    return cache[0][0]
}

---

/*
    0 1 2 3
 0 [0,0,0,0],
 1 [0,0,0,0],
 2 [0,0,0,0]
 3 [0,0,1,0]
    
   f(i, j) = f(i, j+1) + f(i+1, j)
   f(i, j) = 1 if i == m-1 && j == n-1
 */

func uniquePathsWithObstacles(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    dp[m][n-1] = 1
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if grid[i][j] == 1 {
                dp[i][j] = 0
            } else {
                dp[i][j] = dp[i][j+1] + dp[i+1][j]
            }
        }
    }
    return dp[0][0]
}
