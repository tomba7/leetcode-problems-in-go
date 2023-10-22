func maxProfit(prices []int) int {
    var i, profit int
    for i < len(prices) - 1 {
        for i < len(prices) - 1 && prices[i] >= prices[i + 1] { i++ }
        valley := prices[i]
        for i < len(prices) - 1 && prices[i] < prices[i + 1] { i++ }
        peak := prices[i]
        profit += peak - valley
    }
    return profit
}

---

func maxProfit(prices []int) int {
    maxProfit := 0
    for i := 0; i < len(prices) - 1; i++ {
        if prices[i + 1] > prices[i] {
            maxProfit += prices[i + 1] - prices[i]
        }
    }
    return maxProfit
}

---

func maxProfit(prices []int) int {
    var profit int
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i - 1] {
            profit += prices[i] - prices[i - 1]
        }
    }
    return profit
}

---

// Brute force (TLE)
/*
    [7,1,5,3,6,4]

 - Use recursion
 - At each level
   * Buy the current price
   * Then loop through rest of the sell prices
   * Compute the sell price,
   * Then recurse at the next possible index
 - Gather all the result and return the maximum result
 */
func maxProfit(prices []int) int {
    return maxProfitHelper(prices, 0)
}

func maxProfitHelper(prices []int, i int) int {
    if i == len(prices) {
        return 0
    }
    var res int
    for ; i < len(prices); i++ {
        for j := i+1; j < len(prices); j++ {
            res = max(res, prices[j] - prices[i] + maxProfitHelper(prices, j+1))
        }
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }
