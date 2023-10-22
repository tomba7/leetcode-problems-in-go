/*
     0   1
    [1, -1]
         ^
    deq = [-1]
    res = [1, -1]
 */
func maxSlidingWindow(nums []int, k int) []int {
    var res []int
    deq := list.New()
    for i := 0; i < len(nums); i++ {
        if deq.Len() != 0 && deq.Front().Value.(int) == i - k {
            deq.Remove(deq.Front())
        }
        for deq.Len() != 0 && nums[deq.Back().Value.(int)] <= nums[i] {
            deq.Remove(deq.Back())
        }
        deq.PushBack(i)
        if i >= k - 1 {
            res = append(res, nums[deq.Front().Value.(int)])
        }
    }
    return res
}

---

func maxSlidingWindow(nums []int, w int) []int {
    n := len(nums)
    res := make([]int, n - w + 1)
    deq := list.New()
    for i := 0; i < n; i++ {
        for deq.Len() != 0 && nums[deq.Back().Value.(int)] <= nums[i] {
            deq.Remove(deq.Back())
        }
        deq.PushBack(i)
        if i >= w-1 {
            if deq.Front().Value.(int) <= i-w {
                deq.Remove(deq.Front())
            }
            res[i-w+1] = nums[deq.Front().Value.(int)]
        }
    }
    return res
}

---

func maxSlidingWindow(nums []int, k int) []int {
    window := list.New()
    result := []int{}
    for i := 0; i < len(nums); i++ {
        if window.Len() != 0 && window.Front().Value.(int) < i - k + 1 {
            window.Remove(window.Front())
        }
        for window.Len() != 0 && nums[window.Back().Value.(int)] < nums[i] {
            window.Remove(window.Back())
        }
        window.PushBack(i)
        if i < k - 1 { continue }
        result = append(result, nums[window.Front().Value.(int)])
    }
    return result
}
