func isPerfectSquare(num int) bool {
    lo, hi := 0, num
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if mid*mid == num {
            return true
        } else if mid*mid < num {
            lo = mid+1
        } else {
            hi = mid-1
        }
    }
    return false
}
