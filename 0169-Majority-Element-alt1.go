// https://discuss.leetcode.com/topic/6286/share-my-solution-java-count-bits
func majorityElement(nums []int) int {
    majority := 0
    for i := 0; i < 64; i++ {
        numSetBit := 0
        for j := 0; j < len(nums); j++ {
            if nums[j] & (1 << uint(i)) != 0 {
                numSetBit++
            }
        }
        if numSetBit > len(nums)/2 {
            majority |= 1 << uint(i)
        }
    }
    return majority
}
