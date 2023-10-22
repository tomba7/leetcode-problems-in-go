func search(nums []int, target int) int {
    lo, hi := 0, len(nums)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < nums[hi] {
            if nums[mid] < target && target <= nums[hi] {
                lo = mid+1
            } else {
                hi = mid-1
            }
        } else {
            if nums[lo] <= target && target < nums[mid] {
                hi = mid-1
            } else {
                lo = mid+1
            }
        }
    }
    return -1
}

---

func search(nums []int, target int) int {
    lo, hi := 0, len(nums) - 1
    for lo <= hi {
        mid := lo + (hi - lo)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > nums[hi] {
            // Left has normal sorted order. Check for target within that range.
            if nums[lo] <= target && target < nums[mid] {
                hi = mid - 1
            } else {
                lo = mid + 1
            }
        } else {
            // Right has normal sorted order. Check for target within that range.
            if nums[mid] < target && target <= nums[hi] {
                lo = mid + 1
            } else {
                hi = mid - 1
            }
        }
    }
    return -1
}
