func reverseString(s string) string {
    b := []rune(s)
    i, j := 0, utf8.RuneCountInString(s) - 1
    for i < j {
        b[i], b[j] = b[j], b[i]
        i++
        j--
    }
    return string(b)
}
