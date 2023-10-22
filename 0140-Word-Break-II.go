/*  s = "catsanddog"
    wordDict = ["cat","cats","and","sand","dog"]
          i
    catsanddog
       j
*/
func wordBreak(s string, wordDict []string) []string {
    words := make(map[string]bool)
    for _, word := range wordDict {
        words[word] = true
    }
    n := len(s)
    dp := make([][]string, n)
    for i := 0; i < n; i++ {
        if words[s[:i+1]] {
            dp[i] = append(dp[i], s[:i+1])
        }
        for j := 0; j < i; j++ {
            if len(dp[j]) == 0 { continue }
            if words[s[j+1:i+1]] {
                for k := 0; k < len(dp[j]); k++ {
                    dp[i] = append(dp[i], dp[j][k] + " " + s[j+1:i+1])
                }
            }
        }
    }
    return dp[n-1]
}

---

/*  s = "catsanddog"
    wordDict = ["cat","cats","and","sand","dog"]
    i
    catsanddog
             j
    - starting at each index, form all possible words and check
      if they exist in the dict
    - if yes, append the word in the temp and recurse the remaining part
      of the string
    - if all chars are consumed add them to the result set
*/
func wordBreak(s string, wordDict []string) []string {
    words := make(map[string]bool)
    for _, word := range wordDict {
        words[word] = true
    }
    var res []string
    wordBreakHelper(s, 0, "", words, &res)
    return res
}

func wordBreakHelper(s string, i int, path string, words map[string]bool, res *[]string) {
    if i == len(s) {
        *res = append(*res, path)
        return
    }
    for j := i; j < len(s); j++ {
        if !words[s[i:j+1]] { continue }
        if i == 0 {
            wordBreakHelper(s, j+1, s[i:j+1], words, res)
        } else {
            wordBreakHelper(s, j+1, path + " " + s[i:j+1], words, res)
        }
    }
}

---

func wordBreak(s string, wordDict []string) []string {
    dict := map[string]bool{}
    for _, word := range wordDict {
        dict[word] = true
    }
    n := len(s)
    dp := make([][]string, n)
    for i := 0; i < n; i++ {
        if dict[s[:i+1]] {
            dp[i] = append(dp[i], s[:i+1])
        }
        for j := 0; j < i; j++ {
            if len(dp[j]) == 0 || !dict[s[j+1:i+1]] {
                continue
            }
            for k := 0; k < len(dp[j]); k++ {
                dp[i] = append(dp[i], dp[j][k] + " " + s[j+1:i+1])
            }
        }
    }
    return dp[n-1]
}
