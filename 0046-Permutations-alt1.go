func permute(nums []int) [][]int {
    return permuteHelper(nums, 0)
}

func permuteHelper(nums []int, index int) (result [][]int) {
    if len(nums) == index {
        return [][]int{nil}
    }
    x := nums[index]
    perms := permuteHelper(nums, index + 1)
    for _, perm := range perms {
        for i := 0; i <= len(perm); i++ {
            new := append([]int{}, perm[:i]...)
            new = append(new, x)
            new = append(new, perm[i:]...)
            result = append(result, new)
        }
    }
    return result
}
