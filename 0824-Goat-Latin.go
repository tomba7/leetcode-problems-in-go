func toGoatLatin(sentence string) string {
    words := strings.Split(sentence, " ")
    var res string
    for i := 0; i < len(words); i++ {
        goatWord := ""
        trail := bytes.Repeat([]byte{'a'}, i+1)
        if vowel(words[i][0]) {
            goatWord = words[i]
        } else {
            goatWord = words[i][1:] + words[i][0:1]
        }
        goatWord += "ma" + string(trail)
        if i != 0 { res += " " }
        res += goatWord
    }
    return res
}

func vowel(ch byte) bool {
    switch ch {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        return true
    }
    return false
}
