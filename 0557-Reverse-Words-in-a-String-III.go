func reverseWords(str string) string {
    s := []byte(str)
    var i, j int
    for ; j < len(s); j++ {
        if j+1 == len(s) || s[j+1] == ' ' {
            reverse(s, i, j)
            i = j+2
        }
    }
    return string(s)
}

func reverse(s []byte, i, j int) {
    for ; i < j; i, j = i+1, j-1 { s[i], s[j] = s[j], s[i] }
}
