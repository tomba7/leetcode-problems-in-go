func palindromePairs(words []string) [][]int {
    wordSet := map[string]int{}
    for i, word := range words {
        wordSet[word] = i
    }
    res := [][]int{}
    for word, i := range wordSet {
        reversed := reverse(word)
        if _, contains := wordSet[reversed]; contains && wordSet[reversed] != i {
            res = append(res, []int{i, wordSet[reversed]})
        }

        for _, suffix := range validSuffixes(word) {
            reversedSuffix := reverse(suffix)
            if _, contains := wordSet[reversedSuffix]; contains {
                res = append(res, []int{wordSet[reversedSuffix], i})
            }
        }

        for _, prefix := range validPrefixes(word) {
            reversedPrefix := reverse(prefix)
            if _, contains := wordSet[reversedPrefix]; contains {
                res = append(res, []int{i, wordSet[reversedPrefix]})
            }
        }
    }
    return res
}

func validPrefixes(s string) []string {
    var res []string
    for i := 0; i < len(s); i++ {
        if palindrome(s, i, len(s)-1) {
            res = append(res, s[:i])
        }
    }
    return res
}

func validSuffixes(s string) []string {
    var res []string
    for i := 0; i < len(s); i++ {
        if palindrome(s, 0, i) {
            res = append(res, s[i+1:])
        }
    }
    return res
}

func reverse(s string) string {
    cp := []byte(s)
    for i, j := 0, len(cp)-1; i < j; i, j = i+1, j-1 {
        cp[i], cp[j] = cp[j], cp[i]
    }
    return string(cp)
}

func palindrome(s string, i, j int) bool {
    for ; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] { return false }
    }
    return true
}
