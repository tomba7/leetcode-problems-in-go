/*
    index['h'] = 0
    index['l'] = 1
    
 - compare every word to the next starting at i = 0 
 - for each pair of words, compare each char from left to right
 - if the chars are the same, continue
 - if not, return index[ch1] < index[ch2]
 - edge case: if i < m && j == n, return false

 */
func isAlienSorted(words []string, order string) bool {
    index := newCharIndex(order)
    for k := 0; k < len(words)-1; k++ {
        word1, word2 := words[k], words[k+1]
        m, n := len(word1), len(word2)
        var i, j int
        for ; i < m && j < n; i, j = i+1, j+1 {
            if word1[i] == word2[j] {
                continue
            } else if index[word1[i]-'a'] < index[word2[j]-'a'] {
                break
            } else {
                return false
            }
        }
        if i < m && j == n { return false }
    }
    return true
}

func newCharIndex(order string) [26]int {
    index := [26]int{}
    for i := 0; i < len(order); i++ {
        ch := order[i]
        index[ch-'a'] = i
    }
    return index
}

---

func isAlienSorted(words []string, order string) bool {
    index := map[byte]int{}
    for i := 0; i < len(order); i++ {
        index[order[i]] = i
    }
    for i := 0; i < len(words)-1; i++ {
        word1, word2 := words[i], words[i+1]
        for j := 0; j < len(word1); j++ {
            if j >= len(word2) { return false }
            if word1[j] == word2[j] { continue }
            if index[word1[j]] > index[word2[j]] {
                return false
            } else {
                break
            }
        }
    }
    return true
}
