func containsNearbyDuplicate(nums []int, k int) bool {
    indxMap := make(map[int]struct{})
    for i, n := range nums {
        if _, ok := indxMap[n]; ok { return true }
        indxMap[n] = struct{}{} 
        if len(indxMap) > k { delete(indxMap, nums[i - k]) }
    }
    return false
}
