func canPermutePalindrome(s string) bool {
    var freq [26]int
    for i := 0; i < len(s); i++ {
        ch := s[i]
        freq[ch-'a']++
    }
    var count int
    for ch := 0; ch < 26 && count <= 1; ch++ {
        count += freq[ch] % 2
    }
    return count <= 1
}

---

func canPermutePalindrome(s string) bool {
    var freq [26]int
    for i := 0; i < len(s); i++ {
        ch := s[i]
        freq[ch-'a']++
    }
    var oddFound bool
    for _, count := range freq {
        if count % 2 == 1 {
            if oddFound { return false }
            oddFound = true
        }
    }
    return true
}

---

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
