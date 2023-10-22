func maxProfit(prices []int) int {
    n := len(prices)
    leftMin, rightMax := prices[0], prices[n-1]
    left := make([]int, n)
    right := make([]int, n+1)
    for i := 1; i < n; i++ {
        left[i] = max(left[i-1], prices[i]-leftMin)
        leftMin = min(leftMin, prices[i])
        j := n-i-1
        right[j] = max(right[j+1], rightMax-prices[j])
        rightMax = max(rightMax, prices[j])
    }
    var maxProfit int
    for i := 0; i < n; i++ {
        maxProfit = max(maxProfit, left[i]+right[i+1])
    }
    return maxProfit
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }

---

func maxProfit(prices []int) int {
    maxTotalProfit := 0
    firstBuySellProfits := make([]int, len(prices))
    minPriceSoFar := math.MaxInt32
    for i := 0; i < len(prices); i++ {
        minPriceSoFar = min(minPriceSoFar, prices[i])
        maxTotalProfit = max(maxTotalProfit, prices[i] - minPriceSoFar)
        firstBuySellProfits[i] = maxTotalProfit
    }
    maxPriceSoFar := math.MinInt32
    for i := len(prices) - 1; i > 0; i-- {
        maxPriceSoFar = max(maxPriceSoFar, prices[i])
        maxTotalProfit = max(maxTotalProfit, maxPriceSoFar - prices[i] + firstBuySellProfits[i - 1])
    }
    return maxTotalProfit
}

func min(x, y int) int {if x < y {return x}; return y}
func max(x, y int) int {if x > y {return x}; return y}
