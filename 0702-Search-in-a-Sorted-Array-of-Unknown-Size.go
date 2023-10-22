func search(reader ArrayReader, target int) int {
    // We can also set the hi to 20000 since all elements are unique
    // and are in the range of -9999 ~ 9999
    lo, hi := 0, math.MaxInt32
    for lo <= hi {
        mid := lo + (hi-lo)/2
        x := reader.get(mid)
        if x == 2147483647 || x > target {
            hi = mid-1
        } else if x < target {
            lo = mid+1
        } else {
            return mid
        }
    }
    return -1
}
