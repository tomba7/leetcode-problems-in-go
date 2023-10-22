func maxProfit(prices []int) int {
    var maxProfit int
    buy := prices[0]
    for i := 1; i < len(prices); i++ {
        maxProfit = max(maxProfit, prices[i]-buy)
        buy = min(buy, prices[i])
    }
    return maxProfit
}

func max(x, y int) int { if x > y { return x }; return y }
func min(x, y int) int { if x < y { return x }; return y }

---

func maxProfit(prices []int) (maxProfit int) {
    minPurchase := math.MaxInt32
    for i := range prices {
        maxProfit = max(maxProfit, prices[i] - minPurchase)
        minPurchase = min(minPurchase, prices[i])
    }
    return
}

func max(x, y int) int {if x > y {return x}; return y}
func min(x, y int) int {if x < y {return x}; return y}
