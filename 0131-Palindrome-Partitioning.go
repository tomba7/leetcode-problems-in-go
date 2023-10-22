func partition(s string) [][]string {
    var res [][]string
    p := make([]string, len(s))
    partitionHelper(s, 0, p, 0, &res)
    return res
}

func partitionHelper(s string, i int, p []string, k int, res *[][]string) {
    if i == len(s) {
        *res = append(*res, append([]string{}, p[:k]...))
        return
    }
    for j := i; j < len(s); j++ {
        if !palindrome(s, i, j) { continue }
        p[k] = s[i:j+1]
        partitionHelper(s, j+1, p, k+1, res)
    }
}

func palindrome(s string, i, j int) bool {
    for ; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] { return false }
    }
    return true
}

---

func partition(s string) [][]string {
    valid := createPalindromeTable(s)
    res := make([][][]string, len(s))
    for i := 0; i < len(s); i++ {
        if valid[0][i] {
            res[i] = [][]string{{s[:i+1]}}
        }
        for j := 0; j < i; j++ {
            var tmp [][]string
            if !valid[j+1][i] { continue }
            for k := 0; k < len(res[j]); k++ {
                clone := append([]string{}, res[j][k]...)
                tmp = append(tmp, append(clone, s[j+1:i+1]))
            }
            res[i] = append(res[i], tmp...)
        }
    }
    return res[len(s)-1]
}

func createPalindromeTable(s string) [][]bool {
	palindrome := make([][]bool, len(s))
	for i := range palindrome {
		palindrome[i] = make([]bool, len(s))
	}
	for length := 1; length <= len(s); length++ {
		for start := 0; start < len(s); start++ {
			end := start+length-1
			if end >= len(s) { break }
			if length <= 2 {
				palindrome[start][end] = s[start] == s[end]
			} else {
				palindrome[start][end] = s[start] == s[end] && palindrome[start+1][end-1]
			}
		}
	}
	return palindrome
}

---

func partition(s string) [][]string {
    dp := make([][][]string, len(s))
    for l := 0; l < len(s); l++ {
        for i := 0; i < l; i++ {
            if !palindrome(s[i+1:l+1]) { continue }
            for j := 0; j < len(dp[i]); j++ {
                clone := append([]string{}, dp[i][j]...)
                dp[l] = append(dp[l], append(clone, s[i+1:l+1]))
            }
        }
        if palindrome(s[:l+1]) {
            dp[l] = append(dp[l], append([]string{}, s[:l+1]))
        }
    }
    return dp[len(s)-1]
}

func palindrome(s string) bool {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] { return false }
    }
    return true
}
