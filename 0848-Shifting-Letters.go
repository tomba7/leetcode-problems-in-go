func shiftingLetters(s string, shifts []int) string {
    for i := len(shifts)-2; i >= 0; i-- {
        shifts[i] += shifts[i+1]
    }
    var res []byte
    for i := 0; i < len(s); i++ {
        ch := (int(s[i] - 'a') + shifts[i]) % 26
        res = append(res, 'a' + byte(ch))
    }
    return string(res)
}
