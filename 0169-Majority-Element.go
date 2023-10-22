// https://discuss.leetcode.com/topic/17446/6-suggested-solutions-in-c-with-explanations
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}
