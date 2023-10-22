func reverseString(s string) string {
    b := []byte(s)
    i, j := 0, len(s) - 1
    for i < j {
        b[i], b[j] = b[j], b[i]
        i++
        j--
    }
    return string(b)
}
