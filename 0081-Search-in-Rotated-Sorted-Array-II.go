func search(nums []int, target int) bool {
    lo, hi := 0, len(nums) - 1
    for lo <= hi {
        mid := lo + (hi - lo)/2
        if nums[mid] == target {
            return true
        } else if nums[mid] < nums[hi] {
            // Right has normal sorted order. Check for target within that range.
            if nums[mid] < target && target <= nums[hi] {
                lo = mid + 1
            } else {
                hi = mid - 1
            }
        } else if nums[mid] > nums[hi] {
            // Left has normal sorted order. Check for target within that range.
            if nums[lo] <= target && target < nums[mid] {
                hi = mid - 1
            } else {
                lo = mid + 1
            }
        } else {
            // nums[mid] == nums[hi], which means we've duplicates.
            if nums[lo] != nums[mid] {
                // If leftmost element is different then the left side must have the target.
                hi = mid - 1
            } else {
                // If not, either side may have the target. So search both.
                lo++
            }
        }
    }
    return false
}

---

func search(nums []int, target int) bool {
    lo, hi := 0, len(nums)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if nums[mid] == target {
            return true
        } else if nums[mid] < nums[hi] {
            if nums[mid] < target && target <= nums[hi] {
                lo = mid+1
            } else {
                hi = mid-1
            }
        } else if nums[mid] > nums[hi] {
            if nums[lo] <= target && target < nums[mid] {
                hi = mid-1
            } else {
                lo = mid+1
            }
        } else {
            if nums[lo] == nums[mid] {
                lo++
            } else {
                hi = mid-1
            }
        }
    }
    return false
}
