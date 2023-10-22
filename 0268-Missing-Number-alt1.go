// https://discuss.leetcode.com/topic/24535/4-line-simple-java-bit-manipulate-solution-with-explaination
func missingNumber(nums []int) int {
    var result, i int
    for i = 0; i < len(nums); i++ {
        result ^= i ^ nums[i]
    }
    return result ^ i
}
