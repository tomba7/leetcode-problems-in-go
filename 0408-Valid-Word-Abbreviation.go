func validWordAbbreviation(word, abbr string) bool {
    var i, j int
    for i < len(word) && j < len(abbr) {
        if word[i] == abbr[j] {
            i++
            j++
            continue
        }
        if !digit(abbr[j]) || abbr[j] == '0' {
            return false
        }
        start := j
        for ; j < len(abbr) && digit(abbr[j]); j++ {}
        skip, _ := strconv.Atoi(abbr[start:j])
        i += skip
    }
    return i == len(word) && j == len(abbr)
}

func digit(ch byte) bool { return '0' <= ch && ch <= '9' }
