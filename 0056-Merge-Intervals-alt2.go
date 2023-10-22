/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    merged := []Interval{}
    for _, curr := range intervals {
        if n := len(merged); n != 0 && merged[n-1].End >= curr.Start {
            merged[n-1].End = max(merged[n-1].End, curr.End)
        } else {
            merged = append(merged, curr)
        }
    }
    return merged
}

func max(x, y int) int { if x > y { return x }; return y }
