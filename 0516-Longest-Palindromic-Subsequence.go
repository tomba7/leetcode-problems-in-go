// DP
func longestPalindromeSubseq(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        dp[i][i] = 1
    }
    for i := n-2; i >= 0; i-- {
        for j := i+1; j < n; j++ {
            if s[i] == s[j] {
                dp[i][j] = 2 + dp[i+1][j-1]
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
            }
        }
    }
    return dp[0][n-1]
}

func max(x, y int) int { if x > y { return x }; return y }

---

// Memoization
func longestPalindromeSubseq(s string) int {
    n := len(s)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    return helper(s, 0, len(s)-1, memo)
}

func helper(s string, i, j int, memo [][]int) int {
    if i == j {
        return 1
    } else if i > j {
        return 0
    }
    if memo[i][j] != -1 {
        return memo[i][j]
    }
    if s[i] == s[j] {
        memo[i][j] = 2 + helper(s, i+1, j-1, memo)
    } else {
        memo[i][j] = max(helper(s, i+1, j, memo), helper(s, i, j-1, memo))
    }
    return memo[i][j]
}

func max(x, y int) int { if x > y { return x }; return y }

---

// Recursion (TLE)
func longestPalindromeSubseq(s string) int {
    return helper(s, 0, len(s)-1)
}

func helper(s string, lo, hi int) int {
    if lo == hi {
        return 1
    } else if lo > hi {
        return 0
    }
    if s[lo] == s[hi] {
        return 2 + helper(s, lo+1, hi-1)
    } else {
        return max(helper(s, lo+1, hi), helper(s, lo, hi-1))
    }
}

func max(x, y int) int { if x > y { return x }; return y }
