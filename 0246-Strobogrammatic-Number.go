func isStrobogrammatic(num string) bool {
    if len(num) == 0 {return false}
    match := map[byte]byte{'0': '0', '1': '1', '6': '9', '8': '8', '9': '6'}
    for i, j := 0, len(num) - 1; i <= j; i, j = i + 1, j - 1 {
        if match[num[i]] != num[j] {return false}
    }
    return true
}
