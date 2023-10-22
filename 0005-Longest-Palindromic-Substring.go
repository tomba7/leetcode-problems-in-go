func longestPalindrome(s string) string {
    var begin, end int
    for i := 0; i < len(s); i++ {
        odd := expandOnBothSides(s, i, i)
        even := expandOnBothSides(s, i, i + 1)
        length := max(odd, even)
        if length > end - begin {
            begin = i - (length - 1)/2
            end = i + length/2
        }
    }
    return s[begin : end + 1]
}

func expandOnBothSides(s string, i, j int) int {
    for i >= 0 && j < len(s) && s[i] == s[j] {
        i--
        j++
    }
    return j - i - 1
}

func max(x, y int) int { if x > y { return x }; return y }
