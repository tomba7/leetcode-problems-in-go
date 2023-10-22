func longestRepeatingSubstring(s string) int {
    n := len(s)
    lcp := make([][]int, n+1)
    for i := range lcp {
        lcp[i] = make([]int, n+1)
    }
    var lrsLen int
    for i := n-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if i >= j { continue }
            if s[i] == s[j] {
                lcp[i][j] = 1 + lcp[i+1][j+1]
            }
            if lcp[i][j] > lrsLen {
                lrsLen = lcp[i][j]
            }
        }
    }
    return lrsLen
}

---

func longestRepeatingSubstring(s string) int {
    n := len(s)
    left, right := 1, n
    for left <= right {
        l := left + (right-left)/2
        if search(l, s) != -1 {
            left = l + 1
        } else {
            right = l - 1
        }
    }
    return left - 1
}

func search(l int, s string) int {
    n := len(s)
    seen := map[string]bool{}
    for i := 0; i < n-l+1; i++ {
        if seen[s[i:i+l]] { return i }
        seen[s[i:i+l]] = true
    }
    return -1
}

---

func longestRepeatingSubstring(s string) int {
    n := len(s)
    left, right := 1, n
    for left <= right {
        l := left + (right-left)/2
        if search(l, s) != -1 {
            left = l + 1
        } else {
            right = l - 1
        }
    }
    return left - 1
}

func search(l int, s string) int {
    n := len(s)
    seen := map[uint32]bool{}
    var h uint32
    for i := 0; i < n-l+1; i++ {
        h = hash(s[i:i+l])
        if seen[h] { return i }
        seen[h] = true
    }
    return -1
}

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}
