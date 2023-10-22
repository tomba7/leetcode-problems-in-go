/*
       0, 1, 2, 3, 4
    [-10,-5, 0, 3, 7]
                   
       0, 1, 2, 3, 4, 5 
    [-10,-5, 3, 4, 7, 9]

     0, 1, 2, 3, 4
    [0, 2, 5, 8, 17]

       0. 1. 2. 3. 4. 
    [-10,-5, 2, 3, 7]
 */
func fixedPoint(arr []int) int {
    lo, hi := 0, len(arr)-1
    for lo < hi {
        mid := lo + (hi-lo)/2
        if arr[mid] > mid {
            hi = mid - 1
        } else if arr[mid] < mid {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    if arr[lo] != lo {
        return -1
    }
    return arr[lo]
}

---

func fixedPoint(arr []int) int {
    res := -1
    lo, hi := 0, len(arr)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if arr[mid] >= mid {
            if arr[mid] == mid {
                res = mid
            }
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    return res
}
