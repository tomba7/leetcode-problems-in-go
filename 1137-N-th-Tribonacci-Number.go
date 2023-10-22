func tribonacci(n int) int {
    if n < 2 {
        return n
    } else if n == 2 {
        return 1
    }
    dp := make([]int, n+1)
    dp[1], dp[2] = 1, 1
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
    }
    return dp[n]
}
