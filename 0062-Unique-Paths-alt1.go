// bottom-up DP solution

func uniquePaths(m, n int) int {
    cache := make([][]int, m + 1)
    for i := range cache { cache[i] = make([]int, n + 1) }
    cache[1][0] = 1
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            cache[i][j] = cache[i][j - 1] + cache[i - 1][j]
        }
    }
    return cache[m][n]
}
