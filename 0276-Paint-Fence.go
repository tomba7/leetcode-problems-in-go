//
// https://discuss.leetcode.com/topic/83369/golang-solution-0ms-the-secret-is-to-realize-that-the-recursion-is-exactly-like-the-fibonacci-series
//
// The recursion is exactly like calculating the fibonacci series
//
// func numWays(n int, k int) int {
//     if n == 0 || k == 0 {
//         return 0
//     } else if n == 1 {
//         return k
//     } else if n == 2 {
//         return k * k
//     }
//     return (k - 1) * numWays(n - 1, k) + (k - 1) *numWays(n - 2, k))
//  }
//
func numWays(n int, k int) int {
    if n <= 0 || k <= 0 {return 0}
    if n == 1 {return k}
    if n == 2 {return k * k}    // k + k * (k - 1) becomes k * k
    x, y := k * k, k
    for i := 2; i < n; i++ {
        z := (k - 1) * (x + y)  // (k - 1) * f(n - 1) + (k - 1) * f(n - 2)
        y = x
        x = z
    }
    return x
}
