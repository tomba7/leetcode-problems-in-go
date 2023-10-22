func canAttendMeetings(intervals [][]int) bool {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    for i := 0; i < len(intervals)-1; i++ {
        if overlaps(intervals[i], intervals[i+1]) { return false }
    }
    return true
}

func overlaps(i1, i2 []int) bool { return i2[0] < i1[1] }
