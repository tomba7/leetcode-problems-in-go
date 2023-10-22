//
// Range:   1 - 9     10 - 99     100 - 999    1000 - 9999
//
// Digits:  1         2           3            4
//          9         90          900          9000
//
// Total:   1 * 9     2 * 90      3 * 900      4 * 9000
//          9         180         2700         36000
//
func findNthDigit(n int) int {
    base, digits := 1, 1
    for digits * base * 9 < n {
        n -= digits * base * 9
        digits++
        base *= 10
    }
    number := base + (n - 1) / digits
    index := (n - 1) % digits
    result := strconv.Itoa(number)
    digit, _ := strconv.Atoi(string(result[index]))
    return digit
}
