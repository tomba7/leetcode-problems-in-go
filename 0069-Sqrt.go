func mySqrt(x int) int {
    lo, hi := 1, x
    // the following works as well
    //  lo, hi := 0, x 
    var result int
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if mid*mid == x {
            return mid
        } else if mid*mid > x {
            hi = mid-1
        } else {
            lo = mid+1
            result = mid
        }
    }
    return result
}

---

func mySqrt(x int) int {
    if x == 0 {return 0}
    lo, hi, result := 1, x, 0
    for lo <= hi {
        mid := lo + (hi - lo)/2
        if mid * mid <= x {
            lo = mid + 1
            result = mid
        } else {
            hi = mid - 1
        }
    }
    return result
}
