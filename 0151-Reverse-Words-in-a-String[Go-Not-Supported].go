func reverseWords(s string) string {
    result := []byte{}
    j := len(s)
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == ' ' {
            j = i
        } else if i == 0 || s[i-1] == ' ' {
            if len(result) != 0 {
                result = append(result, ' ')
            }
            result = append(result, s[i:j]...)
        }
    }
    return string(result)
}
