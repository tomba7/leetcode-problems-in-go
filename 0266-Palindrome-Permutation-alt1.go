func canPermutePalindrome(s string) bool {
    var freq uint64
    for i := 0; i < len(s); i++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            freq ^= 1 << (s[i] - 'a')
        } else if s[i] >= 'A' && s[i] <= 'Z' {
            freq ^= 1 << (s[i] - 'A' + 27)
        }
    }
    freq &= freq - 1
    return freq == 0
}
