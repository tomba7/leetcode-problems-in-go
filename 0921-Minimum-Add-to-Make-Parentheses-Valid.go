func minAddToMakeValid(s string) int {
    var left, right int
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else {
            if left == 0 {
                right++
            }  else {
                left--
            }
        }
    }
    return left + right
}
