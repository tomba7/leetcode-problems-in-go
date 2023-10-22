func change(amount int, coins []int) int {
    memo := make([]int, amount + 1)
    memo[0] = 1
    for _, coin := range coins {
        for i := coin; i <= amount; i++ {
            memo[i] += memo[i - coin]
        }
    }
    return memo[amount]
}
