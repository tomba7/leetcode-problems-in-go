func uniquePaths(m, n int) int {
    cache := make([]int, n)
    cache[0] = 1
    for i := 0; i < m; i++ {
        for j := 1; j < n; j++ {
            cache[j] += cache[j - 1]
        }
    }
    return cache[n - 1]
}
