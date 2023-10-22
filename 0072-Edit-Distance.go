func minDistance(w1, w2 string) int {
    n, m := len(w1), len(w2)
    dp := newDPTable(n, m)
    for i := n-1; i >= 0; i-- {
        for j := m-1; j >= 0; j-- {
            if w1[i] == w2[j] {
                dp[i][j] = dp[i+1][j+1]
            } else {
                dp[i][j] = 1 + min3(dp[i][j+1], dp[i+1][j], dp[i+1][j+1])
            }
        }
    }
    return dp[0][0]
}

func newDPTable(n, m int) [][]int {
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
        dp[i][m] = n-i
    }
    for j := m-1; j >= 0; j-- {
        dp[n][j] = m-j
    }
    return dp
}

func min(x, y int) int { if x < y { return x }; return y }
func min3(x, y, z int) int { return min(x, min(y, z)) }

---

func minDistance(s, t string) int {
    m, n := len(s), len(t)
    if m == 0 {return n}
    if n == 0 {return m}
    memo := make([][]int, m + 1)
    for i := 0; i < len(memo); i++ {
        memo[i] = make([]int, n + 1)
    }
    for i := 0; i <= m; i++ {
        memo[i][0] = i
    }
    for j := 0; j <= n; j++ {
        memo[0][j] = j
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s[i - 1] == t[j - 1] {
                memo[i][j] = memo[i - 1][j - 1]
            } else {
                memo[i][j] = min(
                    memo[i - 1][j - 1],
                    min(memo[i - 1][j], memo[i][j - 1]),
                ) + 1
            }
        }
    }
    return memo[m][n]
}

func min(x, y int) int {if x < y {return x}; return y}

---

/*
    Recurrence relation

    func minDistance(word1 string, word2 string) int {
        return helper(word1, word2, 0, 0)
    }
    
    func helper(w1, w2 string, i, j int) int {
        if i == len(w1) { return len(w2) - j } 
        if j == len(w2) { return len(w1) - i }
        if w1[i] == w2[j] {
            return helper(w1, w2, i+1, j+1)
        }
        return 1 + min3(
            helper(w1, w2, i, j+1),
            helper(w1, w2, i+1, j),
            helper(w1, w2, i+1, j+1),
        )
    }
    
    func min(x, y int) int { if x < y { return x }; return y }
    func min3(x, y, z int) int { return min(min(x, y), z) }
*/
