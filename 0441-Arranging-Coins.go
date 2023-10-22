func arrangeCoins(n int) int {
    lo, hi := 0, n
    for lo <= hi {
        k := lo + (hi-lo)/2
        curr := k * (k + 1) / 2
        if curr < n {
            lo = k + 1
        } else if curr > n {
            hi = k - 1
        } else {
            return k
        }
    }
    return hi
}

---

// Brute force (Accepted)
func arrangeCoins(n int) int {
    var res int
    for i := 1; n - i >= 0; i++ {
        n -= i
        res++
    }
    return res
}
