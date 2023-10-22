/*
    dividend = 357
    divisor = 5
    
    [1] x = 357                     [2] x = 357 - 320 = 37
    ------------------------        ------------------------
    Total  Copies     PowerOfTwo
        5   1 copy    2^0               5   1 copy    2^0
       10   2 copies  2^1              10   2 copies  2^1
       20   4 copies  2^2              20   4 copies  2^2
       40   8 copies  2^3              40   <-- too large
       80  16 copies  2^4
      160  32 copies  2^5
      320  64 copies  2^6
      640  <-- too large
    
    [3] x = 37 - 20 = 17            [4] x = 17 - 10 = 7
    ------------------------        ------------------------
        5   1 copy    2^0               5   1 copy    2^0
       10   2 copies  2^1              10   <-- too large
       20   <-- too large
                                    [5] x = 7 - 5 = 2
                                    ------------------------
                                        5   <-- too large
    result = 2^6 + 2^2 + 2^1 + 2^0 = 71
*/
func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }
    sign := 1
    if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
        sign = -1
    }
    dividend, divisor = abs(dividend), abs(divisor)
    var res int
    for dividend >= divisor {
        numCopies := 1
        sum := divisor
        for sum << 1 <= dividend {
            numCopies <<= 1
            sum <<= 1
        }
        dividend -= sum
        res += numCopies
    }
    return sign * res
}

func abs(x int) int { if x < 0 { return -x }; return x }

---

func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 { return math.MaxInt32 }
    var result, sign int = 0, 1
    if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
        sign = -1
    }
    dividend, divisor = abs(dividend), abs(divisor)
    for dividend >= divisor {
        total := divisor
        powerOfTwo := 1     // Number of copies (of divisor)
        for total << 1 <= dividend {
            powerOfTwo <<= 1
            total <<= 1
        }
        result += powerOfTwo
        dividend -= total
    }
    return result * sign
}

func abs(x int) int { if x < 0 { return -x }; return x }

---

// Brute force 1 (with computation in -ve space to avoid handling edge cases
func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 { return math.MaxInt32 }
    var negatives = 2
    if dividend > 0 {
        negatives--
        dividend = -dividend
    }
    if divisor > 0 {
        negatives--
        divisor = -divisor
    }
    var result int
    for dividend - divisor <= 0 {
        dividend -= divisor
        result--
    }
    if negatives != 1 { return -result }
    return result
}

---

// Brute force 2
func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 { return math.MaxInt32 }
    var result int
    sign := 1
    if dividend < 0 && divisor > 0 ||dividend > 0 && divisor < 0 {
        sign = -1
    }
    dividend, divisor = abs(dividend), abs(divisor)
    for dividend - divisor >= 0 {
        dividend -= divisor
        result++
    }
    return result * sign
}

func abs(x int) int { if x < 0 { return -x }; return x }

---

func divide(x, y int) int {
    if x == math.MinInt32 && y == -1 { return math.MaxInt32 }
    var result, sign int = 0, 1
    if x < 0 && y > 0 || x > 0 && y < 0 { sign = -1 }
    x, y = abs(x), abs(y)
    for x >= y {
        total, count := y, 1
        for total << 1 <= x {
            count <<= 1
            total <<= 1
        }
        result += count
        x -= total
    }
    return result * sign
}

func abs(x int) int { if x < 0 { return -x }; return x }
