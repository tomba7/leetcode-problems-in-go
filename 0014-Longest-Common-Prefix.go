func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 { return "" }
    prefix := strs[0]
    var i int
    for i = 0; i < len(prefix); i++ {
        for j := 1; j < len(strs); j++ {
            if len(strs[j]) == i || strs[j][i] != prefix[i] {
                return prefix[:i]
            }
        }
    }
    return prefix[:i]
}

---

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 { return "" }
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        for strings.Index(strs[i], prefix) != 0 {
            prefix = prefix[:len(prefix)-1]
            if len(prefix) == 0 { return "" }
        }
    }
    return prefix
}
