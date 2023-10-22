// https://discuss.leetcode.com/topic/33527/math-1-liner-no-log-with-explanation
func isPowerOfThree(n int) bool {
    return n > 0 && (int(math.Pow(3.0, 19.0)) % n == 0)
}
