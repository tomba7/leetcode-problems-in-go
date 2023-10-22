func minEatingSpeed(piles []int, h int) int {
    lo, hi := 1, math.MaxInt32
    for lo < hi {
        mid := lo + (hi-lo)/2
        if possible(piles, h, mid) {
            hi = mid
        } else {
            lo = mid + 1
        }
    }
    return lo
}

func possible(piles []int, h, k int) bool {
    var total int
    for _, p := range piles {
        total += ceil(p, k)
    }
    return total <= h
}

func ceil(p, k int) int {
    res := p/k
    if p%k != 0 { res++ }
    return res
}
