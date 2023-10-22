func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    var m, n int
    if m = len(obstacleGrid); m == 0 { return 0 }
    if n = len(obstacleGrid[0]); n == 0 { return 0 }
    cache := make([]int, n)
    cache[0] = 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                cache[j] = 0
            } else if j > 0 {
                cache[j] += cache[j-1]
            }
        }
    }
    return cache[n-1]
}
