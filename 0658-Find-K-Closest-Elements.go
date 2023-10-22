func findClosestElements(arr []int, k int, x int) []int {
    if x <= arr[0] {
        return arr[:k]
    } else if x >= arr[len(arr)-1] {
        return arr[len(arr)-k:]
    }
    i := sort.Search(len(arr), func(j int) bool {
        return arr[j] >= x
    })
    // lo = i-k (instead of i-k+1) because, if there is a tie,
    // the smaller elements are always preferred
    lo := max(0, i-k)
    hi := min(len(arr)-1, i+k-1)
    for hi-lo+1 > k {
        if lo < 0 || (x-arr[lo] <= arr[hi]-x) {
            hi--
        } else { // if hi > len(arr)-1 || (x-arr[lo] > arr[hi]-x) {
            lo++
        }
    }
    return arr[lo:hi+1]
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }

---

func findClosestElements(arr []int, k int, x int) []int {
    sort.Slice(arr, func(i, j int) bool {
        a, b := abs(arr[i] - x), abs(arr[j] - x)
        // The following can also be shortened as
        //    return a < b || a == b && arr[i] < arr[j]
        if a == b {
            return arr[i] < arr[j]
        } else {
            return a < b
        }
    })
    arr = arr[:k]
    sort.Ints(arr)
    return arr
}

func abs(x int) int { if x < 0 { return -x }; return x }

---

func findClosestElements(arr []int, k int, x int) []int {
    lo, hi := 0, len(arr)-k
    for lo < hi {
        mid := lo + (hi-lo)/2
        if x - arr[mid] > arr[mid+k] - x {
            lo = mid+1
        } else {
            hi = mid
        }
    }
    return arr[lo:lo+k]
}
