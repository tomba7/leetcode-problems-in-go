func lengthOfLongestSubstring(s string) int {
    index := map[byte]int{}
    var startIndex, res int
    for j := 0; j < len(s); j++ {
        if i, exist := index[s[j]]; exist && i >= startIndex {
            startIndex = i + 1
        }
        res = max(res, j - startIndex + 1)
        index[s[j]] = j
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }

---

func lengthOfLongestSubstring(s string) int {
    var i, j, res int
    seen := map[byte]int{}
    for ; j < len(s); j++ {
        ch := s[j]
        k, exist := seen[ch]
        if exist && k >= i {
            i = k + 1
        }
        seen[ch] = j
        res = max(res, j-i+1)
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }

---

func lengthOfLongestSubstring(s string) int {
    indexMap := make(map[byte]int)
    startIndex, subLen, maxSubLen := 0, 0, 0
    for i := 0; i < len(s); i++ {
        indx, exist := indexMap[s[i]]
        if exist && indx >= startIndex {
            subLen = i - indx
            startIndex = indx + 1
        } else {
            subLen++
        }
        indexMap[s[i]] = i
        maxSubLen = max(subLen, maxSubLen)
    }
    return maxSubLen
}

func max(x, y int) int { if x > y { return x }; return y }

---

func lengthOfLongestSubstring(s string) int {
    indexMap := [128]int{}
    startIndex, subLen, maxSubLen := 0, 0, 0
    for i := 0; i < len(indexMap); i++ { indexMap[i] = -1 }

    for i := 0; i < len(s); i++ {
        indx := indexMap[s[i]]
        if indx >= startIndex {
            subLen = i - indx
            startIndex = indx + 1
        } else {
            subLen++
        }
        indexMap[s[i]] = i
        maxSubLen = max(subLen, maxSubLen)
    }
    return maxSubLen
}

func max(x, y int) int { if x > y { return x }; return y }
