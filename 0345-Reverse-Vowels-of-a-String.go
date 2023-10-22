func reverseVowels(s string) string {
    b := []byte(s)
    i, j := 0, len(s) - 1
    for i < j {
        for i < j && !isVowel(b[i]) { i++ }
        for i < j && !isVowel(b[j]) { j-- }
        if i >= j { break }
        b[i], b[j] = b[j], b[i]
        i++
        j--
    }
    return string(b)
}

func isVowel(c byte) bool {
    switch c { case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U': return true }
    return false
}
