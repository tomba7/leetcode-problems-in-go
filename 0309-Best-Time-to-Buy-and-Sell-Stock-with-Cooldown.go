func maxProfit(prices []int) int {
    sold := math.MinInt32
    held := math.MinInt32
    var reset int
    for _, price := range prices {
        preSold := sold
        sold = held + price
        held = max(held, reset - price)
        reset = max(reset, preSold)
    }
    return max(sold, reset)
}

func max(x, y int) int { if x > y { return x }; return y }
