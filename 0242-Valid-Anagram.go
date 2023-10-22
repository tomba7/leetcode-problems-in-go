const maxChar = 26

func isAnagram(s, t string) bool {
    if len(s) != len(t) { return false }
    freq := [maxChar]int{}
    for i := 0; i < len(s); i++ { freq[s[i] - 'a']++ }
    for i := 0; i < len(t); i++ {
        freq[t[i] - 'a']--
        if freq[t[i] - 'a'] < 0 { return false }
    }
    for i := 0; i < maxChar; i++ {
        if freq[i] > 0 { return false }
    }
    return true
}
