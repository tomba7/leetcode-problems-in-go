func reverse(x int) int {
    sign, reverse := 1, 0
    if x < 0 { sign, x = -1, -x }
    for x != 0 {
        if reverse > 214748364 { return 0 } // or 0x7FFFFFFF/10
        reverse = 10 * reverse + x % 10
        x /= 10
    }
    return reverse * sign
}
