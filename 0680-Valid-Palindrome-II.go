func validPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] {
            return palindrome(s, left, right-1) || palindrome(s, left+1, right)
        }
        left++
        right--
    }
    return true
}

func palindrome(s string, left, right int) bool {
    for ; left < right; left, right = left+1, right-1 {
        if s[left] != s[right] { return false }
    }
    return true
}

---

func validPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left <= right {
        if s[left] != s[right] {
            first, second := s[left:right], s[left+1:right+1]
            return first == reverse(first) || second == reverse(second)
        }
        left++
        right--
    }
    return true
}

func reverse(s string) string {
    buf := []byte(s)
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        buf[i], buf[j] = buf[j], buf[i]
    }
    return string(buf)
}
