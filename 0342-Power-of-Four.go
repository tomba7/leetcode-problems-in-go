// https://discuss.leetcode.com/topic/42855/o-1-one-line-solution-without-loops
func isPowerOfFour(num int) bool {
    return num > 0 && ((num & (num - 1)) == 0) && ((num & 0x55555555) == num)
}
