func isPalindrome(x int) bool {
    if x < 0 { return false }
    divisor := 1
    for x / divisor >= 10 {
        divisor *= 10
    }
    for x != 0 {
        left := x / divisor
        right := x % 10
        if (left != right) { return false }
        x %= divisor
        x /= 10
        divisor /= 100
    }
    return true
}
