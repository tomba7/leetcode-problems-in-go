func largestRectangleArea(heights []int) int {
    s := list.New()
    var res int
    n := len(heights)
    for i := 0; i <= n; i++ {
        for s.Len() != 0 && (i == n || heights[i] < heights[top(s)]) {
            j := s.Remove(s.Back()).(int)
            if s.Len() == 0 {
                res = max(res, heights[j] * i)
            } else {
                res = max(res, heights[j] * (i - top(s) - 1))
            }
        }
        if i < n {
            s.PushBack(i)
        }
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }
func top(s *list.List) int { return s.Back().Value.(int) }

---

func largestRectangleArea(heights []int) int {
    s := list.New()
    var res int
    n := len(heights)
    for i := 0; i < n; i++ {
        for s.Len() != 0 && heights[i] < heights[top(s)] {
            j := s.Remove(s.Back()).(int)
            if s.Len() == 0 {
                res = max(res, heights[j] * i)
            } else {
                res = max(res, heights[j] * (i - top(s) - 1))
            }
        }
        s.PushBack(i)
    }
    for s.Len() != 0 {
        j := s.Remove(s.Back()).(int)
        if s.Len() == 0 {
            res = max(res, heights[j] * n)
        } else {
            res = max(res, heights[j] * (n - top(s) - 1))
        }
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }
func top(s *list.List) int { return s.Back().Value.(int) }
