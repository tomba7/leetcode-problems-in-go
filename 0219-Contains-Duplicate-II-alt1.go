func containsNearbyDuplicate(nums []int, k int) bool {
    indxMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        indx, ok := indxMap[nums[i]]
        if ok {
            if i - indx <= k { return true }
        }
        indxMap[nums[i]] = i
    }
    return false
}
