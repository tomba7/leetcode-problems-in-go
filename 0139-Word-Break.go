func wordBreak(s string, wordDict []string) bool {
    dict := map[string]bool{}
    for _, word := range wordDict {
        dict[word] = true
    }
    n := len(s)
    dp := make([]bool, n)
    for i := 0; i < n; i++ {
        if dict[s[:i+1]] {
            dp[i] = true
            continue
        }
        for j := 0; j < n; j++ {
            if dp[j] && dict[s[j+1:i+1]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[n-1]
}
