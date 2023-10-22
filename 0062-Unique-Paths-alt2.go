func uniquePaths(m, n int) int {
    if m < 1 || n < 1 { return 0 }
    table := make([]int, n)
    for j := 0; j < n; j++ { table[j] = 1 }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            table[j] += table[j - 1]
        }
    }
    return table[n - 1]
}
