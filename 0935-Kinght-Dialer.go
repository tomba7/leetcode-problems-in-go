func knightDialer(n int) int {
    mod := 1000000000 + 7
    g := [][]int{
        {4, 6}, {6, 8}, {7, 9}, {4, 8}, {3, 9, 0},
        {}, {0, 1, 7}, {2, 6}, {1, 3}, {2, 4},
    }
    dp := make([][10]int, n)
    for j := range dp[0] {
        dp[0][j] = 1
    }
    for l := 1; l < n; l++ {
        for digit := 0; digit < 10; digit++ {
            for _, nextDigit := range g[digit] {
                dp[l][digit] = (dp[l][digit] + dp[l-1][nextDigit]) % mod
            }
        }
    }
    var total int
    for j := 0; j < 10; j++ {
        total = (total + dp[n-1][j]) % mod
    }
    return total
}

---

func knightDialer(n int) int {
    dp := newDPTable(n)
    for l := 1; l < n; l++ {
        for digit := 0; digit < 10; digit++ {
            for _, nextDigit := range getNextDigits(digit) {
                dp[l][digit] = (dp[l][digit] + dp[l-1][nextDigit]) % (1e9+7)
            }
        }
    }
    var total int
    for digit := 0; digit < 10; digit++ {
        total = (total + dp[n-1][digit])%(1e9+7)
    }
    return total
}

func newDPTable(n int) [][10]int {
    dp := make([][10]int, n)
    for i := 0; i < 10; i++ {
        dp[0][i] = 1
    }
    return dp
}

func getNextDigits(digit int) []int {
    nextDigits := [][]int{
        {4, 6}, {6, 8}, {7, 9}, {4, 8}, {0, 3, 9},
        {}, {0, 1, 7}, {2, 6}, {1, 3}, {2, 4}, 
    }
    return nextDigits[digit]
}
