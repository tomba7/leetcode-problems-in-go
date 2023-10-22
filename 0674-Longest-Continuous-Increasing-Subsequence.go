func findLengthOfLCIS(nums []int) int {
    var seq, longestSeq int
    for i := 0; i < len(nums); i++ {
        if i == 0 || nums[i] > nums[i - 1] {
            seq++
            longestSeq = max(longestSeq, seq)
        } else {
            seq = 1
        }
    }
    return longestSeq
}

func max(x, y int) int { if x > y { return x }; return y }
