func maxArea(h int, w int, horizontal []int, vertical []int) int {
    m, n := len(horizontal), len(vertical)
    sort.Ints(horizontal)
    sort.Ints(vertical)
    maxHeight := max(horizontal[0], h - horizontal[m-1])
    maxWidth := max(vertical[0], w - vertical[n-1])
    for i := 1; i < m; i++ {
        maxHeight = max(maxHeight, horizontal[i] - horizontal[i-1])
    }
    for i := 1; i < n; i++ {
        maxWidth = max(maxWidth, vertical[i] - vertical[i-1])
    }
    return maxHeight * maxWidth % (1e9 + 7)
}

func max(x, y int) int { if x > y { return x }; return y }
