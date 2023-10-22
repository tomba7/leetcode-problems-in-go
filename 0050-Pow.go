func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1/x
        n = -n
    }
    return myPowHelper(x, n)
}

func myPowHelper(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }
    half := myPowHelper(x, n/2)
    if n % 2 != 0 {
        return half * half * x
    }
    return half * half
}

---

// The idea is to break the power into powers of 2. This is
// because powers of 2 are easily calculated using iterated squaring.
//
// e.g  x^20
//              20 = (10100) = (10000 + 100) = (16 + 4)
//
//      x^20 = x^(16 + 4)
//           = x^16 * x^4
//
// Calculating powers of 2 through iterative squaring
//      x^y                    y
//      -------------------------
//      x^1  = x               1
//      x^2  = x   * x        10
//      x^4  = x^2 * x^2     100
//      x^8  = x^4 * x^4    1000
//      x^16 = x^8 * x^8   10000
//
func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1/x
        n = -n
    }
    result := 1.0
    product := x
    for n != 0 {
        if n & 1 == 1 {
            result *= product
        }
        product *= product
        n >>= 1
    }
    return result
}

---

// Brute force
func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1/x
        n = -n
    }
    var res = 1.0
    for i := 0; i < n; i++ {
        res *= x
    }
    return res
}
