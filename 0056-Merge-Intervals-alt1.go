/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
    if len(intervals) <= 1 {return intervals}
    sort.Sort(Intervals(intervals))
    result := []Interval{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        if intervals[i].Start <= result[len(result) - 1].End {
            result[len(result) - 1].End = max(result[len(result) - 1].End, intervals[i].End)
            continue
        }
        result = append(result, intervals[i])
    }
    return result
}

type Intervals []Interval
func (m Intervals) Less(i, j int) bool {return m[i].Start < m[j].Start}
func (m Intervals) Swap(i, j int) {m[i], m[j] = m[j], m[i]}
func (m Intervals) Len() int {return len(m)}

func max(x, y int) int {if x > y {return x}; return y}
