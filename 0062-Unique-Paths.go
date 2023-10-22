// top-down DP solution

func uniquePaths(m int, n int) int {
    if m < 0 || n < 0 { return 0 }
    cache := make([][]int, m + 1)
    for i := 0; i <= m; i++ {
        cache[i] = make([]int, n + 1)
        for j := 0; j <= n; j++ { cache[i][j] = -1 }
    }
    return uniquePathsHelper(0, 0, cache)
}

func uniquePathsHelper(i, j int, cache [][]int) int {
    m, n := len(cache) - 1, len(cache[0]) - 1
    if i == m - 1 && j == n - 1 {
        return 1
    } else if i >= m || j >= n {
        return 0
    }
    if cache[i + 1][j] == -1 {
        cache[i + 1][j] = uniquePathsHelper(i + 1, j, cache)
    }
    if cache[i][j + 1] == -1 {
        cache[i][j + 1] = uniquePathsHelper(i, j + 1, cache)
    }
    return cache[i + 1][j] + cache[i][j + 1]
}
