func trap(height []int) int {
    n := len(height)
    if n == 0 { return 0 }
    left, right := make([]int, n), make([]int, n)
    left[0] = height[0]
    for i := 1; i < n; i++ {
        left[i] = max(height[i], left[i-1])
    }
    right[n-1] = height[n-1]
    for i := n-2; i >= 0; i-- {
        right[i] = max(height[i], right[i+1])
    }
    var total int
    for i := 1; i < n-1; i++ {
        total += min(left[i], right[i])-height[i]
    }
    return total
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }

---

func trap(height []int) int {
    left, right := 0, len(height)-1
    var res, leftMax, rightMax int
    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            } else {
                res += leftMax - height[left]
            }
            left++
        } else {
            if height[right] >= rightMax {
                rightMax = height[right]
            } else {
                res += rightMax - height[right]
            }
            right--
        }
    }
    return res
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }

---

func trap(heights []int) int {
    var res int
    s := list.New()
    for i := 0; i < len(heights); i++ {
        for s.Len() != 0 && heights[s.Back().Value.(int)] < heights[i] {
            prev := s.Remove(s.Back()).(int)
            if s.Len() == 0 {
                break
            }
            width := i - s.Back().Value.(int) - 1
            height := min(heights[i], heights[s.Back().Value.(int)]) - heights[prev]
            res += width * height
        }
        if s.Len() != 0 && heights[s.Back().Value.(int)] == heights[i] {
            // This is optional. Logic still works without this block
            s.Remove(s.Back())
        }
        s.PushBack(i)
    }
    return res
}

func min(x, y int) int { if x < y { return x }; return y }
