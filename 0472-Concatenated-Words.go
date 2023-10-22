func findAllConcatenatedWordsInADict(words []string) []string {
    var res []string
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    dict := map[string]bool{}
    for _, word := range words {
        if word != "" && wordBreak(word, dict) {
            res = append(res, word)
        }
        dict[word] = true
    }
    return res
}

func wordBreak(s string, dict map[string]bool) bool {
    n := len(s)
    dp := make([]bool, n)
    for i := 0; i < n; i++ {
        if dict[s[:i+1]] {
            dp[i] = true
        }
        for j := 0; j < i; j++ {
            if !dp[j] { continue }
            if dict[s[j+1:i+1]] {
                dp[i] = true
            }
        }
    }
    return dp[n-1]
}
