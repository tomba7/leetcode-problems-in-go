func lengthOfLongestSubstringKDistinct(s string, k int) int {
    var i, j, res, uniques int
    freq := make(map[byte]int)
    for ; j < len(s); j++ {
        if freq[s[j]] == 0 {
            uniques++
        }
        freq[s[j]]++
        for uniques > k {
            freq[s[i]]--
            if freq[s[i]] == 0 {
                uniques--
            }
            i++
        }
        res = max(res, j-i+1)
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }

---

func lengthOfLongestSubstringKDistinct(s string, k int) int {
    var i, j, res int
    freq := make(map[byte]int)
    for ; j < len(s); j++ {
        ch := s[j]
        freq[ch]++
        for ; len(freq) > k; i++ {
            freq[s[i]]--
            if freq[s[i]] == 0 {
                delete(freq, s[i])
            }
        }
        res = max(res, j-i+1)
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }
