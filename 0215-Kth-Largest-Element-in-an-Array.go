/* 4:36 PM
   [3,2,1,5,6,4]
   [1,2,3,4,5,6]
   k = 2
   6 - 2 = 4
 */
func findKthLargest(nums []int, k int) int {
    k = len(nums) - k
    lo, hi := 0, len(nums)-1
    for lo <= hi {
        mid := partition(nums, lo, hi)
        if k == mid {
            return nums[mid]
        } else if k < mid {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    return -1
}

func partition(nums []int, lo, hi int) int {
    pivot := nums[hi]
    i := lo
    for j := lo; j < hi; j++ {
        if nums[j] < pivot {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    nums[i], nums[hi] = nums[hi], nums[i]
    return i
}

---

func findKthLargest(nums []int, k int) int {
    h := &Heap{}
    for i := 0; i < len(nums); i++ {
        if i < k {
            heap.Push(h, nums[i])
        } else {
            if nums[i] >= h.Top() {
                heap.Pop(h)
                heap.Push(h, nums[i])
            }
        }
    }
    return h.Top()
}


type Heap []int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Top() int { return (*h)[0] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[:n-1]
    return x
}

---

func findKthLargest(nums []int, k int) int {
    rand.Seed(time.Now().UnixNano())
    lo, hi := 0, len(nums)-1
    target := len(nums)-k
    for lo <= hi {
        mid := partition(nums, lo, hi)
        if mid < target {
            lo = mid + 1
        } else if mid > target {
            hi = mid - 1
        } else {
            return nums[target]
        }
    }
    return -1
}

func partition(nums []int, lo, hi int) int {
    pivotIndex := rand.Intn(hi-lo+1) + lo
    pivot := nums[pivotIndex]
    i, j, k := lo, lo, hi
    for j <= k {
        if nums[j] == pivot {
            j++
        } else if nums[j] < pivot {
            swap(nums, i, j)
            i++
            j++
        } else {
            swap(nums, j, k)
            k--
        }
    }
    return i
}

func swap(nums []int, i, j int) { nums[i], nums[j] = nums[j], nums[i] }

---


func findKthLargest(nums []int, k int) int {
    rand.Seed(time.Now().UnixNano())
    lo, hi, target := 0, len(nums)-1, len(nums)-k
    for lo <= hi {
        mid := partition(nums, lo, hi)
        if mid == target {
            return nums[mid]
        } else if mid < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return -1
}

func partition(nums []int, lo, hi int) int {
    i, j, k := lo, lo, hi
    pivotIndex := lo + rand.Intn(hi-lo+1)
    pivot := nums[pivotIndex]
    for j <= k {
        if nums[j] < pivot {
            nums[i], nums[j] = nums[j], nums[i]
            i, j = i+1, j+1
        } else if nums[j] > pivot {
            nums[j], nums[k] = nums[k], nums[j]
            k--
        } else {
            j++
        }
    }
    return i
}

---

func findKthLargest(nums []int, k int) int {
    if k <= 0 || k > len(nums) { return -1 }
    h := &MinHeap{}
    for i := 0; i < len(nums); i++ {
        if i < k {
            heap.Push(h, nums[i])
        } else {
            if nums[i] < (*h)[0] { continue }
            heap.Pop(h)
            heap.Push(h, nums[i])
        }
    }
    return (*h)[0]
}

// Supporting data structures
type MinHeap []int
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(value interface{}) { *h = append(*h, value.(int)) }
func (h *MinHeap) Pop() interface{} { x := (*h)[len(*h) - 1]; *h = (*h)[:len(*h) - 1]; return x }

---

/*
 - Create a min heap (and a hash map to check dups)
 - Add the first k elements. Make sure we do not insert dups
 - Past k, check if nums[i] > heap.Peek()
 - If yes, the heap.Pop() + heap.Push(nums[i])
 - If no, then continue
 - Return the top of the heap
 
 [3,2,3,1,2,4,5,5,6], k = 4
                  ^
 h = 6, 5, 5, 4
 res = 4

 */
func findKthLargest(nums []int, k int) int {
    h := &MinHeap{}
    for i := 0; i < len(nums); i++ {
        if i < k {
            heap.Push(h, nums[i])
        } else {
            if nums[i] > (*h)[0] {
                heap.Pop(h)
                heap.Push(h, nums[i])
            }
        }
    }
    return (*h)[0]
}

type MinHeap []int
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[:n-1]
    return x
}
