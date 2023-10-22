func countSubstrings(s string) int {
    var res int
    for i := 0; i < len(s); i++ {
        countSubstringsHelper(s, i, i, &res)    // odd length
        countSubstringsHelper(s, i, i+1, &res)  // even length
    }
    return res
}

func countSubstringsHelper(s string, i, j int, res *int) {
    for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
        *res++  // expand outwards from the center and count
    }
}

---

func countSubstrings(s string) int {
    n := len(s)
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }
    for l := 1; l <= n; l++ {
        for i := 0; i < n - l + 1; i++ {
            j := i + l - 1
            if l == 1 {
                dp[i][j] = true
            } else if l == 2 && s[i] == s[j] {
                dp[i][j] = true
            } else if s[i] == s[j] && dp[i+1][j-1] {
                dp[i][j] = true
            }
        }
    }
    var res int
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            if dp[i][j] {
                res++
            }
        }
    }
    return res
}
