func search(nums []int, target int) int {
    lo, hi := 0, len(nums)-1
    for lo <= hi {
        mid := lo + (hi-lo)>>1
        if nums[mid] < target {
            lo = mid+1
        } else if nums[mid] > target {
            hi = mid-1
        } else {
            return mid
        }
    }
    return -1
}
