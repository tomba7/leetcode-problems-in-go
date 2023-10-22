func maximumSwap(num int) int {
    nums := []byte(strconv.Itoa(num))
    left, right, maxIndex := -1, -1, len(nums)-1
    for i := len(nums)-1; i >= 0; i-- {
        if nums[i] > nums[maxIndex] {
            maxIndex = i
        } else if nums[i] < nums[maxIndex] {
            left, right = i, maxIndex
        }
    }
    if left == -1 { return num }
    nums[left], nums[right] = nums[right], nums[left]
    num, _ = strconv.Atoi(string(nums))
    return num
}
