func firstUniqChar(s string) int {
    var freq [26]int
    for i := 0; i < len(s); i++ {
        freq[s[i]-'a']++
    }
    for i := 0; i < len(s); i++ {
        if freq[s[i]-'a'] == 1 {
            return i
        }
    }
    return -1
}

---

func firstUniqChar(s string) int {
    freq := make([]int, 'z' - 'a' + 1)
    for i := 0; i < len(s); i++ {
        freq[s[i] - 'a']++
    }
    for i := 0; i < len(s); i++ {
        if freq[s[i] - 'a'] == 1 { return i }
    }
    return -1
}
