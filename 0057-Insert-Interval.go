func insert(intervals [][]int, newInterval []int) [][]int {
    deque := list.New()
    deque.PushBack(newInterval)
    for _, curr := range intervals {
        back := deque.Back().Value.([]int)
        if overlaps(back, curr) {
            deque.Remove(deque.Back())
            deque.PushBack([]int{
                min(back[0], curr[0]),
                max(back[1], curr[1]),
            })
        } else {
            if curr[0] < back[0] {
                deque.InsertBefore(curr, deque.Back())
            } else {
                deque.PushBack(curr)
            }
        }
    }
    var res [][]int
    for deque.Len() != 0 {
        res = append(res, deque.Remove(deque.Front()).([]int))
    }
    return res
}

func overlaps(interval1, interval2 []int) bool {
    s1, e1 := interval1[0], interval1[1]
    s2, e2 := interval2[0], interval2[1]
    return s2 <= e1 && s1 <= e2
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }
