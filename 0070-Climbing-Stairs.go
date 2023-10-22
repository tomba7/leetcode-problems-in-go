/*
 * Tags: "Dynamic Programming"
 */
func climbStairs(n int) int {
    x, y := 1, 1
    for i := 2; i <= n; i++ {
        tmp := x
        x += y
        y = tmp
    }
    return x
}
