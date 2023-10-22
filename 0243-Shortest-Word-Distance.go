func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    n := len(wordsDict)
    res := n
    i1, i2 := -1, -1
    for i, word := range wordsDict {
        if word == word1 { i1 = i }
        if word == word2 { i2 = i}
        if i1 != -1 && i2 != -1 {
            res = min(res, abs(i1 - i2))
        }
    }
    return res
}

func min(x, y int) int { if x < y { return x }; return y }
func abs(x int) int { if x < 0 { return -x }; return x }

---

// Brute force
func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    n := len(wordsDict)
    res := n
    for i := 0; i < n; i++ {
        if wordsDict[i] != word1 { continue }
        for j := 0; j < n; j++ {
            if wordsDict[j] != word2 { continue }
            res = min(res, abs(i - j))
        }
    }
    return res
}

func min(x, y int) int { if x < y { return x }; return y }
func abs(x int) int { if x < 0 { return -x }; return x }
