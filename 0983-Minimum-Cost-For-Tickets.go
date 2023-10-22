func mincostTickets(days, costs []int) int {
    dp := make([]int, len(days)+1)
    for i := len(days)-1; i >= 0; i-- {
        oneDayCost := costs[0] + dp[i+1]
        var j int
        for j = i; j < len(days) && days[j] < days[i] + 7; j++ {}
        sevenDaysCost := costs[1] + dp[j]
        for j = i; j < len(days) && days[j] < days[i] + 30; j++ {}
        oneMonthCost := costs[2] + dp[j]
        dp[i] = min(oneDayCost, min(sevenDaysCost, oneMonthCost))
    }
    return dp[0]
}

func min(x, y int) int { if x < y { return x }; return y }

---

func mincostTickets(days, costs []int) int {
    memo := make([]int, 366)
    for i := range memo {
        memo[i] = -1
    }
    return mincostTicketsHelper(days, costs, 0, memo)
}

func mincostTicketsHelper(days, costs []int, i int, memo []int) int {
    if i == len(days) { return 0 }
    if memo[i] != -1 {
        return memo[i]
    }
    oneDayCost := costs[0] + mincostTicketsHelper(days, costs, i+1, memo)

    var j int
    for j = i; j < len(days) && days[j] < days[i] + 7; j++ {}
    sevenDaysCost := costs[1] + mincostTicketsHelper(days, costs, j, memo)

    for j = i; j < len(days) && days[j] < days[i] + 30; j++ {}
    oneMonthCost := costs[2] + mincostTicketsHelper(days, costs, j, memo)

    memo[i] = min(oneDayCost, min(sevenDaysCost, oneMonthCost))
    return memo[i]
}

func min(x, y int) int { if x < y { return x }; return y }

---

// Recursive Appraoch (TLE)

func mincostTickets(days, costs []int) int {
    return mincostTicketsHelper(days, costs, 0)
}

func mincostTicketsHelper(days, costs []int, i int) int {
    if i == len(days) { return 0 }
    oneDayCost := costs[0] + mincostTicketsHelper(days, costs, i+1)
    var j int
    for j = i; j < len(days) && days[j] < days[i] + 7; j++ {}
    sevenDaysCost := costs[1] + mincostTicketsHelper(days, costs, j)
    for j = i; j < len(days) && days[j] < days[i] + 30; j++ {}
    oneMonthCost := costs[2] + mincostTicketsHelper(days, costs, j)
    return min(oneDayCost, min(sevenDaysCost, oneMonthCost))
}

func min(x, y int) int { if x < y { return x }; return y }
