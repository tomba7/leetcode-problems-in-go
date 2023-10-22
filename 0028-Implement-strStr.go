func strStr(haystack, needle string) int {
    if len(needle) == 0 {return 0 } // Works with or without this check
    for i := 0; i < len(haystack) - len(needle) + 1; i++ {
        j := 0
        for ; j < len(needle); j++ {
            if haystack[i + j] != needle[j] { break }
        }
        if j == len(needle) {return i}
    }
    return -1
}

---

func strStr(haystack string, needle string) int {
    for i := 0; i < len(haystack)-len(needle)+1; i++ {
        j := 0
        for j < len(needle) && needle[j] == haystack[i+j] { j++ }
        if j == len(needle) { return i }
    }
    return -1
}

---

func strStr(h, n string) int {
    for i := 0; i < len(h)-len(n)+1; i++ {
        var j int
        for ; j < len(n); j++ {
            if h[i+j] != n[j] { break }
        }
        if j == len(n) { return i }
    }
    return -1
}

---

func strStr(s, p string) int {
    n, m := len(s), len(p)
    for i := 0; i <= n-m; i++ {
        found := true
        for j := 0; j < m; j++ {
            if s[i+j] != p[j] {
                found = false
                break
            }
        }
        if found { return i }
    }
    return -1
}

---

func strStr(s string, p string) int {
    n, m := len(s), len(p)
    for i := 0; i < n-m+1; i++ {
        found := true
        for j := 0; j < m; j++ {
            if s[i+j] != p[j] {
                found = false
                break
            }
        }
        if found { return i }
    }
    return -1
}
