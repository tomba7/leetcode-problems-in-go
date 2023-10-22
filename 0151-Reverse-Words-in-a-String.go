func reverseWords(str string) string {
    s := []byte(str)
    var i, j, k int
    for ; k < len(s); k++ {
        if s[k] == ' ' { continue }
        if j != 0 && s[k-1] == ' ' {
            s[j] = ' '; j++
            i = j
        }
        s[j] = s[k]; j++
        if k+1 == len(s) || s[k+1] == ' ' {
            reverse(s, i, j-1)
        }
    }
    s = s[:j]
    reverse(s, 0, len(s)-1)
    return string(s)
}

func reverse(s []byte, i, j int) {
    for ; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
