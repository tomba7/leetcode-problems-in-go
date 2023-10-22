/*
    [5,7,7,8,8,10]
         ^
    [1,1,1,1,1,2,2,2,2,2,3,3,3,3,3,3,3]
                     ^
 */
func searchRange(nums []int, target int) []int {
    lo, hi := 0, len(nums)-1
    res := []int{-1, -1}
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if target > nums[mid] {
            lo = mid + 1
        } else if target < nums[mid] {
            hi = mid - 1
        } else {
            res[0] = searchLeft(nums, lo, mid, target)
            res[1] = searchRight(nums, mid, hi, target)
            break
        }
    }
    return res
}

func searchLeft(nums []int, lo, hi, target int) int {
    for lo < hi {
        mid := lo + (hi-lo)/2
        if nums[mid] == target {
            hi = mid
        } else {
            lo = mid + 1
        }
    }
    return lo
}

func searchRight(nums []int, lo, hi, target int) int {
    for lo < hi {
        mid := lo + (hi-lo)/2 + 1
        if nums[mid] == target {
            lo = mid
        } else {
            hi = mid - 1
        }
    }
    return lo
}

---

func searchRange(nums []int, target int) []int {
    lo, hi := 0, len(nums)-1
    if lo > hi {
        return []int{-1, -1}
    }
    left := searchLeft(nums, lo, hi, target)
    if nums[left] != target {
        return []int{-1, -1}
    }
    right := searchRight(nums, lo, hi, target)
    return []int{left, right}
}

func searchLeft(nums []int, lo, hi, target int) int {
    for lo < hi {
        mid := lo + (hi-lo)/2
        if nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    return lo
}

func searchRight(nums []int, lo, hi, target int) int {
    for lo < hi {
        mid := lo + (hi-lo)/2 + 1
        if nums[mid] > target {
            hi = mid - 1
        } else {
            lo = mid
        }
    }
    return lo
}

---

func searchRange(nums []int, target int) []int {
    result := []int{-1, -1}
    if len(nums) == 0 { return result }
    // Find the left index
    lo, hi := 0, len(nums)-1
    for lo < hi {
        mid := lo + (hi-lo)/2
        if nums[mid] < target {
            lo = mid+1
        } else {
            hi = mid
        }
    }
    if nums[lo] != target { return result }
    result[0] = lo

    // Find the right index
    lo, hi = 0, len(nums)-1
    for lo < hi {
        mid := lo + (hi-lo)/2 + 1
        if nums[mid] > target {
            hi = mid-1
        } else {
            lo = mid
        }
    }
    result[1] = lo
    return result
}

