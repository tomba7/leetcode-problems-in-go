func isOneEditDistance(s string, t string) bool {
    m, n := len(s), len(t)
    if m > n { return isOneEditDistance(t, s) }
    if n - m > 1 { return false }
    for i := 0; i < m; i++ {
        if s[i] != t[i] {
            if m == n {
                return s[i+1:] == t[i+1:]
            } else {
                return s[i:] == t[i+1:]
            }
        }
    }
    return m + 1 == n
}

---

func isOneEditDistance(s, t string) bool {
    if len(s) > len(t) { s, t = t, s }
    if len(t) - len(s) > 1 || s == t { return false }
    diffFound := false
    for i, j := 0, 0; i < len(s); j++ {
        if s[i] != t[j] {
            if diffFound { return false }
            diffFound = true
            if len(s) == len(t) { i++ }
        } else {
            i++
        }
    }
    return true
}

---

func isOneEditDistance(s string, t string) bool {
    if math.Abs(float64(len(s) - len(t))) > 1 || s == t { return false }
    if len(s) > len(t) { s, t = t, s }
    i, j := 0, 0
    replacementFound := false
    for i < len(s) && j < len(t) {
        if s[i] != t[j] {
            if len(s) == len(t) {
                if replacementFound {
                    return false
                }
                replacementFound = true
                i++
                j++

            } else {
                if i != j {
                    return false
                }
                j++
            }
        } else {
            i++
            j++
        }
    }
    return true
}
