func summaryRanges(nums []int) (result []string) {
    for i, j := 0, 0; j < len(nums); j++ {
        if j + 1 == len(nums) || nums[j] != nums[j + 1] - 1 {
            var numRange string
            if i == j {
                numRange = fmt.Sprintf("%d", nums[i])
            } else {
                numRange = fmt.Sprintf("%d->%d", nums[i], nums[j])
            }
            result = append(result, numRange)
            i = j + 1
        }
    }
    return
}
