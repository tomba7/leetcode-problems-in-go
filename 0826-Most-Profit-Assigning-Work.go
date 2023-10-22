func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    var jobs [][]int
    n := len(profit)
    for i := 0; i < n; i++ {
        jobs = append(jobs, []int{
            difficulty[i], profit[i],
        })
    }

    sort.Slice(jobs, func(i, j int) bool {
        return jobs[i][0] < jobs[j][0]
    })
    sort.Ints(worker)

    var res, i, best int
    for _, ability := range worker {
        for i < n && ability >= jobs[i][0] {
            best = max(best, jobs[i][1])
            i++
        }
        res += best
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }
