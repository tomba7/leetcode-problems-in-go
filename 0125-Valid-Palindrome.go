func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
    for i < j {
        for i < j && !isAlphaNumeric(s[i]) {
            i++
        }
        for i < j && !isAlphaNumeric(s[j]) {
            j--
        }
        if toLower(s[i]) != toLower(s[j]) {
            return false
        }
        i++
        j--
    }
    return true
}

func toLower(ch byte) byte {
    if ch >= 'A' && ch <= 'Z' {
        return 'a' +  ch - 'A'
    }
    return ch
}

func isAlphaNumeric(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9'
}

---

func isPalindrome(s string) bool {
    i, j := 0, len(s) - 1
    for i < j {
        if !isAlphaNumeric(s[i]) {
            i++
            continue
        }
        if !isAlphaNumeric(s[j]) {
            j--
            continue
        }
        if toLower(s[i]) != toLower(s[j]) { return false }
        i++
        j--
    }
    return true
}

func isAlphaNumeric(b byte) bool {
    if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' || b >= '0' && b <= '9' { return true }
    return false
}

func toLower(b byte) byte {
    if b >= 'A' && b <= 'Z' { return 'a' + b - 'A' }
    return b
}

---

func isPalindrome(s string) bool {
    i, j := 0, len(s) - 1
    for i < j {
        if !isAlphaNumeric(s[i]) { i++; continue }
        if !isAlphaNumeric(s[j]) { j--; continue }
        if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) { return false }
        i, j = i + 1, j - 1
    }
    return true
}

func isAlphaNumeric(b byte) bool {
    return unicode.IsLetter(rune(b)) || unicode.IsNumber(rune(b))
}

---

func isPalindrome(s string) bool {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        for i < j && !unicode.IsDigit(rune(s[i])) && !unicode.IsLetter(rune(s[i])) { i++ }
        for i < j && !unicode.IsDigit(rune(s[j])) && !unicode.IsLetter(rune(s[j])) { j-- }
        if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) { return false }
    }
    return true
}
