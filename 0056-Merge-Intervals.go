/*
  prev:  --------
  curr:       --------
- Sort the intervals
- Check if curr and next intervals overlap using s2 <= e1
- If so, merge into [s1, max(e1, e2)]
*/
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    stack := list.New()
    stack.PushBack(intervals[0])
    for i := 1; i < len(intervals); i++ {
        prev := stack.Back().Value.([]int) // first
        curr := intervals[i]               // second
        if curr[0] <= prev[1] {
            stack.Remove(stack.Back())
            stack.PushBack([]int{
                prev[0], max(prev[1], curr[1]),
            })
        } else {
            stack.PushBack(curr)
        }
    }
    var res [][]int
    for stack.Len() != 0 {
        res = append(res, stack.Remove(stack.Front()).([]int))
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }

---

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    n := len(intervals)
    var i int
    for j := 1; j < n; j++ {
        if intervals[i][1] >= intervals[j][0] {
            intervals[i][1] = max(intervals[i][1], intervals[j][1])
        } else {
            i++
            intervals[i] = intervals[j]
        }
    }
    return intervals[:i+1]
}

func max(x, y int) int { if x > y { return x }; return y }

---

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    n := len(intervals)
    s := list.New()
    s.PushBack(intervals[0])
    for i := 1; i < n; i++ {
        if s.Back().Value.([]int)[1] >= intervals[i][0] {
            top := s.Remove(s.Back()).([]int)
            top[1] = max(top[1], intervals[i][1])
            s.PushBack(top)
        } else {
            s.PushBack(intervals[i])
        }
    }
    var res [][]int
    for s.Len() != 0 { res = append(res, s.Remove(s.Front()).([]int)) }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }

---

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    n := len(intervals)
    for i := 1; i < n; i++ {
        if intervals[i-1][1] >= intervals[i][0] {
            intervals[i] = []int{
                intervals[i-1][0],
                max(intervals[i][1], intervals[i-1][1]),
            }
            intervals[i-1] = []int{}
        }
    }
    var res [][]int
    for i := 0; i < n; i++ {
        if len(intervals[i]) == 0 { continue }
        res = append(res, intervals[i])
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }

---

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
    if len(intervals) <= 1 { return intervals }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    result := []Interval{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        prev, curr := &result[len(result)-1], intervals[i]
        if prev.End >= curr.Start {
            prev.End = max(prev.End, curr.End)
        } else {
            result = append(result, curr)
        }
    }
    return result
}

func max(x, y int) int { if x > y { return x }; return y }
