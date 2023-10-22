const (
    Red = iota
    White
    blue
)

func sortColors(nums []int)  {
    i, j, k := 0, 0, len(nums) - 1
    for j <= k {
        if nums[j] == Red {
            nums[i], nums[j] = nums[j], nums[i]
            i++
            j++
        } else if nums[j] == White {
            j++
        } else {
            nums[j], nums[k] = nums[k], nums[j]
            k--
        }
    }
}
