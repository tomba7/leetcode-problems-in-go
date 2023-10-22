func numSquares(n int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = 10000
    }
    for i := 1; i <= n; i++ {
        for j := 1; j*j <= i; j++ {
            dp[i] = min(dp[i], dp[i-j*j])
        }
        dp[i]++
    }
    return dp[n]
}

func min(x, y int) int { if x < y { return x }; return y }

---

func numSquares(n int) int {
    cache := make([]int, n+1)
    return numSquaresHelper(n, cache)
}

func numSquaresHelper(n int, cache []int) int {
    if n < 0 {
        return 10000
    } else if n == 0 {
        return 0
    }
    if cache[n] != 0 {
        return cache[n]
    }
    var res = 10000
    for i := 1; i*i <= n; i++ {
        res = min(res, numSquaresHelper(n - i*i, cache))
    }
    if res != 1000 {
        res += 1
    }
    cache[n] = res
    return res
}

func min(x, y int) int { if x < y { return x }; return y }
